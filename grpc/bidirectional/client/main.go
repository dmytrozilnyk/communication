package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	wearablepb "github.com/dmytrozilnyk/communication/grpc/gen/go/proto/wearable/v1"
)

func main() {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.Dial("localhost:9879", opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	client := wearablepb.NewWearableServiceClient(conn)
	stream, err := client.ConsumeBeatsPerMinute(context.Background())
	if err != nil {
		log.Fatalln("Consuming stream", err)
	}

	for i := 0; i < 5; i++ {
		err := stream.Send(&wearablepb.ConsumeBeatsPerMinuteRequest{Uuid: "Dimi", Value: 100, Minute: uint32(i)})
		if err != nil {
			log.Fatalln("Sending value", err)
		}
		time.Sleep(10 * time.Millisecond)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalln("Closing", err)
	}

	fmt.Println("Total messages", res.GetTotal())
}
