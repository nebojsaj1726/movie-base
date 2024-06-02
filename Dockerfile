FROM golang:1.21-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN GOOS=linux GOARCH=amd64 go build -o server cmd/api/main.go

FROM alpine:latest

WORKDIR /app

RUN apk add --no-cache postgresql-client

COPY ./wait-for-pg.sh ./
RUN chmod +x wait-for-pg.sh

COPY --from=builder /app/server .
COPY .env /app/.env

COPY ui/dist /app/ui/dist

EXPOSE 8080
