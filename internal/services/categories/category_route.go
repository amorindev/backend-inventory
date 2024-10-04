package categories

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetCategories godoc
// @Summary      Get all categories
// @Description  Return the list of categories
// @Tags         categories
// @Produce      json
// @Success      200 {array} Category
// @Router       /categories [get]
func GetCategoriesHandler(c *gin.Context) {
	ctgs, err := GetCategoriesDto()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request categories"})
		return
	}
	c.JSON(http.StatusOK, ctgs)
}

