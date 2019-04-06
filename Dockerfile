FROM golang:alpine

WORKDIR /go/src/github.com/andreoav/brazilian-utils-service

ADD . .

RUN go install -v -tags heroku github.com/andreoav/brazilian-utils-service/cmd/server

EXPOSE 50051

CMD ["/go/bin/server"]