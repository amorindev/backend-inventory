package kardex

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func GetKardexHandler(c *gin.Context) {
	kardexes, err := GetKardexDto()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error retrieving Kardex records"})
		return
	}
	c.JSON(http.StatusOK, kardexes)
}

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
