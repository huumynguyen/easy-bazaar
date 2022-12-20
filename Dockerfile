FROM golang:alpine as builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o /app/main

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/ /app
EXPOSE 8080
CMD [ "/app/main" ]

