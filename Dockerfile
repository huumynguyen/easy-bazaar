FROM golang:latest as builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build .

FROM alpine:latest
COPY --from=builder /app/main /app
EXPOSE 8080
ENTRYPOINT [ "main" ]

