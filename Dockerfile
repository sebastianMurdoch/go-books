FROM golang:1.8.1-alpine

RUN apk update && apk upgrade && apk add --no-cache bash git

RUN go get github.com/sebastianMurdoch/go-books

ENV SOURCES /go/src/github.com/sebastianMurdoch/go-books/
RUN cd ${SOURCES} && CGO_ENABLED=0 go build

WORKDIR ${SOURCES}
CMD ${SOURCES}go-books
EXPOSE 8080

