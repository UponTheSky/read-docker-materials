FROM golang:1.16-alpine AS build-stage

WORKDIR /usr/src/app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

# ENV REQUEST_ORIGIN=http://localhost:5000 
ENV REQUEST_ORIGIN=http://localhost

# see: https://stackoverflow.com/questions/49955097/how-do-i-add-a-user-when-im-using-alpine-as-a-base-image
# also check: https://reakwon.tistory.com/137
RUN CGO_ENABLED=0 GOOS=linux go build
RUN echo "nobody:x:65534:65534:Nobody:/:" > ./etc_passwd && chmod +x ./server

FROM scratch

WORKDIR /usr/src/app

# for adduser, see: https://gist.github.com/kszarek/00f3cc3117a56e92956a06ffabe30169
COPY --from=build-stage /usr/src/app/server ./server
COPY --from=build-stage /usr/src/app/etc_passwd /etc/passwd

EXPOSE 8080
USER nobody

CMD ["./server"]
