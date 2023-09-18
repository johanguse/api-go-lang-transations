FROM golang:1.16-alpine
WORKDIR /app
COPY . /app
RUN go build -o main .
EXPOSE 8000
CMD ["./main"]