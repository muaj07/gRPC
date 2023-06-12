package main

import (
	"log"
	"io"
	pb "github.com/muaj07/grpc-tutorial/invoicer"
)

func (s *myInvoicerServer) CreateBidirectionalStreaming(stream pb.Invoicer_CreateBidirectionalStreamingServer) error{
	for {
        req, err := stream.Recv()
        if err == io.EOF {
            return nil
        }
        if err!= nil {
            return err
        }
        log.Printf("Got request with name: %v", req.Name)
		res := &pb.CreateResponse{
			Message: "Hello!"+ req.Name,
		}
		if err := stream.Send(res); err!=nil{
			return err
		}
    }
}