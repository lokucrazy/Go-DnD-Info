FROM golang:latest

RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go mod download
RUN go build -o main .
EXPOSE 8081
CMD ["./main", "8081"]