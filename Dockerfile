FROM golang:alpine AS builder
WORKDIR /app
COPY . .
RUN go build -ldflags="-w -s" -o chatgpt-arkose-token-api main.go

FROM alpine
WORKDIR /app
COPY --from=builder /app/chatgpt-arkose-token-api .
RUN apk add --no-cache tzdata
ENV TZ=Asia/Shanghai
EXPOSE 8080
CMD ["/app/chatgpt-arkose-token-api"]
