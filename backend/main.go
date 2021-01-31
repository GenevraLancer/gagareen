package main

import (
	"io/ioutil"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"

	"github.com/GenevraLancer/gagarin/backend/gateway"
	myapi "github.com/GenevraLancer/gagarin/backend/gen/api"
	handlers "github.com/GenevraLancer/gagarin/backend/handlers"

	//Static files
	_ "github.com/GenevraLancer/gagarin/backend/statik"
)

func main() {
	//Adds gRPC internal logs
	log := grpclog.NewLoggerV2(os.Stdout, ioutil.Discard, ioutil.Discard)
	grpclog.SetLoggerV2(log)

	//Specify the port to listen for client requests
	addr := ":9244"
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	//Create an instance of the gRPC server
	// TODO: Replace with your own certificate!
	creds, err := credentials.NewServerTLSFromFile("cert/server.crt", "cert/server.key")
	if err != nil {
		log.Fatalln("Failed to load a self-signed certtificate server.crt:", err)
	}

	s := grpc.NewServer(grpc.Creds(creds))

	myapi.RegisterCommonServiceServer(s, handlers.New())

	// Serve gRPC Server
	log.Info("Serving gRPC on https://", addr)
	go func() {
		log.Fatal(s.Serve(lis))
	}()

	err = gateway.Run("dns:///" + addr)
	log.Fatalln(err)
}
