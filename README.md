# GRPC Proof Of Concept w/ GoKit - Lorem Ipsum Generator

This is proof of concept for a gRPC microserve using GoKit that generates lorem ipsum text between a minimum and maximum number of characters. Can generate words, sentences or paragraphs with the cmd flags on the client.

Format: [client Directory] 
```sh
go run main.go lorem (word|sentence|paragraph) (minChars) (maxChars)

```

### Installation
Step One:
```sh
Clone this repo
$ git clone https://github.com/JohnAntonusMaximus/gokit-grpc-http-microservice.git
```

Step Two:
```sh
$ go mod download
```
Note: If you receive errors with finding commit hashes, just run the command again. 

### gRPC Usage
Step One:
```sh
$ go run $GOPATH/src/github.com/johnantonusmaximus/gokit-grpc-http-microservice/server/main.go
```

Step Two:
Open another terminal window:
```sh
$ go run $GOPATH/src/github.com/johnantonusmaximus/gokit-grpc-http-microservice/client/main.go lorem sentence 10 1000
```

### HTTP Usage
Step One:
```sh
Make a POST Request to the following endpoint:
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
