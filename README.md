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

### Response

    HTTP/1.1 200 OK
    Date: Mon, 18 Sept 2023 12:36:33 GMT
    Status: 200 OK
    Connection: close
    Content-Type: application/json
    Content-Length: 287

    [
        {
            "id": 2,
            "title": "Ace the first round",
            "description": "go through the first rounds of interview",
            "status": "A",
            "userId": 1,
            "createdAt": "2023-09-18T18:58:29.063666Z",
            "modifiedAt": "2023-09-18T18:58:29.063666Z"
        }
    ]

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

### Response

    HTTP/1.1 200 OK
    Date: Thu, 24 Feb 2011 12:36:32 GMT
    Status: 200 OK
    Connection: close
    Content-Type: application/json
    Content-Length: 1

    3

## Patch a task
### Request
```PATCH /v1/task/{id} ```

    curl --location --request PATCH 'http://localhost:8080/v1/task/3' \
        --header 'Content-Type: application/json' \
        --data '{
            "title": "Ace the second round",
            "description": "go through the second rounds of interview",
            "status": "C"
        }'

### Response

    HTTP/1.1 204 No Content
    Date: Mon, 18 Sept 2023 12:36:33 GMT
    Status: 204 No Content
    Connection: close

## Delete a task
### Request
```DELETE /v1/task/{id} ```

    curl --location --request DELETE 'http://localhost:8080/v1/task/3'

### Response

    HTTP/1.1 204 No Content
    Date: Mon, 18 Sept 2023 12:36:33 GMT
    Status: 204 No Content
    Connection: close
