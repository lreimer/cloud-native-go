FROM golang:1.7.3-alpine

ENV SOURCES /go/src/github.com/lreimer/cloud-native-go/

COPY . ${SOURCES}

RUN cd ${SOURCES} && CGO_ENABLED=0 go install -a

ENTRYPOINT cloud-native-go
EXPOSE 8080