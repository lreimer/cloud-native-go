FROM golang:1.7.3-alpine

ENV SOURCES /go/src/github.com/lreimer/cloud-native-go/

COPY . ${SOURCES}

RUN cd ${SOURCES} && go install

ENTRYPOINT cloud-native-go
EXPOSE 8080