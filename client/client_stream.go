package main

import (
	"context"
	"log"
	"time"

	pb "github.com/muaj07/grpc-tutorial/invoicer"
)

func callCreateClientStreaming(client pb.InvoicerClient, names *pb.NameList) {
	log.Printf("Client streaming Side")
	stream, err := client.CreateClientStreaming(context.Background())
	if err!=nil{
		log.Fatalf("Could not send names: %v", err)
	}
	for _, name := range *&names.Names {
       req := &pb.CreateRequest{
		Amount: nil,
		From: "Alice",
		To: "Bob",
	}
	if err := stream.Send(req); err!=nil{
        log.Fatalf("Could not send names: %v", err)
    }
	log.Printf("Send the request with name: %s", name)
	time.Sleep(2*time.Second)
	}
	res, err := stream.CloseAndRecv()
	log.Printf("Client streaming finished")
	if err!= nil{
		log.Fatalf("Error while receiving %v", err)
	}
	log.Printf("%v", res.Message)
    }
