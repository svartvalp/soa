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
        "/product/index": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "responses": {
                    "200": {
                        "description": "OK"
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
                            "$ref": "#/definitions/models.Filter"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "integer"
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
                        "description": "Update products",
                        "name": "products",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.ProductInfo"
                            }
                        }
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
        "models.Filter": {
            "type": "object",
            "properties": {
                "brand": {
                    "type": "string"
                },
                "cat_id": {
                    "type": "integer"
                },
                "price_from": {
                    "type": "integer"
                },
                "price_to": {
                    "type": "integer"
                },
                "query": {
                    "type": "string"
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
        "models.ProductInfo": {
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
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:7001",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Search API",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
