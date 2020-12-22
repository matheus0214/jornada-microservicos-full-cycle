package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Product struct {
	Uuid    string  `json:"uuid"`
	Product string  `json:"product"`
	Price   float64 `json:"price,string"`
}

type Products struct {
	Products []Product `json:"products"`
}

func loadData() []byte {
	jsonFile, err := os.Open("products.json")
	if err != nil {
		fmt.Println(err.Error())
	}

	defer jsonFile.Close()

	data, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err.Error())
	}

	return data
}

func ListProducts(w http.ResponseWriter, r *http.Request) {
	products := loadData()

	w.Write(products)
}

func GetProductById(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("id")

	data := loadData()
	var products Products

	json.Unmarshal(data, &products)

	for _, v := range products.Products {
		if v.Uuid == q {
			product, _ := json.Marshal(v)
			w.Write([]byte(product))
			return
		}
	}

	w.Write([]byte(`{ "message": "Product does not found" }`))
}

func main() {
	r := http.NewServeMux()

	r.HandleFunc("/products", ListProducts)
	r.HandleFunc("/product", GetProductById)

	http.ListenAndServe(":8081", r)
}
