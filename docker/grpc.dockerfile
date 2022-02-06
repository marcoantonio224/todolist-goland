FROM golang:alpine

WORKDIR /grpc

ADD ./grpc .

RUN go get github.com/githubnemo/CompileDaemon
COPY . .

COPY ./entrypoint/grpc.sh /entrypoint/grpc.sh

ADD https://raw.githubusercontent.com/eficode/wait-for/v2.2.1/wait-for /usr/local/bin/wait-for

RUN chmod +rx /usr/local/bin/wait-for /entrypoint/grpc.sh

ENTRYPOINT ["sh", "/entrypoint/grpc.sh"]