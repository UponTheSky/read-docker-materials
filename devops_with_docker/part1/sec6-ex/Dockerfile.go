FROM golang:1.16-alpine

WORKDIR /usr/src/app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN go build

EXPOSE 8080

ENV REQUEST_ORIGIN=http://localhost:3000

CMD ["./server"]
