FROM golang:1.16-buster

WORKDIR /app

COPY ./ ./

RUN apt-get update
RUN apt-get -y install postgresql-client

RUN chmod +x wait-for-postgres.sh

RUN go mod download
RUN go build -o accounts ./cmd/main.go

EXPOSE 8002
CMD ["./accounts"]