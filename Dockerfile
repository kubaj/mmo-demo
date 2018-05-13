FROM golang:1.9

WORKDIR /go/src/github.com/kubaj/mmo-demo
ADD . .

RUN mkdir -p /build

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /build/auth auth/cmd/auth/main.go
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /build/order order/cmd/order/main.go

FROM alpine:latest

WORKDIR /root/

RUN apk --no-cache add ca-certificates
COPY --from=0 /build .