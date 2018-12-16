FROM golang:1.11

COPY . /go/src/exampleddd
WORKDIR /go/src/exampleddd

RUN export GO111MODULE=on && go get github.com/oxequa/realize/...
RUN  export GO111MODULE=on && go get