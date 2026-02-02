FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
COPY app.env .

RUN apk update && apk add --no-cache \
    tzdata \
    libgdiplus \
    iputils \
    curl \
    && cp /usr/share/zoneinfo/Asia/Bangkok /etc/localtime \
    && echo "Asia/Bangkok" > /etc/timezone \
    && ln -s /usr/lib/libgdiplus.so /usr/lib/libgdiplus-native.so

EXPOSE 8080
CMD [ "/app/main" ]