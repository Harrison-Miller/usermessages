# Golang REST API exercise
A simple program that allows users to set and get messages.
It uses a fake external api for authenticating users (with any password).

## Building and Running
```
make build
chmod +x ./server
./server
```

## Endpoints

/api/message - GET

/api/message - POST

## Testing
`go test`

`go run .`

```
# see that you need to pass basic auth
curl localhost:8080/api/message

# see friendly message for if you don't have a message saved yet
curl -u george.bluth@reqres.in:password localhost:8080/api/message

# set a message
curl -XPOST -u george.bluth@reqres.in:password localhost:8080/api/message -d '{"message": "hello, world"}'

# retrieve the message you set
curl -u george.bluth@reqres.in:password localhost:8080/api/message
```