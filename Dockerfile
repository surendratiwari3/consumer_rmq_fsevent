FROM golang:1.16.3-alpine as builder
ENV GO111MODULE=on
WORKDIR /go/src/github.com/consumer_rmq_fsevent
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /
COPY --from=builder /go/src/github.com/consumer_rmq_fsevent .

CMD ["/app"]