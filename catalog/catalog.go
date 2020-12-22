package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

type Product struct {
	Uuid    string  `json:"uuid"`
	Product string  `json:"product"`
	Price   float64 `json:"price,string"`
}

type Products struct {
	Products []Product `json:"products"`
}

var productsUrl string

func loadProducts() []Product {
	response, err := http.Get(productsUrl + "/products")
	if err != nil {
		fmt.Println("Error to get products")
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error to get products")
	}

	var products Products
	json.Unmarshal(body, &products)

	return products.Products
}

func ListCatalog(w http.ResponseWriter, r *http.Request) {
	products := loadProducts()

	t := template.Must(template.ParseFiles("template/catalog.html"))
	t.Execute(w, products)
}

func ListProduct(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	request, err := http.Get(productsUrl + "/product?id=" + id)
	if err != nil {
		fmt.Println("Erro to get product")
	}

	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		fmt.Println("Erro to get product")
	}

	var product Product
	json.Unmarshal([]byte(body), &product)

	t := template.Must(template.ParseFiles("template/view.html"))

	t.Execute(w, product)
}

func init() {
	// productsUrl = os.Getenv("PRODUCT_URL")
	productsUrl = "http://localhost:8081"
}

func main() {
	r := http.NewServeMux()

	r.HandleFunc("/catalog", ListCatalog)
	r.HandleFunc("/catalog-show", ListProduct)

	http.ListenAndServe(":9091", r)
}
