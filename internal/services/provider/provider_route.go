package provider

import (
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
