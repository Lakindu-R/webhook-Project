package config
import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
var DB *mongo.Database

func ConnectDB(){
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil{
		log.Fatal("Error creating MongoDB client: ", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil{
		log.Fatal("Error connecting to MongoDB: ", err)
	}
	DB = client.Database(os.Getenv("DB_NAME"))
	log.Println("Connected to MongoDB")
}