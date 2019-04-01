package libs

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

type RequestID struct {
	ContextKey string
}

func (rID RequestID) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
			ctx = r.Context()
			rID = RequestID{
				ContextKey: "requestId",
			}
		)
		reqID := r.Header.Get("X-Request-Id")
		if len(reqID) <= 0 {
			reqID = uuid.New().String()
		}
		r.Header.Set("X-Request-Id", reqID)

		ctx = context.WithValue(ctx, rID.ContextKey, reqID)
		r = r.WithContext(ctx)
		if next != nil {
			next.ServeHTTP(w, r)
		}
	})
}
