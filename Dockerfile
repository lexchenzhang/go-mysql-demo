FROM golang:1.13-buster AS builder
WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .

RUN go build -o /go-mysql

FROM gcr.io/distroless/base-debian10
WORKDIR /
COPY --from=builder /go-mysql /go-mysql
EXPOSE 8000
CMD ["/go-mysql"]