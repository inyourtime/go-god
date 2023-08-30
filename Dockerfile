FROM golang:1.20.7-alpine AS builder
WORKDIR /go/src
COPY . .
RUN go get
RUN go build -o /go/bin/server

FROM alpine
COPY --from=builder /go/bin/server /app/server
COPY --from=builder /go/src/config.yaml /app

WORKDIR /app
EXPOSE 8000
CMD ["./server"]