FROM golang:1.13.3

WORKDIR /go/src/gomicro/
COPY . .
RUN go fmt ./... && go build web_server.go

CMD ["./web_server"]

