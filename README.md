# TODO Web Service

TODO Web Service is a web API that would manage a to do list. Authentication is left out on purpose. So validation for users are left out on purpose.

## Install

    go mod vendor
    go build

## Run the service

    go run .

## Run the tests

    go test

## REST API
The REST API for the TODO web service is described below.

## Get Task for user
### Request
```GET /v1/task/ ```

    curl -i -H 'Accept: application/json' http://localhost:8080/v1/task

## Post a task
### Request
```POST /v1/task/ ```

## Patch a task
### Request
```PATCH /v1/task/{id} ```

## Delete a task
### Request
```DELETE /v1/task/{id} ```


