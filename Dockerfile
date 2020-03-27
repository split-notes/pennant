FROM golang:1.10

WORKDIR /go/src/bili

RUN go get github.com/githubnemo/CompileDaemon

## This container exposes port 8080 to the outside world
EXPOSE 8080

#ENTRYPOINT CompileDaemon -log-prefix=false -build="go build ./" -command="./bili dev"
ENTRYPOINT CompileDaemon -log-prefix=false -build="go build -v -o bili-dev ./" -command="./bili-dev dev"
