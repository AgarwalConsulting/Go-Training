FROM golang:latest AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN mkdir -p /run
ENV CGO_ENABLED=0
RUN go build -o /run/server ./cmd/server

FROM scratch
COPY --from=builder /run/server /
ENV GRPC_PORT 50001
EXPOSE 50001
ENV DIAGNOSTICS_PORT 8081
EXPOSE 8081
CMD [ "/server" ]
