package main

import(
	pb "github.com/muaj07/grpc-tutorial/invoicer"
	context "context"
)

func(m *myInvoicerServer) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error){

	return &pb.CreateResponse{
		Message: "Hello!",
	}, nil
}


