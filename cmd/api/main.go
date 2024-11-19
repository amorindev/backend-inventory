package main

import (
	"log"

	"net/http"
	"os"

	// project

	//router

	"github.com/gin-gonic/gin"

	// enviroment variables
	"github.com/joho/godotenv"

	// controller for database
	_ "github.com/lib/pq"

	"github.com/amorindev/backend-inventory/internal/cors"
	"github.com/amorindev/backend-inventory/internal/db"
	"github.com/amorindev/backend-inventory/internal/services/auth"
	"github.com/amorindev/backend-inventory/internal/services/categories"
	"github.com/amorindev/backend-inventory/internal/services/company"
	"github.com/amorindev/backend-inventory/internal/services/kardex"
	"github.com/amorindev/backend-inventory/internal/services/product"
	"github.com/amorindev/backend-inventory/internal/services/provider"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("PORT")
	ginMode := os.Getenv("GIN_MODE")
	mode := os.Getenv("BACK_ENV")
	corsenv := os.Getenv("CORS")

	if port == "" || ginMode == "" || mode == ""{
		log.Fatal("one or various enviroment variable is not set on main.go")
	}

	db.DBConnection()

	// that you need to set this before creating the router
	gin.SetMode(ginMode) 

	r := gin.Default()

	if corsenv != "" {
		r.Use(cors.CorsMiddleware(corsenv))
	}

	v1 := r.Group("/api/v1")
	{

		v1.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "pong"})
		})

		// -------------------------  AUTH  --------------------------------
		v1.POST("/authentication", auth.PostLoginHandler)

		// ------------------------- CATEGORIES ------------------------------
		v1.GET("/companies/:id", company.GetCompanyByIdHandler)

		// ------------------------- CATEGORIES ------------------------------
		v1.GET("/categories", categories.GetCategoriesHandler)
		v1.POST("/categories", categories.PostCategoryHandler)

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

		// -------------------------  PROVIDERS  --------------------------------
		v1.GET("/providers", provider.GETProvidersHandler)
		v1.POST("/providers", provider.POSTProviderHandler)

	}

	r.Run(":"+port)
}
