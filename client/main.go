package main

import (
	"log"
	pb "github.com/muaj07/grpc-tutorial/invoicer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)
const(
	port= ":8089"
)
func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err!= nil{
		log.Fatalf("Client failed to connect %v", err)
	}
	defer conn.Close()
	client := pb.NewInvoicerClient(conn)
	names := &pb.NameList{
		Names: []string{"Alice", "Bob", "Eva"},
	}
	//callCreateServerStreaming(client, names)
	//callCreateClientStreaming(client, names)
	callCreateBidirectionalStreaming(client, names)
}

