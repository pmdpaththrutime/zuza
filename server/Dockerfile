FROM golang:1.27.0 AS builder
WORKDIR /app

COPY go.mod ./
RUN go mod download
COPY . .

# CGO required to compile SQLite dependency
RUN CGO_ENABLED=1 go build -o /server

FROM alpine:3.23.5
RUN apk add --no-cache sqlite-libs
WORKDIR /app
COPY --from=builder /server /server

EXPOSE 3333
ENTRYPOINT ["/server"]
