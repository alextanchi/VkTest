FROM golang:1.20-alpine
RUN apk add build-base

EXPOSE 8080
WORKDIR /app
COPY ./ /app
WORKDIR /app/

RUN go mod download && go mod verify && go build -o ./service
CMD ["./service"]