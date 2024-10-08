package provider

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GETProvidersHandler(c *gin.Context) {
	prov, err := GetProviders()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request providers"})
		return
	}
	c.JSON(http.StatusOK, prov)
}

func POSTProviderHandler(c *gin.Context) {
	var newProvider ProviderEntity
	
	if err := c.ShouldBindJSON(&newProvider); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Post provider bad request deserializar"})
	}

	// validar 
	if newProvider.ProvName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Provider name must not be empty"})
		return
	}
	if newProvider.ProvAddress == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Provider address must not be empty"})
		return
	}
	if newProvider.ProvEmail == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Provider email must not be empty"})
		return
	}
	if newProvider.ProvPhone == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Provider phone must not be empty"})
		return
	}

	if newProvider.ComID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Provider companyID must not be empty"})
		return
	}

	// Registrar en la base de datos
	prov, err := CreateProvider(newProvider)
	if err != nil {
		msg := fmt.Sprintf("The provider could not be created in the database %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": msg})
		return
	}
	c.JSON(http.StatusOK, prov)
}