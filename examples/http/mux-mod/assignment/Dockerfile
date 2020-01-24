FROM golang:latest
WORKDIR /go/src/app
COPY . .
RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure
ENV GOOS=linux
RUN go build .
EXPOSE 9000
ENTRYPOINT [ "/go/src/app/app" ]
