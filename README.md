# GRPC Proof Of Concept w/ GoKit - Lorem Ipsum Generator

This is proof of concept for a gRPC microserve using GoKit that generates lorem ipsum text between a minimum and maximum number of characters. Can generate words, sentences or paragraphs with the cmd flags on the client.

Format: [client Directory] 
```sh
go run main.go lorem (word|sentence|paragraph) (minChars) (maxChars)
```

### Usage
Step One:
```sh
$ go get github.com/johnantonusmaximus/gokit-grpc-poc
```

Step Two:
```sh
$ go run $GOPATH/src/github.com/johnantonusmaximus/gokit-grpc-poc/server/main.go
```

Step Three:
```sh
$ go run $GOPATH/src/github.com/johnantonusmaximus/gokit-grpc-poc/client/main.go lorem paragraph 10 4000
```

### To Do
  - Add HTTP REST Proxy in the proto file (commented out right now)