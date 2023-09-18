# TODO Web Service

TODO Web Service is a web API that would manage a to do list. Authentication is left out on purpose. So validation for users are left out on purpose.

## Install

    go mod vendor
    go build

## Run the service

    go run .

## Run with Docker

    docker build ./
    docker run -d -p 8080:8080 {id}

## REST API
The REST API for the TODO web service is described below.

## Get Task for user
### Request
```GET /v1/task/ ```

    curl --location 'http://localhost:8080/v1/task?userId=1'

## Post a task
### Request
```POST /v1/task/ ```

    curl --location 'http://localhost:8080/v1/task' \
        --header 'Content-Type: application/json' \
        --data '{
            "title": "Ace the second round",
            "description": "go through the second rounds of interview",
            "userId": 1
        }'

## Patch a task
### Request
```PATCH /v1/task/{id} ```

    curl --location --request PATCH 'http://localhost:8080/v1/task/3' \
        --header 'Content-Type: application/json' \
        --data '{
            "status": "C"
        }'

## Delete a task
### Request
```DELETE /v1/task/{id} ```

    curl --location --request DELETE 'http://localhost:8080/v1/task/3'
