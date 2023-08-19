package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	pb "github.com/dzoniops/common/pkg/notification"
	"github.com/dzoniops/notification-service/db"
	"github.com/dzoniops/notification-service/services"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	//uri := fmt.Sprintf("mongodb+srv://%s:%s@cluster06443.hcrrodp.mongodb.net/?retryWrites=true&w=majority", os.Getenv("MONGO_USERNAME"), os.Getenv("MONGO_PASSWORD"))
	uri := "mongodb://localhost:27017"
	// Create a new client and connect to the server
	client, err := db.InitDb(uri)
	if err != nil {
		log.Fatal("Failed to connect to Mongodb:", err)
	}

	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("PORT")))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterNotificationServiceServer(s, &services.Server{
		DB: *client,
	})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
