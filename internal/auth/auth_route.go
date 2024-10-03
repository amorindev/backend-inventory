package auth

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func PostLoginHandler(c *gin.Context) {
	var user UserEntity

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Post user bad request deserializar"})
		return
	}

	//validate
	validateNewProduct := validator.New()
	err := validateNewProduct.Struct(user)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		msg := fmt.Sprintf("message post user: %v", errors)
		c.JSON(http.StatusBadRequest, gin.H{"message": msg})
		return
	}

	if !(user.Email == "calidad@gmail.com" && user.Password == "Calidad2024") {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials"})
		return
	}
	c.JSON(http.StatusOK, nil)
}
