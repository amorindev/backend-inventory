package company

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)


// apply status not found
func GetCompanyByIdHandler(c *gin.Context) {
	companyIDParam := c.Param("id")
	companyID, err := strconv.ParseInt(companyIDParam,10,64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request Company - Invalid ID"})
		return
	}

	company, err := GetCompanyByIdDto(companyID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request Company - Fetching company"})
		return
	}
	c.JSON(http.StatusOK, company)
}
