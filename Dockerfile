FROM golang:1.3-cross
ADD . /go/src/github.com/cglewis/dockerdash
WORKDIR /go/src/github.com/cglewis/dockerdash
ENV GOOS linux
ENV GOARCH amd64
RUN go get
ENTRYPOINT ["/go/src/github.com/cglewis/dockerdash/make.sh"]
