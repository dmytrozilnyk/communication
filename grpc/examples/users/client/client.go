package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"

	userpb "github.com/dmytrozilnyk/communication/grpc/gen/go/proto/user/v1"
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

	client := userpb.NewUserServiceClient(conn)
	response, err := client.GetUser(context.Background(), &userpb.GetUserRequest{Uuid: "124Dimi"})
	if err != nil {
		log.Fatalln("Consuming stream", err)
	}

	fmt.Println(response.GetUser())
}
