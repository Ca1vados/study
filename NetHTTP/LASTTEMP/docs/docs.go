// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/convert": {
            "post": {
                "description": "возвращает ответ на запрос по вычислению криптовалют",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "convert"
                ],
                "summary": "Get response",
                "parameters": [
                    {
                        "description": "форма для ковертации",
                        "name": "convertrequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.ConvertRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entity.ConvertResponse"
                            }
                        }
                    }
                }
            }
        },
        "/get": {
            "get": {
                "description": "получает курсы криптовалют и сохраняет их в бд (комментарий для пользователя)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "cryptorates"
                ],
                "summary": "Get rates",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entity.ApiBinanceResponse"
                            }
                        }
                    }
                }
            }
        },
        "/history": {
            "get": {
                "description": "возвращает историю хранимых курсов валют",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "cryptorates"
                ],
                "summary": "Get history",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entity.ApiBinanceResponse"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "entity.ApiBinanceResponse": {
            "type": "object",
            "properties": {
                "price": {
                    "type": "string"
                },
                "symbol": {
                    "type": "string"
                }
            }
        },
        "entity.ConvertRequest": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "number"
                },
                "from": {
                    "type": "string"
                },
                "to": {
                    "type": "string"
                }
            }
        },
        "entity.ConvertResponse": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "number"
                },
                "from": {
                    "type": "string"
                },
                "result": {
                    "type": "number"
                },
                "to": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Simple cryptoconverter",
	Description:      "Service for convertion crypto........",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}