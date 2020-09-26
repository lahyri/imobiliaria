FROM golang:1.13.9-stretch as builder

RUN apt-get update && apt-get install -y xz-utils && rm -rf /var/lib/apt/lists/*
ADD https://github.com/upx/upx/releases/download/v3.96/upx-3.96-amd64_linux.tar.xz /usr/local
RUN xz -d -c /usr/local/upx-3.96-amd64_linux.tar.xz | tar -xOf - upx-3.96-amd64_linux/upx > /bin/upx && chmod a+x /bin/upx
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-s -w" ./cmd/core/main.go; upx main

######## Start a new stage from scratch #######
FROM alpine:latest

RUN apk --no-cache add ca-certificates
EXPOSE 80
WORKDIR /root/
COPY --from=builder /app/main .
CMD ["./main"] 