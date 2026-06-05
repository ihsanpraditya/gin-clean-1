FROM golang:1.26-alpine

WORKDIR /usr/src/app

RUN go install github.com/air-verse/air@latest

COPY go.mod go.su[m] ./
RUN go mod download

EXPOSE 8080

CMD ["air"]
