FROM golang:latest AS builder

WORKDIR /app

COPY . .
RUN go mod tidy
RUN go mod vendor
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/basicnats ./cmd

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/basicnats /app/basicnats

EXPOSE 8080

CMD ["/app/basicnats"] 

