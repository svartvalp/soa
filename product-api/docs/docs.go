// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
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
        "/category/create": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "description": "Create category",
                        "name": "products",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreateCategoryReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/category/list": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Category"
                            }
                        }
                    }
                }
            }
        },
        "/category/update": {
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "description": "Update category by id",
                        "name": "product",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Category"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/category/{id}/delete": {
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Delete category by id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/characteristic/create": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "description": "Create characteristic",
                        "name": "products",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreateCharacteristicReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/characteristic/list": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Characteristic"
                            }
                        }
                    }
                }
            }
        },
        "/characteristic/update": {
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "description": "Update characteristic by id",
                        "name": "product",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Characteristic"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/characteristic/{id}/delete": {
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Delete characteristic by id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/product/brand/list": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/product/create": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "description": "Create product list",
                        "name": "products",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreateProductReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/product/full-info": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "description": "List product filters",
                        "name": "filter",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ProductFilters"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.FullProductInfo"
                            }
                        }
                    }
                }
            }
        },
        "/product/list": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "description": "List product filters",
                        "name": "filter",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ProductFilters"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Product"
                            }
                        }
                    }
                }
            }
        },
        "/product/update": {
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "description": "Update product by id",
                        "name": "product",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Product"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/product/{id}/add": {
            "post": {
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "type": "file",
                        "description": "Product image",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Set product image by id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/product/{id}/delete": {
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Delete product by id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.CreateCategoryReq": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "level": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "parent_id": {
                    "type": "integer"
                }
            }
        },
        "dto.CreateCharacteristicReq": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "parent_id": {
                    "type": "string"
                }
            }
        },
        "dto.CreateProductReq": {
            "type": "object",
            "properties": {
                "brand": {
                    "type": "string"
                },
                "category_id": {
                    "type": "integer"
                },
                "description": {
                    "type": "string"
                },
                "image": {
                    "$ref": "#/definitions/dto.Image"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "integer"
                }
            }
        },
        "dto.Image": {
            "type": "object",
            "properties": {
                "body": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "models.Category": {
            "type": "object",
            "properties": {
                "ID": {
                    "type": "integer"
                },
                "description": {
                    "type": "string"
                },
                "level": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "parentID": {
                    "type": "integer"
                }
            }
        },
        "models.Characteristic": {
            "type": "object",
            "properties": {
                "ID": {
                    "type": "integer"
                },
                "chType": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "models.FullProductInfo": {
            "type": "object",
            "properties": {
                "brand": {
                    "type": "string"
                },
                "category_id": {
                    "type": "integer"
                },
                "categorys": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Category"
                    }
                },
                "characteristics": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.ProductCharacteristic"
                    }
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "image": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "integer"
                }
            }
        },
        "models.Product": {
            "type": "object",
            "properties": {
                "brand": {
                    "type": "string"
                },
                "category_id": {
                    "type": "integer"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "image": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "integer"
                }
            }
        },
        "models.ProductCharacteristic": {
            "type": "object",
            "properties": {
                "ID": {
                    "type": "integer"
                },
                "chType": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "productId": {
                    "type": "integer"
                }
            }
        },
        "models.ProductFilters": {
            "type": "object",
            "properties": {
                "category_ids": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "characteristic_ids": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "ids": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "names": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "product_characteristic_ids": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:7002",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Product API",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
