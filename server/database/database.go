package database

import (
	"context"
	"crypto/tls"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	// "gorm.io/gorm"
)

// Connections
var DBConn *mongo.Client
var Blog_Coll_Conn *mongo.Collection
var Song_Coll_Conn *mongo.Collection
var SongCover_Coll_Conn *gridfs.Bucket

func ConnectDB() {
	// Access MongoDB credentials from environment
	username := os.Getenv("mongodb_username")
	password := os.Getenv("mongodb_password")

	url := "mongodb+srv://" + username + ":" + password + "@cluster0.9th4r7y.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"
	clientOptions := options.Client().ApplyURI(url)

	// Configuring TLS
	clientOptions.SetTLSConfig(&tls.Config{
		InsecureSkipVerify: false,
	})

	// Connecting the target MongDB
	mongoTestClient, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		log.Fatal("error while connecting mongodb", err)
	}
	log.Println("Mongodb successfully connected.")

	// Ping the target MongoDB
	err = mongoTestClient.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("Ping failed", err)
	}
	log.Println("Ping success")

	// Output the connection
	DBConn = mongoTestClient
	Blog_Coll_Conn = DBConn.Database("maindb").Collection("Blog")
	Song_Coll_Conn = DBConn.Database("maindb").Collection("Song")
	bucket, err := gridfs.NewBucket(DBConn.Database("maindb"))

	if err != nil {
		log.Fatal("Some problems on grid", err)
	}

	SongCover_Coll_Conn = bucket

}
