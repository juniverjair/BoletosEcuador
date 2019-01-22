FROM golang

RUN mkdir -p /go/src/BoletosEcuador

ADD . /go/src/BoletosEcuador

RUN go get  -t -v ./...
RUN go get  github.com/canthefason/go-watcher
RUN go install github.com/canthefason/go-watcher/cmd/watcher

ENTRYPOINT  watcher -run BoletosEcuador/reporte/cmd  -watch BoletosEcuador/reporte
