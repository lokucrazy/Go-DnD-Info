FROM golang:latest

RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go get -v github.com/mattn/go-sqlite3
RUN go get -v github.com/go-chi/chi
RUN go get -v github.com/go-chi/cors
RUN go build -o main .
EXPOSE 8080
CMD ["./main"]