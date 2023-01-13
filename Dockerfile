FROM golang:1.20rc3
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
WORKDIR /go/src/github.com/stellar/go

COPY . .
RUN dep ensure -v
RUN go install github.com/stellar/go/tools/...
RUN go install github.com/stellar/go/services/...
