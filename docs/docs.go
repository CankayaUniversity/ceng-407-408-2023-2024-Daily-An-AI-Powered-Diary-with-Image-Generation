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
        "/api/CreateDaily": {
            "post": {
                "description": "creates a new daily resource",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Daily"
                ],
                "summary": "returns the created daily",
                "parameters": [
                    {
                        "description": "CreateDailyDTO",
                        "name": "daily",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CreateDailyDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Daily"
                        }
                    },
                    "400": {
                        "description": "Bad Request {\"message\": \"Invalid JSON data\"}",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "502": {
                        "description": "Bad Gateway {\"message': \"Couldn't fetch the image\"}",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/api/GetDailies": {
            "get": {
                "description": "returns a list of dailies",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Daily"
                ],
                "summary": "returns a list of dailies",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Daily"
                            }
                        }
                    },
                    "500": {
                        "description": "Bad Gateway {\"message': \"Couldn't fetch the image\"}",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/api/getDaily": {
            "get": {
                "description": "return a specific daily via daily.ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Daily"
                ],
                "summary": "return a daily",
                "parameters": [
                    {
                        "description": "DailyRequestDTO",
                        "name": "daily",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.DailyRequestDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Daily"
                        }
                    },
                    "400": {
                        "description": "Bad Request {\"message\": \"Invalid JSON data\"}",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "502": {
                        "description": "Bad Gateway {\"message': \"mongo: no documents in result\"}",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.CreateDailyDTO": {
            "type": "object",
            "required": [
                "isShared",
                "text"
            ],
            "properties": {
                "image": {
                    "type": "string"
                },
                "isShared": {
                    "type": "boolean"
                },
                "text": {
                    "type": "string"
                }
            }
        },
        "model.Daily": {
            "type": "object",
            "required": [
                "text"
            ],
            "properties": {
                "author": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "integer"
                },
                "emotions": {
                    "$ref": "#/definitions/model.Emotion"
                },
                "favourites": {
                    "type": "integer"
                },
                "id": {
                    "type": "string"
                },
                "image": {
                    "type": "string"
                },
                "isShared": {
                    "type": "boolean"
                },
                "keywords": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "text": {
                    "type": "string"
                },
                "viewers": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "model.DailyRequestDTO": {
            "type": "object",
            "required": [
                "id"
            ],
            "properties": {
                "id": {
                    "type": "string"
                }
            }
        },
        "model.Emotion": {
            "type": "object",
            "properties": {
                "anger": {
                    "type": "integer"
                },
                "happiness": {
                    "type": "integer"
                },
                "sadness": {
                    "type": "integer"
                },
                "shock": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:9090",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "Daily API",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
