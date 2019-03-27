FROM golang:1.11 as base

WORKDIR /go/src/github.com/astaphobia/rajasinga/server
COPY . .
RUN if ! which realize > /dev/null;\
    then go get github.com/oxequa/realize; \
    fi
RUN if ! which dep > /dev/null;\
    then go get github.com/golang/dep/...; \
    fi
RUN dep ensure -v
RUN go build -o server .

FROM base as prod 
ENTRYPOINT [ "./server" ]

FROM base as dev
ENTRYPOINT ["realize", "start"]
