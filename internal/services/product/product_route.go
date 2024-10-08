package product

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// GetProducts godoc
// @Summary      Get all products
// @Description  Return the list of products
// @Tags         products
// @Produce      json
// @Success      200 {array} Product
// @Router       /products [get]
func GetProductsHandler(c *gin.Context) {
	pro, err := GetProductsWithCategoryDto()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Get Products bad Request"})
		return
	}
	c.JSON(http.StatusOK, pro)
}


// PostProduct godoc
// @Summary      Create product
// @Description  Create a new product
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        product body Product true "Product to be created"
// @Success      200 {object} Product
// @Router       /products [post]
func PostProductHandler(c *gin.Context) {
	var newProduct ProductEntity

	if err := c.ShouldBindJSON(&newProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Post products bad request deserializar"})
		return
	}

	//validate
	validateNewProduct := validator.New()
	err := validateNewProduct.Struct(newProduct)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		msg := fmt.Sprintf("messsage post product: %v", errors)
		c.JSON(http.StatusBadRequest, gin.H{"message": msg})
		return
	}

	//add register to databases
	p, err := CreateProduct(newProduct)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Post product bad request testt"})
		return
	}
	c.JSON(http.StatusOK, p)
}


func ValidateProduct(p ProductEntity) (isValid bool, err error) {
	if p.ProductName == "" || p.ProductDescription == ""{
		return false, fmt.Errorf("product name or product description connot be empty")
	}

	if  p.ProductPrice <= 0 || p.ProductStk <= 0 {
		return false, fmt.Errorf("product price or product stock must be assigned")
	}
	if p.ProductDiscount < 0 {
		return false, fmt.Errorf("invalid product discount")
	}
	return true, nil
}



// PutProduct godoc
// @Summary      Update product
// @Description  Update a product by its ID
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Product ID"
// @Param        product body Product true "Product to be updated"
// @Success      200 {object} Product
// @Router       /products/{id} [put]
func PutProductHandler(c *gin.Context) {
	id := c.Param("id")
	productId, err := strconv.Atoi(id)
	if err != nil { 
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
		return
	}
	var product ProductEntity
	err = c.ShouldBindJSON(&product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request put product deserializer"})
		return
	}

	validateProduct := validator.New()
	err = validateProduct.Struct(product)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		msg := fmt.Sprintf("messsage post user: %v", errors)
		c.JSON(http.StatusBadRequest, gin.H{"message": msg})
		return
	}

	err = UpdateProduct(int64(productId), product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request update product"})
		return
	}
	//tiene que devolver el  producto?
	c.JSON(http.StatusNoContent, nil)
}


// DeleteProduct godoc
// @Summary      Delete product
// @Description  Delete a product by its ID
// @Tags         products
// @Param        id   path      int  true  "Product ID"
// @Success      204  "Product successfully deleted"
// @Router       /products/{id} [delete]
func DeleteProductHandler(c *gin.Context) {
	id := c.Param("id")
	idProduct, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
		return
	}
	if err = DeleteProduct(int64(idProduct)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request delete product"})
		return
	}
	c.JSON(http.StatusOK, nil)
}
