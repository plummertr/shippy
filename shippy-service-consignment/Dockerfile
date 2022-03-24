FROM golang:1.17-alpine as builder

RUN apk update && apk upgrade && apk add --no-cache git

RUN mkdir /app
WORKDIR /app

COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -o shippy-service-consignment

FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=builder /app/shippy-service-consignment .

CMD ["./shippy-service-consignment"]