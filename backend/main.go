package main

import (
	"io/ioutil"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	"github.com/GenevraLancer/gagarin/backend/gateway"
	pbField "github.com/GenevraLancer/gagarin/backend/myapi"
	"github.com/GenevraLancer/gagarin/backend/server"

	//Static files
	_ "github.com/GenevraLancer/gagarin/backend/statik"
)

func main() {
	//Adds gRPC internal logs
	log := grpclog.NewLoggerV2(os.Stdout, ioutil.Discard, ioutil.Discard)
	grpclog.SetLoggerV2(log)

	//Specify the port to listen for client requests
	addr := "localhost:82000"
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	//Create an instance of the gRPC server
	s := grpc.NewServer(
	// TODO: Replace with your own certificate!
	//grpc.Creds(credentials.NewServerTLSFromCert(&insecure.Cert)),
	)
	pbField.RegisterFieldServiceServer(s, server.New())

	// Serve gRPC Server
	log.Info("Serving gRPC on https://", addr)
	go func() {
		log.Fatal(s.Serve(lis))
	}()

	err = gateway.Run("dns:///" + addr)
	log.Fatalln(err)
}
