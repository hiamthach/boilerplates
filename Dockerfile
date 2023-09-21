FROM golang:1.20-alpine AS builder

WORKDIR /app
COPY go.mod ./
COPY go.sum ./
COPY app.env ./
#RUN go mod download
RUN go get -d -v ./...
#RUN go mod vendor

COPY . .

RUN go build -o /go-microservices

EXPOSE 8080

CMD ["/go-microservices"]