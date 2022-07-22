FROM golang:alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o app ./cmd/ 

WORKDIR /dist

RUN cp /build/app .
RUN cp /build/.env .

FROM scratch

COPY --from=builder /dist/app . /dist/.env ./

ENTRYPOINT ["/app"]
