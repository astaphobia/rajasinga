FROM golang:1.11
WORKDIR /go/src/github.com/astaphobia/rajasinga/
COPY . .
RUN if ! which dep > /dev/null;\
    then go get github.com/golang/dep/...; \
    fi
RUN dep ensure -v
RUN go build -o server .
ENTRYPOINT [ "./server" ]
