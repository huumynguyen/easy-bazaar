FROM golang:alpine as builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o /main

FROM alpine:latest
WORKDIR /
COPY --from=builder /main /main
EXPOSE 8080
CMD [ "/main" ]

