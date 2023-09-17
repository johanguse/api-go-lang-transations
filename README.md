# Transactions API
In 2022, I took a front-end test to work at Warren. The API used in the test was hosted on Heroku, but it's currently offline. To continue consuming my front-end code, I decided to create my own API using Go and Fiber. You can find the front-end code for the project at https://github.com/johanguse/.

I'm happy to say that I passed the test and worked at the company for some time.

## Getting Started

### Prerequisites

- Go 1.21 or laterg
- MySQL 5.7 or later

### Installing

1. Clone the repository:
```console
git clone https://github.com/johanguse/api-go-lang-transations.git
```
2. Install the dependencies:
```console
cd transactions-api
go mod download
```
3. Set up the database:
```console
mysql -u root -p < database/schema.sql
```
4. Start the server:
```console
go run main.go
```
The server will start on http://127.0.0.1:8000/api/.

You can use the healthcheck to 

## API Reference

#### Postman Collection

[![Run in Postman](https://run.pstmn.io/button.svg)](https://app.getpostman.com/run-collection/2499608-f2fe1930-17a0-4b66-9d67-1278fe0dc427?action=collection%2Ffork&source=rip_markdown&collection-url=entityId%3D2499608-f2fe1930-17a0-4b66-9d67-1278fe0dc427%26entityType%3Dcollection%26workspaceId%3Dc837a68d-911f-48c1-8f68-a9b87bc7c51a)

Base URL

```console
http://127.0.0.1:8000/api
```

### Healthcheck
```console
GET /healthchecker
```
You can use the healthcheck to verify the status and proper functioning of your application. It returns a status of success and a welcome message.

```go
"status":  "success",
"message": "Welcome to Golang, Fiber, and GORM",
```

### List Transactions
```console
GET /transactions
```
List transactions with pagination.

### Create Transaction
```console
POST /transactions
```
Create a new transaction.

### Find Transaction by ID
```console
GET /transactions/:id
```
Find a transaction by its unique ID.

### Update Transaction
```console
PUT /transactions/:id
```
Update a transaction by its unique ID.

### Delete Transaction
```console
DELETE /transactions/:id
```
Delete a transaction by its unique ID.

## License

This project is licensed under the MIT License - see the LICENSE file for details.