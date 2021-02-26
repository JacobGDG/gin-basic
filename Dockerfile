FROM golang:1.16.0-buster
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go mod tidy
RUN go build -o main .

CMD ["/app/main"]

