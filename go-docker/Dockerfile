# Build stage
FROM golang:1.11-alpine as builder

RUN mkdir /app
COPY  . /app

WORKDIR /app

RUN go build -o main .

# Run stage
FROM scratch

WORKDIR /

COPY --from=builder /app/main .

CMD ["./main"]
