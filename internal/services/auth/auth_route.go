package auth

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
)


func PostLoginHandler(c *gin.Context) {
	var user UserEntity

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Post user bad request deserializar"})
		return
	}

	//validate
	_, err := ValidateUserAuth(user)
	if err != nil {
		msg := fmt.Sprintf("message post user: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": msg})
		return
	}


	c.JSON(http.StatusOK, nil)
}

func ValidateUserAuth(u UserEntity) (isValid bool, err error){
	if u.Email == "" || u.Password == "" {
		return false, fmt.Errorf("email or password connot be empty")
	}
	
	regexEmail :=  `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(regexEmail)
	
	if !re.MatchString(u.Email) {
		return false, fmt.Errorf("invalid email")
	}

	if(u.Email != "calidad@gmail.com" && u.Password != "Calidad2024"){
		return false, fmt.Errorf("invalid credentials")
	}
	return true, nil
}
