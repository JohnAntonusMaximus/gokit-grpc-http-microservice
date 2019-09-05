# GRPC Proof Of Concept w/ GoKit - Lorem Ipsum Generator

This is proof of concept for a gRPC microserve using GoKit that generates lorem ipsum text between a minimum and maximum number of characters. Can generate words, sentences or paragraphs with the cmd flags on the client.

Format: [client Directory] 
```sh
go run main.go lorem (word|sentence|paragraph) (minChars) (maxChars)

```

### Installation
Step One:
```sh
$ go get github.com/johnantonusmaximus/gokit-grpc-poc
```

Step Two:
```sh
$ go get -u github.com/kardianos/govendor
$ cd $GOPATH/src/github.com/kardianos/govendor
$ go install .
```

Step Three:
```sh
$ cd $GOPATH/src/github.com/johnantonusmaximus/gokit-grpc-poc
$ govendor init 
```
Note: Make sure your the path to GOBIN is in your PATH variable

Step Four:
```sh
$ govendor sync
```

### gRPC Usage
Step One:
```sh
$ go run $GOPATH/src/github.com/johnantonusmaximus/gokit-grpc-poc/server/main.go
```

Step Two:
```sh
$ go run $GOPATH/src/github.com/johnantonusmaximus/gokit-grpc-poc/client/main.go lorem paragraph 10 4000
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
$ POST localhost:8080/v1/lorem/sentence/10/100
```
