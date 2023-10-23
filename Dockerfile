FROM golang:alpine AS build

WORKDIR /app
COPY . .

RUN go build -o gochat .

FROM alpine

WORKDIR /app
COPY --from=build /app/gochat .

CMD ["./gochat"]
