FROM golang:1.21 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY ./cmd/*.go ./

RUN CGO_ENABLED=0 GOOS=linux go build main.go

FROM alpine:3.18 AS final

WORKDIR /build

COPY --from=builder /app/main /build/main

EXPOSE 3000

ENTRYPOINT ["/build/main"]