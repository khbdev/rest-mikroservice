package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// --- PRODUCT MODEL ---
type Product struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
	Image string `json:"image"`
}

var Products []Product

// --- STORAGE FUNKSIYALARI ---
func AddProduct(p Product) {
	Products = append(Products, p)
}

func GetProducts() []Product {
	return Products
}

// --- HANDLERLAR ---
func GetProductsHandler(c *gin.Context) {
	products := GetProducts()
	c.JSON(http.StatusOK, products)
}

func CreateProductHandler(c *gin.Context) {
	var newProduct Product
	if err := c.ShouldBindJSON(&newProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	AddProduct(newProduct)
	c.JSON(http.StatusOK, newProduct)
}

// --- MAIN ---
func main() {
	r := gin.Default()

	r.GET("/products", GetProductsHandler)
	r.POST("/products", CreateProductHandler)

	// product servisini 8082 portda ishga tushuramiz
	r.Run(":8084")
}
