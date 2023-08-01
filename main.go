package main

import (
	"context"
	"log"
	"github.com/dzoniops/notification-service/db"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	//uri := fmt.Sprintf("mongodb+srv://%s:%s@cluster06443.hcrrodp.mongodb.net/?retryWrites=true&w=majority", os.Getenv("MONGO_USERNAME"), os.Getenv("MONGO_PASSWORD"))
	uri := "mongodb://localhost:27017"	
	// Create a new client and connect to the server
	client, err := db.InitDb(uri) 
	if err != nil {
		log.Fatal("Failed to connect to Mongodb:",err)	
	}

	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	
	db.InsertData(client)
	// // Send a ping to confirm a successful connection
	// if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Err(); err != nil {
	// 	panic(err)
	// }
	// fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
}
