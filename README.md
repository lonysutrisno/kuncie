# Kuncie

Kuncie is a checkout system app develop on Go Language with interfaces REST API

## Features

- Checkout

## Tech Stack

Kuncie uses a number of open source projects to work properly:

- Golang 1.12
- Mysql 5.7
- Docker-compose v3

## Installation

Kuncie requires Docker and Docker-compose to create image based on [Golang](https://golang.org/doc/go1.12) v1.12 and [MYSQL](https://dev.mysql.com/downloads/mysql/5.7.html) v5.7 to run.


Clone this repo to specific folder on your go project folder ```go/src/github.com/lonysutrisno/kuncie```
Next step:

```sh
cd kuncie
docker-compose up
```

For running the unit-test...

```sh
cd kuncie
go test ./...
```

# Documentation
For Rest API documentation you can find it in project folder: ```go/src/github.com/lonysutrisno/kuncie/kuncie.postman_collection.json```
