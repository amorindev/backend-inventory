package main

import (
	"log"
	"net/http"
	"os"

	// project

	//router
	"github.com/amorindev/backend-inventory/db"
	"github.com/amorindev/backend-inventory/internal/auth"
	"github.com/amorindev/backend-inventory/internal/categories"
	"github.com/amorindev/backend-inventory/internal/kardex"
	"github.com/amorindev/backend-inventory/internal/product"
	"github.com/gin-gonic/gin"

	// enviroment variables
	"github.com/joho/godotenv"

	// controller for database
	_ "github.com/lib/pq"

	//swagger
	_ "github.com/amorindev/backend-inventory/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Inventory api
// @version 1.0
// @description Api for an app of inventory
// @host localhost:7000
// @BasePath /v1
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("PORT")
	ginMode := os.Getenv("GIN_MODE")
	mode := os.Getenv("BACK_ENV")

	if port == "" || ginMode == "" || mode == ""{
		log.Fatal("one or various enviroment variable is not set on main.go")
	}

	db.DBConnection()

	// that you need to set this before creating the router
	gin.SetMode(ginMode) 

	r := gin.Default()

	r.Use(func(c *gin.Context) {
		
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, Origin")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		// Manejo de preflight request (OPTIONS)
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	})

	v1 := r.Group("/api/v1")
	{
		if mode != "prod" {
			v1.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		}
		v1.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "pong"})
		})

		// -------------------------  AUTH  --------------------------------
		//v1.GET("/products", product.GetProductsHandler)
		v1.POST("/authentication", auth.PostLoginHandler)

		// ------------------------- CATEGORIES ------------------------------
		v1.GET("/categories", categories.GetCategoriesHandler)

		// -------------------------  PRODUCTS  --------------------------------
		v1.GET("/products", product.GetProductsHandler)
		v1.POST("/products", product.PostProductHandler)
		v1.PUT("/products/:id", product.PutProductHandler)
		v1.DELETE("/products/:id", product.DeleteProductHandler)
	
		// -------------------------  KARDEX  --------------------------------
		v1.GET("/kardex", kardex.GetKardexHandler)
		v1.POST("/kardex", kardex.PostKardexHandler)
		//v1.PUT("/products/:id", products.PutProductHandler)
		//v1.DELETE("/products/:id", products.DeleteProductHandler)
	}

	r.Run(":"+port)

}
