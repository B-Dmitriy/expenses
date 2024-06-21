// Package api Code generated by swaggo/swag. DO NOT EDIT
package api

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "API Support",
            "email": "example@mail.com"
        },
        "license": {
            "name": "MIT",
            "url": "https://opensource.org/license/mit"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/login": {
            "post": {
                "description": "Вход под учётной записью",
                "tags": [
                    "Authentication"
                ],
                "parameters": [
                    {
                        "description": "query params",
                        "name": "request",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/model.LoginCredentials"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.LoginResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/web.WebError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/web.WebError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/web.WebError"
                        }
                    }
                }
            }
        },
        "/api/v1/logout": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Выход из учётной записи",
                "tags": [
                    "Authentication"
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/web.WebError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/web.WebError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/web.WebError"
                        }
                    }
                }
            }
        },
        "/api/v1/refresh": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Перегенерировать токен пользователя",
                "tags": [
                    "Authentication"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/tokens.Tokens"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/web.WebError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/web.WebError"
                        }
                    }
                }
            }
        },
        "/api/v1/registration": {
            "post": {
                "description": "Регистрация пользователя",
                "tags": [
                    "Authentication"
                ],
                "parameters": [
                    {
                        "description": "query params",
                        "name": "request",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/model.RegistrationData"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/web.WebError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/web.WebError"
                        }
                    }
                }
            }
        },
        "/api/v1/users": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Получить список пользователей (только для админа)",
                "tags": [
                    "Users"
                ],
                "parameters": [
                    {
                        "maximum": 10,
                        "minimum": 1,
                        "type": "integer",
                        "default": 1,
                        "description": "positive int",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "maximum": 100,
                        "minimum": 1,
                        "type": "integer",
                        "default": 25,
                        "description": "positive int",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "maxLength": 256,
                        "type": "string",
                        "description": "any string",
                        "name": "search",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.UserInfo"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/web.WebError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/web.WebError"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Создать пользователя вручную. (Только админ)",
                "tags": [
                    "Users"
                ],
                "parameters": [
                    {
                        "description": "query params",
                        "name": "request",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/model.CreateUserBody"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/web.WebError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/web.WebError"
                        }
                    }
                }
            }
        },
        "/api/v1/users/{id}": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Получить информацию о пользователе (Пользователь - о семе, Админ о любом)",
                "tags": [
                    "Users"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.UserInfo"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/web.WebError"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/web.WebError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/web.WebError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/web.WebError"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Изменить информацию о пользователе (Пользователь - о семе, Админ о любом)",
                "tags": [
                    "Users"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "query params",
                        "name": "request",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/model.EditUserBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/web.WebError"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/web.WebError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/web.WebError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/web.WebError"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Удалить пользователя (Пользователь - только себя, Админ - любого)",
                "tags": [
                    "Users"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/web.WebError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/web.WebError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/web.WebError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.CreateUserBody": {
            "type": "object",
            "required": [
                "email",
                "login",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "login": {
                    "type": "string",
                    "maxLength": 255,
                    "minLength": 2
                },
                "password": {
                    "type": "string",
                    "maxLength": 20,
                    "minLength": 4
                }
            }
        },
        "model.EditUserBody": {
            "type": "object",
            "required": [
                "email",
                "login"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "login": {
                    "type": "string",
                    "maxLength": 255,
                    "minLength": 2
                }
            }
        },
        "model.LoginCredentials": {
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
                    "type": "string",
                    "maxLength": 20,
                    "minLength": 4
                }
            }
        },
        "model.LoginResponse": {
            "type": "object",
            "properties": {
                "accessToken": {
                    "type": "string"
                },
                "login": {
                    "type": "string"
                },
                "refreshToken": {
                    "type": "string"
                },
                "userID": {
                    "type": "integer"
                },
                "userRoleID": {
                    "type": "integer"
                }
            }
        },
        "model.RegistrationData": {
            "type": "object",
            "required": [
                "email",
                "login",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "login": {
                    "type": "string",
                    "maxLength": 255,
                    "minLength": 2
                },
                "password": {
                    "type": "string",
                    "maxLength": 20,
                    "minLength": 4
                }
            }
        },
        "model.UserInfo": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "emailConfirmed": {
                    "type": "boolean"
                },
                "id": {
                    "type": "integer"
                },
                "login": {
                    "type": "string"
                },
                "roleID": {
                    "type": "integer"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "tokens.Tokens": {
            "type": "object",
            "properties": {
                "acceptToken": {
                    "type": "string"
                },
                "refreshToken": {
                    "type": "string"
                }
            }
        },
        "web.WebError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "0.0.0.0:7070",
	BasePath:         "/",
	Schemes:          []string{"http"},
	Title:            "Swagger Expenses API",
	Description:      "This is server Expenses application.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
