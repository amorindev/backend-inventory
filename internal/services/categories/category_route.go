package categories

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// GetCategories godoc
// @Summary      Get all categories
// @Description  Return the list of categories
// @Tags         categories
// @Produce      json
// @Success      200 {array} CategoryEntity
// @Router       /categories [get]
func GetCategoriesHandler(c *gin.Context) {
	ctgs, err := GetCategoriesDto()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request categories"})
		return
	}
	c.JSON(http.StatusOK, ctgs)
}

func PostCategoryHandler(c *gin.Context) {
	var newCategory CategoryEntity
	if err := c.ShouldBindJSON(&newCategory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Post category bad request deserializar"})
		return
	}

	// validate
	if newCategory.CatName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Category name must not be empty"})
		return
	}

	newCategory.CatName = strings.ToUpper(newCategory.CatName)

	// Register in the database
	ctg, err := CreateCategoryDto(newCategory)
	if err != nil {
		msg := fmt.Sprintf("The category could not be created in the database %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": msg})
		return
	}
	c.JSON(http.StatusOK,ctg)
}
