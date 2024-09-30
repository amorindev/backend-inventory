// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/categories": {
            "get": {
                "description": "Return the list of categories",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "categories"
                ],
                "summary": "Get all categories",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/categories.Category"
                            }
                        }
                    }
                }
            }
        },
        "/kardex": {
            "get": {
                "description": "Return the list of kardex entries",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "kardex"
                ],
                "summary": "Get all kardex entries",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/kardex.Kardex"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new kardex",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "kardex"
                ],
                "summary": "Create kardex",
                "parameters": [
                    {
                        "description": "Kardex to be created",
                        "name": "kardex",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/kardex.Kardex"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/kardex.Kardex"
                        }
                    }
                }
            }
        },
        "/products": {
            "get": {
                "description": "Return the list of products",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "products"
                ],
                "summary": "Get all products",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/product.Product"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new product",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "products"
                ],
                "summary": "Create product",
                "parameters": [
                    {
                        "description": "Product to be created",
                        "name": "product",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/product.Product"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/product.Product"
                        }
                    }
                }
            }
        },
        "/products/{id}": {
            "put": {
                "description": "Update a product by its ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "products"
                ],
                "summary": "Update product",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Product ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Product to be updated",
                        "name": "product",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/product.Product"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/product.Product"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a product by its ID",
                "tags": [
                    "products"
                ],
                "summary": "Delete product",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Product ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "Product successfully deleted"
                    }
                }
            }
        }
    },
    "definitions": {
        "categories.Category": {
            "type": "object",
            "required": [
                "cat_name"
            ],
            "properties": {
                "cat_id": {
                    "type": "integer"
                },
                "cat_name": {
                    "type": "string"
                }
            }
        },
        "kardex.Kardex": {
            "type": "object",
            "required": [
                "kardex_description",
                "kardex_products",
                "kardex_type"
            ],
            "properties": {
                "kardex_created_at": {
                    "type": "string"
                },
                "kardex_description": {
                    "type": "string"
                },
                "kardex_id": {
                    "type": "integer"
                },
                "kardex_products": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/kardex.KardexProduct"
                    }
                },
                "kardex_type": {
                    "type": "string"
                }
            }
        },
        "kardex.KardexProduct": {
            "type": "object",
            "properties": {
                "pro_kar_amount": {
                    "description": "Cantidad movida en la",
                    "type": "integer",
                    "minimum": 0
                },
                "prod_id": {
                    "description": "ID del producto",
                    "type": "integer"
                },
                "prod_name": {
                    "type": "string"
                }
            }
        },
        "product.Product": {
            "type": "object",
            "required": [
                "cat_id",
                "prod_desc",
                "prod_name",
                "prod_price",
                "prod_stk"
            ],
            "properties": {
                "cat_id": {
                    "type": "integer"
                },
                "prod_desc": {
                    "type": "string"
                },
                "prod_discount": {
                    "description": "campo omitido 0",
                    "type": "integer",
                    "minimum": 0
                },
                "prod_id": {
                    "type": "integer"
                },
                "prod_name": {
                    "type": "string"
                },
                "prod_price": {
                    "type": "number",
                    "minimum": 0
                },
                "prod_stk": {
                    "type": "integer",
                    "minimum": 0
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:7000",
	BasePath:         "/v1",
	Schemes:          []string{},
	Title:            "Inventory api",
	Description:      "Api for an app of inventory",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
