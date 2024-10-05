package product

import (
	"fmt"
	"testing"
)


func TestProduct(t *testing.T) {
	p := ProductEntity{
		ID: 6,
		ProductName: "Zapatilla addidas",
		ProductDescription: "Zapatilla adidas para hombre color azul",
		ProductPrice: 120,
		ProductStk: 50,
		ProductDiscount: 0,
		CatID: 2,
	}

	isValid, err := ValidateProduct(p)
	fmt.Println("Is valid: ", isValid)
	fmt.Println("Error: ", err)
}

