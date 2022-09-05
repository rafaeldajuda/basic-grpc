# GO-GRPC

Aplicação básica feita em gRPC.

link do tutorial: https://www.youtube.com/watch?v=LuS59XHdKG8&t=388s

## Estrutura

```txt
├── cmd
│   ├── client
│   │   └── main.go
│   └── server
│       └── main.go
├── pb
│   ├── message_service_grpc.pb.go
│   └── message_service.pb.go
├── proto
│   └── message_service.proto
├── go.mod
├── go.sum
└── README.md
```

## Pasta proto

Aqui é onde fica o arquivo de definição do *gRPC*. O arquivo define como vai ser os tipos de requisições e respostas, e também quais métodos irão existir no serviço.

A partir desse arquivo são gerados mais dois arquivos que conterão as funções que utilizaremos no nosso *server* e *client*.

Comando para geração dos arquivos:
```cmd
protoc --go_out=. --go-grpc_out=. proto/*.proto
```

### Arquivo message_service.proto

**syntax:** seta qual sintaxe iremos utilizar no arquivo *.proto*.

```proto3
syntax = "proto3";
```

**option go_package:** aqui fica o caminho dos arquivos que serão gerados pelo arquivo *.proto*.

```proto3
option go_package = "/pb";
```

**message:** define como serão as entradas e saidas das funções.

```proto3
message Request {
    string Message = 1;
}

message Response {
    int32 Status = 1;
}
```

**service:** define quais as funções existirão no serviço. As entradas e saídas das funções devem utilizar os **messages**.

```proto3
service SendMessage {
    rpc RequestMessage (Request) returns (Response){}
}
```
## Pasta pb

Aqui fica os arquivos gerados pelo arquivo *.proto*.

## Pasta cmd/server

Pasta do nosso server que executará as funções do arquivo *.proto*.

```go
func main() {

	// create grpc server
	grpcServer := grpc.NewServer()

	// register server
	pb.RegisterSendMessageServer(grpcServer, &Server{})

	// open tcp connection
	port := ":5000"
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}

	// start grpc server
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}

}
```

**pb.RegisterSendMessageServer():** função que registra as funções do nosso serviço no server. Sem registrar as funções não será possível chamar as funções do nosso serviço.

```go
pb.RegisterSendMessageServer(grpcServer, &Server{})
```

**struct Server:** estrutura de um serviço a ser implementado. A estrutura Server possui a estrutura **pb.UnimplementedSendMessageServer** que foi gerado pelo arquivo *.proto*. 

```go
type Server struct {
	pb.UnimplementedSendMessageServer
}
```

A partir dessa estrutura podemos utilizar e "sobrescrever" suas funções.

```go
func (UnimplementedSendMessageServer) RequestMessage(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestMessage not implemented")
}
func (UnimplementedSendMessageServer) mustEmbedUnimplementedSendMessageServer() {}
```

```go
func (service *Server) RequestMessage(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	log.Print("mensagem recebida: ", req.GetMessage())

	response := &pb.Response{
		Status: 1,
	}

	return response, nil
}
func (service *Server) mustEmbedUnimplementedSendMessageServer() {}
```

## Pasta cmd/client

Pasta do nosso client que consumirá o nosso serviço.

**grpc.Dial()**: função que conecta ao server.

```go
conn, err := grpc.Dial("localhost:5000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
```

O segundo parâmetro dessa função deve passar uma forma de autenticação. Caso não tenha denhema forma de autencação podemos utilizar o seguinte trecho para podemos testar.

```go
grpc.WithTransportCredentials(insecure.NewCredentials())
```

**pb.NewSendMessageClient():** registra o nosso client no serviço. Sem isso não consiguiremos consumir os serviços. Devemos passar a conexão como parâmetro.

```go
client := pb.NewSendMessageClient(conn)
```

**client.RequestMessage():** chama a função do serviço.

```go 
res, err := client.RequestMessage(context.Background(), req)
if err != nil {
	log.Fatal(err)
}
```
