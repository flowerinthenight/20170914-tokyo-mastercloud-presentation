FROM golang

ADD . /go/src/demoapp
WORKDIR /go/src/demoapp
RUN go build -v
ENTRYPOINT ["/go/src/demoapp/demoapp"]
