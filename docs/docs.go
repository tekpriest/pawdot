// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "consumes": [
        "application/json"
    ],
    "produces": [
        "application/json"
    ],
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {
            "name": "MIT"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/auth/login": {
            "post": {
                "tags": [
                    "Auth"
                ],
                "operationId": "login",
                "parameters": [
                    {
                        "description": "Body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/Login"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/LoginSuccessfulResponse"
                        }
                    }
                }
            }
        },
        "/api/auth/register": {
            "post": {
                "tags": [
                    "Auth"
                ],
                "operationId": "register",
                "parameters": [
                    {
                        "description": "Body",
                        "name": "Register",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/Register"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/RegistrationSuccessfulResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "Login": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "LoginSuccessfulResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/auth.UserAuthData"
                },
                "mesage": {
                    "type": "string",
                    "example": "account created successfully"
                },
                "success": {
                    "type": "boolean",
                    "example": true
                }
            }
        },
        "Register": {
            "type": "object",
            "required": [
                "accountType",
                "email",
                "password",
                "username"
            ],
            "properties": {
                "accountType": {
                    "$ref": "#/definitions/models.AccountType"
                },
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "RegistrationSuccessfulResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/auth.UserAuthData"
                },
                "mesage": {
                    "type": "string",
                    "example": "account created successfully"
                },
                "success": {
                    "type": "boolean",
                    "example": true
                }
            }
        },
        "auth.UserAuthData": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/models.User"
                }
            }
        },
        "models.AccountType": {
            "type": "string",
            "enum": [
                "BUYER",
                "SELLER"
            ],
            "x-enum-varnames": [
                "BUYER",
                "SELLER"
            ]
        },
        "models.Bid": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "number"
                },
                "createdAt": {
                    "type": "string",
                    "readOnly": true,
                    "example": "2022-08-21 21:08"
                },
                "id": {
                    "type": "string",
                    "readOnly": true
                },
                "saleId": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                },
                "userId": {
                    "type": "string"
                }
            }
        },
        "models.Sale": {
            "type": "object",
            "properties": {
                "bidCount": {
                    "type": "integer"
                },
                "bids": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Bid"
                    }
                },
                "breed": {
                    "type": "string"
                },
                "category": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string",
                    "readOnly": true,
                    "example": "2022-08-21 21:08"
                },
                "description": {
                    "type": "string"
                },
                "expiresBy": {
                    "type": "string"
                },
                "id": {
                    "type": "string",
                    "readOnly": true
                },
                "priority": {
                    "$ref": "#/definitions/models.SalePriority"
                },
                "sold": {
                    "type": "boolean"
                },
                "soldTo": {
                    "type": "string"
                },
                "startingBig": {
                    "type": "number"
                },
                "status": {
                    "$ref": "#/definitions/models.SaleStatus"
                },
                "title": {
                    "type": "string"
                },
                "traderId": {
                    "type": "string"
                },
                "type": {
                    "$ref": "#/definitions/models.SaleType"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "models.SalePriority": {
            "type": "integer",
            "enum": [
                0,
                1,
                2
            ],
            "x-enum-varnames": [
                "LOW",
                "NORMAL",
                "HIGH"
            ]
        },
        "models.SaleStatus": {
            "type": "string",
            "enum": [
                "PENDING",
                "PUBLISHED",
                "CANCELLED",
                "CLOSED"
            ],
            "x-enum-varnames": [
                "PENDING",
                "PUBLISHED",
                "CANCELLED",
                "CLOSED"
            ]
        },
        "models.SaleType": {
            "type": "string",
            "enum": [
                "NEW",
                "PROMOTED",
                "REPUBLISHED"
            ],
            "x-enum-varnames": [
                "NEW",
                "PROMOTED",
                "REPUBLISHED"
            ]
        },
        "models.User": {
            "type": "object",
            "properties": {
                "accountType": {
                    "$ref": "#/definitions/models.AccountType"
                },
                "bids": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Bid"
                    }
                },
                "createdAt": {
                    "type": "string",
                    "readOnly": true,
                    "example": "2022-08-21 21:08"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string",
                    "readOnly": true
                },
                "password": {
                    "type": "string"
                },
                "profileImg": {
                    "type": "string"
                },
                "sales": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Sale"
                    }
                },
                "updatedAt": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                },
                "wallet": {
                    "$ref": "#/definitions/models.Wallet"
                }
            }
        },
        "models.Wallet": {
            "type": "object",
            "properties": {
                "balance": {
                    "type": "number"
                },
                "createdAt": {
                    "type": "string",
                    "readOnly": true,
                    "example": "2022-08-21 21:08"
                },
                "id": {
                    "type": "string",
                    "readOnly": true
                },
                "updatedAt": {
                    "type": "string"
                },
                "userID": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:3000",
	BasePath:         "/api/",
	Schemes:          []string{},
	Title:            "Pawdot API Service",
	Description:      "Pawdow API Service",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
