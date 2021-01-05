# GRPC Proof Of Concept w/ GoKit - Lorem Ipsum Generator

This is proof of concept for a gRPC microserve using GoKit that generates lorem ipsum text between a minimum and maximum number of characters. Can generate words, sentences or paragraphs with the cmd flags on the client.

### Installation
Step #1: Clone this repo...
```sh
$ git clone https://github.com/kaizenlabs/gokit-grpc-http-microservice.git
```

Step #2: Install dependencies...
```sh
$ cd $GOPATH/src/github.com/kaizenlabs/gokit-grpc-http-microservice
$ go mod download
$ go mod vendor
```

### gRPC Usage
Step #1: Start the server...
```sh
$ go run $GOPATH/src/github.com/kaizenlabs/gokit-grpc-http-microservice/server/main.go
```

Step #2: Open another terminal window and run...
```sh
$ go run $GOPATH/src/github.com/kaizenlabs/gokit-grpc-http-microservice/client/main.go lorem sentence 10 1000
```

### HTTP Usage
Step #1: While running the server, make a POST request to the following endpoint w/ Postman or curl:
```sh
localhost:8080/v1/lorem/{requestType}/{min}/{max}
```

WHERE:
  - requestType = (word|sentence|paragraph)
  - min = some positive integer
  - max = some positive integer

Example:
```sh
$ POST localhost:8080/v1/lorem/sentence/10/1000
```

For gRPC, use this format:
```sh
go run main.go lorem (word|sentence|paragraph) (minChars) (maxChars)

```