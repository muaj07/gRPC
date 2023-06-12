package main

import(
	"context"
	"log"
	"io"
	"time"
	pb "github.com/muaj07/grpc-tutorial/invoicer"
)


func callCreateBidirectionalStreaming(client pb.InvoicerClient, names *pb.NameList) {

	stream, err := client.CreateBidirectionalStreaming(context.Background())
    if err!= nil {
        log.Fatalf("Client failed to create bidirectional streaming: %v", err)
    }
	waitc := make(chan struct{})
	go func() {
    	for {
        	message, err := stream.Recv()
        	if err == io.EOF {
            	break
        	}

        if err!= nil {
            log.Fatalf("Client failed to receive bidirectional streaming request: %v", err)
        }

		log.Println(message)
	}
        close(waitc)
    } ()


	for _, name :=range names.Names{
		req := &pb.CreateRequest{
		Amount: nil,
		From: "Alice",
		To: name,
		}
	if err := stream.Send(req); err !=nil{
		log.Fatalf("Client failed to send bidirectional streaming request: %v", err)
	}
	time.Sleep(2*time.Second)
}
stream.CloseSend()
<-waitc
log.Printf("Bidirectional streaming finish")
}



