FROM golang:alpine

WORKDIR /todolist

ADD . .

RUN go mod download && go mod verify

RUN go get github.com/githubnemo/CompileDaemon
COPY . .
COPY ./entrypoint/app.sh /entrypoint/app.sh

ADD https://raw.githubusercontent.com/eficode/wait-for/v2.2.1/wait-for /usr/local/bin/wait-for
RUN chmod +rx /usr/local/bin/wait-for /entrypoint/app.sh

ENTRYPOINT ["sh", "/entrypoint/app.sh"]