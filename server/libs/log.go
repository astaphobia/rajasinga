package libs

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

type (
	fimHook struct {
		Writer    *os.File
		LogLevels []logrus.Level
	}
	timer interface {
		Now() time.Time
		Since(time.Time) time.Duration
	}
	realClock struct{}
	Logger    struct {
		logger *logrus.Logger
		clock  timer
	}
)

func (rc *realClock) Now() time.Time {
	return time.Now()
}
func (rc *realClock) Since(t time.Time) time.Duration {
	return time.Since(t)
}

func (hook *fimHook) Fire(entry *logrus.Entry) error {
	line, err := entry.String()
	if err != nil {
		fmt.Fprintf(os.Stderr, string(line))
	}
	_, err = hook.Writer.Write([]byte(line))
	return err
}
func (hook *fimHook) Levels() []logrus.Level {
	return hook.LogLevels
}

func (l *Logger) LoggerInit() {
	var (
		logLevels []logrus.Level
		log       *logrus.Logger
	)

	logFile, err := os.OpenFile("debug.log", os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		panic(err)
	}

	levelConf := os.Getenv("LOG_LEVEL")
	if len(levelConf) >= 1 {
		for _, level := range strings.Split(levelConf, ",") {
			switch level {
			case "info":
				logLevels = append(logLevels, logrus.InfoLevel)
			case "warning":
				logLevels = append(logLevels, logrus.WarnLevel)
			case "debug":
				logLevels = append(logLevels, logrus.DebugLevel)
			case "error":
				logLevels = append(logLevels, logrus.ErrorLevel)
			default:
				logLevels = append(logLevels, logrus.ErrorLevel)
			}
		}
	} else {
		logLevels = append(logLevels, logrus.InfoLevel)
		logLevels = append(logLevels, logrus.ErrorLevel)
	}

	log = logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetReportCaller(true)
	log.Hooks.Add(&fimHook{
		Writer:    logFile,
		LogLevels: logLevels,
	})
	l.logger = log
	l.clock = &realClock{}
}
func (l Logger) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
			entry = logrus.NewEntry(l.logger)
			start = l.clock.Now()
		)
		if reqID := r.Header.Get("X-Request-Id"); len(reqID) >= 1 {
			entry = entry.WithField("requestId", reqID)
		}

		entry.WithFields(logrus.Fields{
			"request": r.RequestURI,
			"method":  r.Method,
			"ip":      r.RemoteAddr,
			"latency": l.clock.Since(start),
		}).Info("Kings")
	})
}
