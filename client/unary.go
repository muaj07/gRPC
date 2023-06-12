package main

import (
	"context"
	"log"
	"time"
	pb "github.com/muaj07/grpc-tutorial/invoicer"
)

func callCreateRequest(client pb.InvoicerClient){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	req, err := client.Create(ctx, &pb.CreateRequest{
		Amount: nil,
		From: "Alice",
		To: "Bob",
	})
	if err!=nil{
		log.Fatalf("Could not send Create Message: %v", err)
	}
	log.Printf("Response: %v", req.Message)
}