package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)
type Product struct {
	Name string `json:"name"`
}
type UpdateProducts struct {
	Name string `json:"name"`
	NewName string `json:"newname"`
}
var products = []string{"banana", "apple", "cherry"}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/products", giveAllProductsHandler).Methods("GET")
	r.HandleFunc("/products",createNewProductHandler).Methods("POST")
	r.HandleFunc("/products",removeProductHandler).Methods("DELETE")
		r.HandleFunc("/products",updateHandler).Methods("PUT")

	http.ListenAndServe(":8080",r)
	fmt.Println("server is running on port 8080")
}
func giveAllProductsHandler(w http.ResponseWriter, r *http.Request){
	data := giveAllProducts()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
	fmt.Println("succesfully sent products")
}
func giveAllProducts() []string{
	return products
}
func createNewProductHandler(w http.ResponseWriter, r *http.Request){
	var p Product
	json.NewDecoder(r.Body).Decode(&p)
	product := createNewProduct(p.Name)
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(product)
}
func createNewProduct(product string) string {
	products = append(products, product)
	return product
}
func removeProductHandler(w http.ResponseWriter, r *http.Request){
	var p Product
	json.NewDecoder(r.Body).Decode(&p)
	newProductSlice := removeProduct(p.Name)
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(newProductSlice)
}
func removeProduct(product string) []string {
	for i,p := range products {
		if p == product {
			products = append(products[:i], products[i+1:]... )
			break
		}
	}
	return products
}
func updateHandler(w http.ResponseWriter, r *http.Request){
	var p UpdateProducts
	json.NewDecoder(r.Body).Decode(&p)
	newProducts := updateProduct(p.Name, p.NewName)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newProducts)
}

func updateProduct(product string, newName string) []string {
	for i,p := range products {
		if p == product {
			products[i] = newName 
			break
		}
	}
	return products
}