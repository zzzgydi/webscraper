FROM golang:latest AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go mod download
COPY . .

RUN mkdir -p output
RUN CGO_ENABLED=0 go build -tags netgo -ldflags '-extldflags "-static"' -o output/webscraper ./cmd/main.go


FROM alpine:latest

WORKDIR /app


RUN mkdir -p /app/log

COPY --from=builder /app/output /app
COPY --from=builder /app/assets /app/assets
COPY --from=builder /app/config /app/config
COPY --from=builder /app/static /app/static

RUN chmod +x /app/webscraper

ENV CONFIG_ENV=prod

EXPOSE 4090
CMD ["./webscraper"]
