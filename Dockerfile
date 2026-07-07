FROM golang:1.26-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY * .
RUN go build -o /identity-svc

EXPOSE 8080

CMD ["/identity-svc"]