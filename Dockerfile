FROM golang:alpine

RUN  mkdir -p /go/src \
  && mkdir -p /go/bin \
  && mkdir -p /go/pkg

RUN mkdir -p $GOPATH/src/app 
ADD . $GOPATH/src/app
WORKDIR $GOPATH/src/app 
RUN go get

RUN go build -o main . 
RUN go install
CMD ["app"]

EXPOSE 8080