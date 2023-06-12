package main

import(
	pb "github.com/muaj07/grpc-tutorial/invoicer"
	"io"
	"log"
)


func (s *myInvoicerServer) CreateClientStreaming(stream pb.Invoicer_CreateClientStreamingServer) error {
	var messages []string
	for {
		req, err := stream.Recv()
        if err == io.EOF {
            return stream.SendAndClose(&pb.MessageList{Messages: massages})
        }
        if err!= nil {
            return err
        }
		log.Printf("Got request with name: %v", req.Name)
        messages = append(messages, "Hello!", req.Name)
	}
}