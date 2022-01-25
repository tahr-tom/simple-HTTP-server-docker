# Simple HTTP Server Docker
Written in Go with [Go-Chi](https://github.com/go-chi/chi)
## Getting Started
1. Install [Docker](https://docs.docker.com/install/)
2. Clone this repository
3. In the repository directory, run `$ docker build -t test-server .`
4. Start the container with `$ docker run --rm -p 80:80 test-server`
5. Now the server should be running on http://localhost
##  Endpoints
### `GET /list`

Accept a GET request and return a JSON response, each objects include two string, namely `key` and `value` and a `timestamp` 

The objects are sorted in descending order by `timestamp`

#### Example

Running
```shell script
$ curl http://localhost/list | python -m json.tool
```
Output with formatting
```json
[
    {
        "key": "asdf",
        "timestamp": "2019-12-02T06:53:35Z",
        "value": "some other Value"
    },
    {
        "key": "a",
        "timestamp": "2019-12-02T06:53:32Z",
        "value": "same Value"
    }
]
```

This two items are hardcoded in the server the demonstration purpose
### `POST /add`
Accept a POST request with json payload that includes two string attribute `key` and `value`
#### Example
Running
```shell script
$ curl --header "Content-Type: application/json" \
      --request POST \
      --data '{"key": "key a", "value": "value a"}' \
      http://localhost:80/add

$ curl --header "Content-Type: application/json" \
      --request POST \
      --data '{"key": "key b", "value": "value b"}' \
      http://localhost:80/add

$ curl --header "Content-Type: application/json" \
      --request POST \
      --data '{"key": "key c", "value": "value c"}' \
      http://localhost:80/add

$ curl http://localhost/list | python -m json.tool
```
Output with formatting
```json
[
    {
        "key": "key c",
        "timestamp": "2020-02-20T20:19:02Z",
        "value": "value c"
    },
    {
        "key": "key b",
        "timestamp": "2020-02-20T20:18:58Z",
        "value": "value b"
    },
    {
        "key": "key a",
        "timestamp": "2020-02-20T20:18:45Z",
        "value": "value a"
    },
    {
        "key": "asdf",
        "timestamp": "2019-12-02T06:53:35Z",
        "value": "some other Value"
    },
    {
        "key": "a",
        "timestamp": "2019-12-02T06:53:32Z",
        "value": "same Value"
    }
]
```


