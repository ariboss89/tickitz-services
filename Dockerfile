FROM golang:alpine3.22 AS build

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o server cmd/main.go

FROM alpine:edge

WORKDIR /app

COPY --from=build /app/server ./server

COPY public ./public

RUN chmod +x server

EXPOSE 8080

CMD [ "./server" ]