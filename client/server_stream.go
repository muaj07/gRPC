package main

import (
	"context"
	"log"
	"io"
	pb "github.com/muaj07/grpc-tutorial/invoicer"
)

func callCreateServerStreaming(client pb.InvoicerClient, names *pb.NameList){
	log.Printf("Streaming Started")
	stream, err := client.CreateServerStreaming(context.Background(), names)
	if err!=nil{
		log.Fatalf("Could not send names: %v",err)
	}
	for{
		message, err := stream.Recv()
		if err == io.EOF{
			break
		}
		if err!=nil{
			log.Fatalf("Could not receive message: %v",err)
		}
		log.Printf("Received message: %v",message)
	}
	log.Printf("Streaming Finished")
}