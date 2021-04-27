package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var e error

type Product struct {
	ID      uint   `json:"id"`
	Product string `json:"product"`
	Ptype   string `json:"ptype"`
	Price   int    `json:"price"`
	Os      string `json:"os"`
}

func main() {
	db, e = gorm.Open("sqlite3", "./qh.db")
	if e != nil {
		fmt.Println(e)
	}
	defer db.Close()

	db.AutoMigrate(&Product{})

	r := gin.Default()
	// Get products
	r.GET("/products", getProducts)
	// Get products by id
	r.GET("/products/:id", getProductById)
	// Insert new products
	r.POST("/products", insertProduct)
	// Update products
	r.PUT("/products/:id", updateProduct)
	// Delete products
	r.DELETE("/products/:id", deleteProduct)
	r.Run(":1991")
}

// Get products
func getProducts(c *gin.Context) {
	var products []Product
	if e := db.Find(&products).Error; e != nil {
		c.AbortWithStatus(404)
		fmt.Println(e)
	} else {
		c.JSON(200, products)
	}
}

// Get products by id
func getProductById(c *gin.Context) {
	var product Product
	id := c.Params.ByName("id")
	if e := db.Where("id = ?", id).First(&product).Error; e != nil {
		c.AbortWithStatus(404)
		fmt.Println(e)
	} else {
		c.JSON(200, product)
	}
}

// Insert new product
func insertProduct(c *gin.Context) {
	var product Product
	c.BindJSON(&product)
	db.Create(&product)
	c.JSON(200, product)
}

// Update product
func updateProduct(c *gin.Context) {
	var product Product
	id := c.Params.ByName("id")
	if e := db.Where("id = ?", id).First(&product).Error; e != nil {
		c.AbortWithStatus(404)
		fmt.Println(e)
	} else {
		c.BindJSON(&product)
		db.Save(&product)
		c.JSON(200, product)
	}
}

// Delete product
func deleteProduct(c *gin.Context) {
	var product Product
	id := c.Params.ByName("id")
	d := db.Where("id = ?", id).Delete(&product)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
