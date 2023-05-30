FROM ubuntu:20.04

RUN apt-get update && apt-get install curl -y

WORKDIR /usr/src/app

COPY ./ex7.sh .

RUN chmod +x ex7.sh

CMD ./ex7.sh
