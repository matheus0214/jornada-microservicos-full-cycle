FROM golang:1.15.2

WORKDIR /go/src
COPY . .

RUN GOOS=linux go build -ldflags="-s -w"
CMD ["./product"]