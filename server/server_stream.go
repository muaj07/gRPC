package main

import (
	"log"
	"time"

	pb "github.com/muaj07/grpc-tutorial/invoicer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/profiling/service"
)

func (s *myInvoicerServer) CreateServerStreaming(req *pb.NameList, stream pb.Invoicer_CreateServerStreamingServer) error{
	log.Printf("Got request with names : %v", req.Names)
	for _, name := range req.Names{
		res := &pb.CreateResponse{
			Message: "Hello" + name,
		}
		if err := stream.Send(res); err!= nil {
		return err
	}
	time.Sleep(2 *time.Second)
}
return nil
}