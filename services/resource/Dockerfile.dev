FROM golang:1.17

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

COPY ./ /app

RUN go mod download

RUN go get github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon --build="go build cmd/serve.go" --command=./serve
