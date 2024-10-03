package kardex

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// PostKardexHandler godoc
// @Summary      Create kardex
// @Description  Create a new kardex
// @Tags         kardex
// @Accept       json
// @Produce      json
// @Param        kardex body Kardex true "Kardex to be created"
// @Success      200 {object} Kardex
// @Router       /kardex [post]
func GetKardexHandler(c *gin.Context) {
	kardexes, err := GetKardexDto()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error retrieving Kardex records"})
		return
	}
	c.JSON(http.StatusOK, kardexes)
}

// GetKardexHandler godoc
// @Summary      Get all kardex entries
// @Description  Return the list of kardex entries
// @Tags         kardex
// @Produce      json
// @Success      200 {array} Kardex
// @Router       /kardex [get]
func PostKardexHandler(c *gin.Context) {
	var kardex Kardex

	if err := c.ShouldBindJSON(&kardex); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Post products bad request deserializar"})
		return
	}


	validateNewProduct := validator.New()
	err := validateNewProduct.Struct(kardex)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		msg := fmt.Sprintf("messsage post user: %v", errors)
		c.JSON(http.StatusBadRequest, gin.H{"message": msg})
		return
	}

	if(kardex.Type != "SALIDA" && kardex.Type != "ENTRADA"){
		c.JSON(http.StatusBadRequest, gin.H{"message": "Please insert vaid type (ENTRADA/SALIDA)"})
		return
	}

	k,kp,err := CreateKardexDto(kardex)
	if err != nil {
		msg := fmt.Sprintf("Post kardex bad request testt: %v", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": msg})
		return
	} 
	
	if(len(kp)>0) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Some products have insufficient stock",
			"insufficient_stock_products": kp,
		})	
		return
	}

	c.JSON(http.StatusOK,k)
}
