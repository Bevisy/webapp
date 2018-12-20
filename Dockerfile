FROM docker.io/golang:1.11.3-alpine3.8
WORKDIR $GOPATH/src
RUN apk add --no-cache git \
    && git clone https://github.com/Bevisy/webapp.git \
    && cd webapp \
    && go build webapp.go \
    && mv webapp /
ENTRYPOINT ["/webapp"]
