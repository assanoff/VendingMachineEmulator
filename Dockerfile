
FROM golang:1.13 AS builder
RUN mkdir /app
ADD . /app
WORKDIR /app
# We want to build our application's binary executable
RUN CGO_ENABLED=0 GOOS=linux make build

FROM alpine:latest AS production

COPY --from=builder /app/bin/ .
CMD ["./vendingmachine"]