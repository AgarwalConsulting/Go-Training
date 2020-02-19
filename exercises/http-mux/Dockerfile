FROM golang:latest AS builder
WORKDIR /go/src/app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
RUN go build -tags netgo -a -v .

FROM scratch
WORKDIR /app
COPY --from=builder /go/src/app/library /app/library
EXPOSE 9000
ENTRYPOINT [ "/app/library" ]
