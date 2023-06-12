package main

import (
	"log"
	"net"
	pb "github.com/muaj07/grpc-tutorial/invoicer"
	grpc "google.golang.org/grpc"
	
)
const(
	port = ":8089"
)

type myInvoicerServer struct{
	pb.UnimplementedInvoicerServer
}


func main() {
	lis, err := net.Listen("tcp", port)
	if err!= nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
	grpcServer := grpc.NewServer()
	service := &myInvoicerServer{}
	pb.RegisterInvoicerServer(grpcServer, service) 
	log.Printf("server started at %v", lis.Addr())
	err = grpcServer.Serve(lis)
	if err!=nil{
		log.Fatalf("cannot start server: %s", err)
	}
}