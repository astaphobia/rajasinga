FROM golang:1.11 as base
ARG mode=development
WORKDIR /go/src/github.com/astaphobia/rajasinga/server
COPY . .
RUN if ! which realize > /dev/null;\
    then go get github.com/oxequa/realize; \
    fi
RUN if ! which dep > /dev/null;\
    then go get github.com/golang/dep/...; \
    fi
RUN dep ensure -v

FROM base as prod 
RUN go build -o server .
ENTRYPOINT [ "./server" ]

FROM base as dev
ENTRYPOINT ["realize", "start"]
