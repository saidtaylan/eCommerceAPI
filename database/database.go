package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database
var ProdCol *mongo.Collection
var SellerCol *mongo.Collection

func ConnectToDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	options := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(ctx, options)
	if err != nil {
		panic(err)
	}
	DB = client.Database("project")
	if DB == nil {
		panic(DB)
	}
	ProdCol = DB.Collection("Products")
	SellerCol = DB.Collection("Sellers")
}
