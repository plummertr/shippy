FROM golang:1.17-alpine as builder

RUN apk update && apk upgrade && apk add --no-cache git

RUN mkdir /app
WORKDIR /app

COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -o shippy-cli-consignment

FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
ADD consignment.json /app/consignment.json
COPY --from=builder /app/shippy-cli-consignment .

CMD ["./shippy-service-consignment"]