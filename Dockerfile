#FROM golang:latest
#
#ENV GOPROXY https://goproxy.cn,direct
#WORKDIR $GOPATH/src/github.com/Pluto00/Template
#COPY . $GOPATH/src/github.com/Pluto00/Template
#
#RUN go build .
#
#EXPOSE 8000
#ENTRYPOINT ["./Template"]

FROM scratch

WORKDIR $GOPATH/src/github.com/Pluto00/Template
COPY . $GOPATH/src/github.com/Pluto00/Template

EXPOSE 8000
ENTRYPOINT ["./Template"]