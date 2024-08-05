FROM golang:1.18 as builder
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o finance-planner

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/finance-planner .
CMD ["./finance-planner"]
