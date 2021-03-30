package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID     primitive.ObjectID `bson:"_id"`
	Title  string             `bson:"title,omitempty"`
	Desc   string             `bson:"desc,omitempty"`
	Price  string             `bson:"price,omitempty"`
	Count  uint               `bson:"count,omitempty"`
	Seller primitive.ObjectID `bson:"seller_id"`
}

type Seller struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Name     string             `bson:"name,emitempty"`
	LastName string             `bson:"lastname,omitempty"`
	Email    string             `bson:"email,omitempty"`
	Password string             `bson:"password,omitempty"`
}

type AllProducts struct {
	Title  string `json:"title,omitempty"`
	Desc   string `json:"desc,omitempty"`
	Price  string `json:"price,omitempty"`
	Count  uint   `json:"count,omitempty"`
	Seller struct {
		Name     string `json:"name,emitempty"`
		LastName string `json:"lastname,omitempty"`
		Email    string `json:"email,omitempty"`
	} `json:"seller"`
}

type AllSellers struct {
	Name     string `bson:"name,emitempty"`
	LastName string `bson:"lastname,omitempty"`
	Email    string `bson:"email,omitempty"`
}
