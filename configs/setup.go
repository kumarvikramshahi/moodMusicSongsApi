package configs

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"time"

// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// func ConnectDB() *mongo.Client {
// 	client, err := mongo.NewClient(options.Client().ApplyURI(EnvMongoURI()))
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
// 	err = client.Connect(ctx)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	//ping the database
// 	err = client.Ping(ctx, nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	} else {
// 		fmt.Println("Connected to MongoDB")
// 	}
// 	fmt.Println(client)
// 	return client
// }

//Client instance
// var DB *mongo.Client = ConnectDB()

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var DefaultDB *mongo.Database

func DatabaseConnection() {
	// env variable retrieval
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable.")
	}

	// connection to DB
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	// err = client.Connect(ctx)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer client.Disconnect(ctx)

	// ping DB
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection made with mongoDB")

	// Global variable of DB
	DefaultDB = client.Database("forcucekians")
}

//getting database collections
func GetCollection(collectionName string) *mongo.Collection {
	log.Println(DefaultDB)
	if DefaultDB != nil {
		return DefaultDB.Collection(collectionName)
	} else {
		log.Fatal("not getting Db name")
	}
	return nil
}
