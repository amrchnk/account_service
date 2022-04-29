FROM golang:1.16

RUN mkdir /app

COPY ./ ./

RUN go mod download
RUN go build -o accounts ./cmd/main.go
CMD ["/accounts"]
EXPOSE 8002