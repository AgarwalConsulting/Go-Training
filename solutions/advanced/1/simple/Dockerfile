FROM golang:latest AS builder
WORKDIR /go/src/app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
ENV CGO_ENABLED=0
RUN go build -tags netgo -a .

FROM scratch
WORKDIR /app
COPY --from=builder /go/src/app/biblioteca .
EXPOSE 9000
ENTRYPOINT [ "/app/biblioteca" ]
