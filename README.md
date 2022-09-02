# gRPC

1 - install Golang: https://go.dev/doc/install

2 - install Protocol Buffer Compiler: https://grpc.io/docs/protoc-installation/

    2.1 - sudo apt-get update
    2.2 - sudo apt-get install -y protobuf-compiler
    2.3 - protoc --version (Ensure compiler version is 3+)

3 - Go plugins for the protocol compiler: https://grpc.io/docs/languages/go/quickstart/

    3.1 - Install the protocol compiler plugins for Go using the following commands

        $ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
        $ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

    3.2 - Update your PATH so that the protoc compiler can find the plugins

        export PATH="$PATH:$(go env GOPATH)/bin"

4 - Get the example code

    4.1 - git clone -b v1.48.0 --depth 1 https://github.com/grpc/grpc-go
    4.2 - cd grpc-go/examples/helloworld

5 - Run the example

    5.1 - Compile and execute the server code --> go run greeter_server/main.go
    5.2 - From a different terminal, compile and execute the client code to see the client output: 
        --> go run greeter_client/main.go
    5.3 - Congratulations! You’ve just run a client-server application with gRPC

# Update the gRPC service

1 - Add the new method *SayHelloAgain()* in the file *helloworld/helloworld.proto* in the method *service Greeter*

    // The greeting service definition.
    service Greeter {
        // Sends a greeting
        rpc SayHello (HelloRequest) returns (HelloReply) {}
        // Sends another greeting
        rpc SayHelloAgain (HelloRequest) returns (HelloReply) {}
    }

2 - Regenerate gRPC code

    protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    helloworld/helloworld.proto

This will regenerate the helloworld/helloworld.pb.go and helloworld/helloworld_grpc.pb.go files, which contain:

* Code for populating, serializing, and retrieving HelloRequest and HelloReply message types.
* Generated client and server code.

3 - Update the server - Open greeter_server/main.go

    func (s *server) SayHelloAgain(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
        return &pb.HelloReply{Message: "Hello again " + in.GetName()}, nil
    }

4 - Update the client - greeter_client/main.go

```go
r, err = c.SayHelloAgain(ctx, &pb.HelloRequest{Name: *name})
if err != nil {
        log.Fatalf("could not greet: %v", err)
}
log.Printf("Greeting: %s", r.GetMessage())
```

5 - Run

    5.1 - Run the server
        go run greeter_server/main.go
    5.2 - From another terminal, run the client. This time, add a name as a command-line argument:
        go run greeter_client/main.go
        go run greeter_client/main.go --name=Rafael
