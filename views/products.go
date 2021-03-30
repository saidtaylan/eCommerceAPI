package views

import (
	"context"
	"ecommerce/database"
	"ecommerce/helpers"
	"ecommerce/models"
	"encoding/json"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

func Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome"))
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	cur, err := database.ProdCol.Find(context.TODO(), bson.M{})
	helpers.PrintErr(err)
	var theProd models.Product
	var theSeller models.Seller
	var Result models.AllProducts
	encoder := json.NewEncoder(w)
	for cur.Next(context.TODO()) {
		cur.Decode(&theProd)
		helpers.PrintErr(err)
		err = database.SellerCol.FindOne(context.TODO(), bson.M{"_id": theProd.Seller}).Decode(&theSeller)
		Result.Count = theProd.Count
		Result.Desc = theProd.Desc
		Result.Price = theProd.Price
		Result.Title = theProd.Title
		Result.Seller.Email = theSeller.Email
		Result.Seller.Name = theSeller.Name
		Result.Seller.LastName = theSeller.LastName
		helpers.PrintErr(err)
		w.WriteHeader(200)
		encoder.Encode(Result)
	}
	helpers.JSON(w, 200, "Successful")
	return
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	if w.Header().Get("Content-Type") == "application/json" {
		var theProd models.Product
		err := json.NewDecoder(r.Body).Decode(&theProd)
		fmt.Fprintln(w, theProd)
		helpers.PrintErr(err)
		result, err := database.ProdCol.InsertOne(context.TODO(), theProd)
		helpers.PrintErr(err)
		fmt.Fprintln(w, result)
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(result)
		return
	}
	w.WriteHeader(415)
	w.Write([]byte("Content-Type mismatched. Content-Type should be application/json"))
	return
}

func CreateSeller(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") == "application/json" {
		w.Header().Set("Content-Type", "application/json")
		var seller models.Seller
		err := json.NewDecoder(r.Body).Decode(&seller)
		helpers.PrintErr(err)
		result, err := database.SellerCol.InsertOne(context.Background(), seller)
		fmt.Fprintln(w, result)
		if result == nil {
			w.WriteHeader(500)
			w.Write([]byte("Seller could not created"))
			return
		}
	}
	w.WriteHeader(415)
	w.Write([]byte("Content type mismatched. Content-Type should be application/json"))
	return
}

func GetSellers(w http.ResponseWriter, r *http.Request) {
	cur, _ := database.SellerCol.Find(context.TODO(), bson.M{})
	var theSeller models.AllSellers
	encoder := json.NewEncoder(w)
	//helpers.PrintErr(err)
	w.WriteHeader(200)
	for cur.Next(context.Background()) {
		err := cur.Decode(&theSeller)
		helpers.PrintErr(err)
		encoder.Encode(theSeller)
	}
}
