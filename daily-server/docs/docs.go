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
        "/api/daily": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
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
                        "description": "Bad Request {\"message': \"Invalid JSON data\"}",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "401": {
                        "description": "Unauthorized {\"message': \"Unauthorized\"}",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error {\"message': \"Couldn't fetch the image\"}",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "502": {
                        "description": "Bad Gateway {\"message': \"Couldn't fetch the image / DB error\"}",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/api/daily/explore": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "returns a list of shared dailies",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Daily"
                ],
                "summary": "returns a list of shared dailies",
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
                        "description": "Bad Gateway {\"message': \"Failed to fetch Dailies\"}",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "502": {
                        "description": "Bad Gateway {\"message': \"No user\"}",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/api/daily/explorevs": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "returns 5 shared dailies",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Daily"
                ],
                "summary": "returns 5 shared dailies",
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
                        "description": "Internal Server Error {\"message': \"Failed to fetch Dailies\"}",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "502": {
                        "description": "Bad Gateway {\"message': \"No user\"}",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/api/daily/fav/{id}": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "fav a daily",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Daily"
                ],
                "summary": "update daily \u0026 user to apply fav feature",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Daily ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success {\"message\": \"Favourite Success\"}",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "400": {
                        "description": "Bad Request {\"message\": \"Invalid JSON data\"}",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "401": {
                        "description": "Bad Gateway {\"message\": \"Unauthorized\"}",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error {\"message\": \"Database error\"}",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/api/daily/image": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "edit a daily's image",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Daily"
                ],
                "summary": "update daily image",
                "parameters": [
                    {
                        "description": "EditDailyImageDTO",
                        "name": "daily",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.EditDailyImageDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success {\"message\": \"Image Edited Successfully\"}",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "400": {
                        "description": "Bad Request {\"message': \"Invalid JSON data\"}",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "500": {
                        "description": "Bad Gateway {\"message\": \"Database Error\"}",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/api/daily/list": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
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
                "parameters": [
                    {
                        "type": "integer",
                        "description": "limit by q",
                        "name": "limit",
                        "in": "query"
                    }
                ],
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
                        "description": "Bad Gateway {\"message': \"Couldn't fetch daily list\"}",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "502": {
                        "description": "Bad Gateway {\"message': \"No user\"}",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/api/daily/report": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "report a daily",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Daily"
                ],
                "summary": "update daily to apply report feature",
                "parameters": [
                    {
                        "description": "ReportedDaily",
                        "name": "daily",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.ReportedDaily"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success {\"message\": \"Created Successfully\"}",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "400": {
                        "description": "Bad Request {\"message': \"Invalid JSON data\"}",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "502": {
                        "description": "Bad Gateway {\"message\": \"Failed to update daily\"}",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/api/daily/statistics": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "provides statistical data about a user's activity including likes, views, number of dailies written, current mood, streak, and a predefined topic",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Statistics"
                ],
                "summary": "Get user statistics",
                "responses": {
                    "200": {
                        "description": "An object of statistics including likes, views, dailies written, mood, streak, and topic",
                        "schema": {
                            "$ref": "#/definitions/model.StatisticsDTO"
                        }
                    },
                    "400": {
                        "description": "bad request - error message",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "unauthorized - error message",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/daily/view": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "view a daily",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Daily"
                ],
                "summary": "update daily \u0026 user to apply view feature",
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
                        "description": "Success {\"message\": \"Viewed Successfully\"}",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "400": {
                        "description": "Bad Request {\"message\": \"Invalid JSON data\"}",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "401": {
                        "description": "Bad Gateway {\"message\": \"Wrong user id\"}",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "500": {
                        "description": "Bad Gateway {\"message\": \"Database error\"}",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/api/daily/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "returns a specific daily via daily.ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Daily"
                ],
                "summary": "returns a daily",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Daily ID",
                        "name": "id",
                        "in": "path",
                        "required": true
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
                    "500": {
                        "description": "Internal Server Error {\"message': \"mongo: no documents in result\"}",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "report a daily",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Daily"
                ],
                "summary": "delete the given daily",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Daily ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success {\"message\": \"Deleted Successfully\"}",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "400": {
                        "description": "Unauthorized {\"message': \"Unauthorized\"}",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "502": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/api/login": {
            "post": {
                "description": "Authenticate a user using the provided email and password, and return a token on successful authentication.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "User login",
                "parameters": [
                    {
                        "description": "User login details",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UserLoginDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Token",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request {\"message': \"Invalid JSON data\"}",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/api/register": {
            "post": {
                "description": "Create a new user with the given email and password, if they don't exist already.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Register a new user",
                "parameters": [
                    {
                        "description": "User Registration",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UserRegisterDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    },
                    "400": {
                        "description": "Bad Request {\"message': \"Invalid JSON data\"}",
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
                "prompt": {
                    "type": "string"
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
                "embedding": {
                    "type": "array",
                    "items": {
                        "type": "number"
                    }
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
                "topics": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
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
        "model.EditDailyImageDTO": {
            "type": "object",
            "required": [
                "id"
            ],
            "properties": {
                "id": {
                    "type": "string"
                },
                "image": {
                    "type": "string"
                }
            }
        },
        "model.Emotion": {
            "type": "object",
            "properties": {
                "anger": {
                    "type": "number"
                },
                "fear": {
                    "type": "number"
                },
                "joy": {
                    "type": "number"
                },
                "love": {
                    "type": "number"
                },
                "sadness": {
                    "type": "number"
                },
                "surprise": {
                    "type": "number"
                }
            }
        },
        "model.ReportedDaily": {
            "type": "object",
            "required": [
                "dailyId",
                "title"
            ],
            "properties": {
                "content": {
                    "type": "string"
                },
                "dailyId": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "reportedAt": {
                    "type": "integer"
                },
                "reports": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "model.StatisticsDTO": {
            "type": "object",
            "properties": {
                "dailiesWritten": {
                    "description": "Number of dailies written",
                    "type": "integer"
                },
                "date": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "likes": {
                    "description": "Number of likes received",
                    "type": "integer"
                },
                "mood": {
                    "description": "Current mood based on user's entries",
                    "type": "string"
                },
                "streak": {
                    "description": "Current streak of daily entries",
                    "type": "integer"
                },
                "topics": {
                    "description": "Currently focused topic",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "views": {
                    "description": "Number of views",
                    "type": "integer"
                }
            }
        },
        "model.User": {
            "type": "object",
            "required": [
                "email",
                "isVerified",
                "password"
            ],
            "properties": {
                "createdAt": {
                    "type": "integer"
                },
                "email": {
                    "type": "string"
                },
                "favouriteDailies": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "id": {
                    "type": "string"
                },
                "isVerified": {
                    "type": "boolean"
                },
                "password": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                },
                "viewedDailies": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "model.UserLoginDTO": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "model.UserRegisterDTO": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
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
