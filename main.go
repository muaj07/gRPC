package main

import (
	"log"
	"net"
	"github.com/muaj07/grpc-tutorial/invoicer"
	"google.golang.org/grpc"
	"context"
)

type myInvoicerServer struct{
	invoicer.UnimplementedInvoicerServer
}



func(m myInvoicerServer) Create(ctx context.Context, req *invoicer.CreateRequest) (*invoicer.CreateResponse, error){

	return &invoicer.CreateResponse{
		Pdf: []byte(req.From),
		Docx:[]byte("Test1"),
	}, nil
}


func main() {
	lis, err := net.Listen("tcp", ":8089")
	if err!=nil {
		log.Fatalf("cannot create listener: %s", err)
	}
	serverRegistrar := grpc.NewServer()
	service := &myInvoicerServer{}
	invoicer.RegisterInvoicerServer(serverRegistrar, service) 
	err = serverRegistrar.Serve(lis)
	if err!=nil{
		log.Fatalf("cannot start server: %s", err)
	}

}