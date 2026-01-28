FROM golang:tip-trixie AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

RUN go build -o server cmd/main.go

FROM alpine:edge

WORKDIR /app

COPY --from=builder /build/server ./server

RUN chmod +x server

EXPOSE 8080

CMD [ "./server" ]


