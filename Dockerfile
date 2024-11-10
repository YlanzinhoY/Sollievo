FROM golang:latest AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o sollievo .


FROM alpine:latest as final
WORKDIR /app
RUN apk add --no-cache bash
COPY --from=builder /app/sollievo .
ENTRYPOINT ["/app/sollievo"]
