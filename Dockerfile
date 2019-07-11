FROM docker.io/golang:1.11.3-alpine3.8
WORKDIR $GOPATH/src
RUN mkdir -p webapp
COPY webapp.go webapp/webapp.go
RUN cd webapp \
    && go build webapp.go \
    && chmod +x webapp \
    && mv webapp /
ENTRYPOINT ["/webapp"]
