package main

import (
	userpb "github.com/dmytrozilnyk/communication/grpc/gen/go/proto/user/v1"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"

	wearablepb "github.com/dmytrozilnyk/communication/grpc/gen/go/proto/wearable/v1"
)

func main() {
	lis, err := net.Listen("tcp", "localhost:9879")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	wearableServer := &wearableService{}

	healthServer := health.NewServer()

	go func() {
		for {
			status := healthpb.HealthCheckResponse_SERVING
			// Check if user Service is valid
			if time.Now().Second()%2 == 0 {
				status = healthpb.HealthCheckResponse_NOT_SERVING
			}

			healthServer.SetServingStatus(wearablepb.WearableService_ServiceDesc.ServiceName, status)

			time.Sleep(1 * time.Second)
		}
	}()

	healthServer.SetServingStatus("", healthpb.HealthCheckResponse_SERVING)
	healthServer.SetServingStatus(userpb.UserService_ServiceDesc.ServiceName, healthpb.HealthCheckResponse_SERVING)

	healthpb.RegisterHealthServer(grpcServer, healthServer)
	wearablepb.RegisterWearableServiceServer(grpcServer, wearableServer)

	grpcServer.Serve(lis)
}
