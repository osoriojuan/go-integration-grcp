FROM golang:1.19.1-alpine3.16

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go install github.com/osoriojuan/go-integration-grcp

EXPOSE 8080

CMD go run main.go