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
        "/receipts/process": {
            "post": {
                "description": "Create a new receipt by providing the retailer, purchase date, time, items, and total amount",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "receipts"
                ],
                "summary": "Create a new receipt",
                "parameters": [
                    {
                        "description": "Receipt input data",
                        "name": "receiptInput",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReceiptInput"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Receipt created successfully",
                        "schema": {
                            "$ref": "#/definitions/models.Receipt"
                        }
                    },
                    "400": {
                        "description": "Bad request, invalid input",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/receipts/{id}/points": {
            "get": {
                "description": "Fetch the receipt by its ID and calculate the points based on predefined rules",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "receipts"
                ],
                "summary": "Get points11 for a receipt by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Receipt ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Points calculated successfully",
                        "schema": {
                            "$ref": "#/definitions/models.PointsResponse"
                        }
                    },
                    "400": {
                        "description": "Bad request, receipt ID is required or receipt not found",
                        "schema": {
                            "$ref": "#/definitions/models.PointsErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.ErrorResponse": {
            "type": "object",
            "properties": {
                "data": {},
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "models.Item": {
            "type": "object",
            "properties": {
                "price": {
                    "type": "string"
                },
                "shortDescription": {
                    "type": "string"
                }
            }
        },
        "models.PointsErrorResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "models.PointsResponse": {
            "type": "object",
            "properties": {
                "points": {
                    "type": "integer"
                }
            }
        },
        "models.Receipt": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Item"
                    }
                },
                "purchaseDate": {
                    "type": "string"
                },
                "purchaseTime": {
                    "type": "string"
                },
                "retailer": {
                    "type": "string"
                },
                "total": {
                    "type": "string"
                }
            }
        },
        "models.ReceiptInput": {
            "type": "object",
            "properties": {
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Item"
                    }
                },
                "purchaseDate": {
                    "type": "string"
                },
                "purchaseTime": {
                    "type": "string"
                },
                "retailer": {
                    "type": "string"
                },
                "total": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
