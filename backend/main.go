package main

import (
	"io/ioutil"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	myapi "github.com/GenevraLancer/gagarin/backend/gen/api"
	queryhandler "github.com/GenevraLancer/gagarin/backend/queryhandler"

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
	myapi.RegisterFieldServiceServer(s, queryhandler.New())

	// Serve gRPC Server
	log.Info("Serving gRPC on https://", addr)
	go func() {
		log.Fatal(s.Serve(lis))
	}()
}
