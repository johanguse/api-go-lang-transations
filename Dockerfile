FROM golang:alpine

RUN apk add --no-cache git ca-certificates


ARG GITHUB_TOKEN
ENV CGO_ENABLED=0 GO111MODULE=on GOOS=linux TOKEN=$GITHUB_TOKEN

WORKDIR /app
COPY . /app
RUN go mod download 
RUN go build -o main .
EXPOSE 8000
CMD ["./main"]