FROM golang:1.16-alpine

WORKDIR /usr/src/app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

# ENV REQUEST_ORIGIN=http://localhost:5000 
ENV REQUEST_ORIGIN=http://localhost

EXPOSE 8080

RUN go build

# see: https://stackoverflow.com/questions/49955097/how-do-i-add-a-user-when-im-using-alpine-as-a-base-image
RUN addgroup -S usergroup && adduser -S backuser -G usergroup
USER backuser

CMD ["./server"]
