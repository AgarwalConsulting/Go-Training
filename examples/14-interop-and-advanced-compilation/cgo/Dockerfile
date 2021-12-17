FROM golang AS builder
WORKDIR /app
COPY . .
ENV CGO_ENABLED 0
RUN go build -o app simple.go

FROM scratch
WORKDIR /app
COPY --from=builder /app/app .
CMD [ "/app/app" ]
