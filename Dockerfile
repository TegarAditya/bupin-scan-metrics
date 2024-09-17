FROM golang:1.23.1

WORKDIR /usr/src/app

COPY . .

RUN go mod tidy

EXPOSE 3030