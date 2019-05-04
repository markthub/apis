FROM golang
MAINTAINER Davide Berdin "davideberdin@gmail.com"

WORKDIR /go/src/github.com/markthub/apis/api

RUN apt-get update
RUN apt-get install -y vim

COPY . .

RUN mkdir -p ./bin && make build

EXPOSE 8000

CMD ["./bin/apis", "server"]