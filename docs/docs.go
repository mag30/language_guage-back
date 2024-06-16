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
        "/user/:user-id/update": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Update User All Fields",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Update User All Fields",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "User data",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UpdateUserAllFieldRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/base.ResponseOK"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/base.ResponseFailure"
                        }
                    },
                    "500": {
                        "description": "Internal error (server fault)",
                        "schema": {
                            "$ref": "#/definitions/base.ResponseFailure"
                        }
                    }
                }
            }
        },
        "/user/authorizationFields/update": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Update User Authorization Fields",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Update User Authorization Fields",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "User data",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UpdateUserAuthorizationFieldsRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/base.ResponseOK"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/base.ResponseFailure"
                        }
                    },
                    "500": {
                        "description": "Internal error (server fault)",
                        "schema": {
                            "$ref": "#/definitions/base.ResponseFailure"
                        }
                    }
                }
            }
        },
        "/user/get": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get all users",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Get all users",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.GetUsersResponse"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/base.ResponseFailure"
                        }
                    },
                    "500": {
                        "description": "Internal error (server fault)",
                        "schema": {
                            "$ref": "#/definitions/base.ResponseFailure"
                        }
                    }
                }
            }
        },
        "/user/login": {
            "post": {
                "description": "User authorisation",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "User authorisation",
                "parameters": [
                    {
                        "description": "User request",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.LoginRequest"
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
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/base.ResponseFailure"
                        }
                    },
                    "500": {
                        "description": "Internal error (server fault)",
                        "schema": {
                            "$ref": "#/definitions/base.ResponseFailure"
                        }
                    }
                }
            }
        },
        "/user/logout": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Unauthorized users",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Unauthorized users",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "User request",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.RecreateJWTRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/base.ResponseOK"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/base.ResponseFailure"
                        }
                    },
                    "500": {
                        "description": "Internal error (server fault)",
                        "schema": {
                            "$ref": "#/definitions/base.ResponseFailure"
                        }
                    }
                }
            }
        },
        "/user/refresh": {
            "post": {
                "description": "Re-create refresh token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Re-create refresh token",
                "parameters": [
                    {
                        "description": "User request",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.RecreateJWTRequest"
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
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/base.ResponseFailure"
                        }
                    },
                    "500": {
                        "description": "Internal error (server fault)",
                        "schema": {
                            "$ref": "#/definitions/base.ResponseFailure"
                        }
                    }
                }
            }
        },
        "/user/register": {
            "post": {
                "description": "User registration",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "User registration",
                "parameters": [
                    {
                        "description": "User request",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.RegisterRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/base.ResponseOKWithID"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/base.ResponseFailure"
                        }
                    },
                    "500": {
                        "description": "Internal error (server fault)",
                        "schema": {
                            "$ref": "#/definitions/base.ResponseFailure"
                        }
                    }
                }
            }
        },
        "/user/retrieve": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Retrieve data of an authorised user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Retrieve data of an authorised user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.GetUserResponse"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/base.ResponseFailure"
                        }
                    },
                    "500": {
                        "description": "Internal error (server fault)",
                        "schema": {
                            "$ref": "#/definitions/base.ResponseFailure"
                        }
                    }
                }
            }
        },
        "/usersByIdList": {
            "post": {
                "description": "Retrieve user information by id list",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Retrieve user information by id list",
                "parameters": [
                    {
                        "description": "User data",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UsersByIdListRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.GetUsersResponse"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/base.ResponseFailure"
                        }
                    },
                    "500": {
                        "description": "Internal error (server fault)",
                        "schema": {
                            "$ref": "#/definitions/base.ResponseFailure"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "base.Blame": {
            "type": "string",
            "enum": [
                "User",
                "Postgres",
                "Server"
            ],
            "x-enum-varnames": [
                "BlameUser",
                "BlamePostgres",
                "BlameServer"
            ]
        },
        "base.ResponseFailure": {
            "type": "object",
            "properties": {
                "blame": {
                    "allOf": [
                        {
                            "$ref": "#/definitions/base.Blame"
                        }
                    ],
                    "example": "Guilty System"
                },
                "message": {
                    "type": "string",
                    "example": "error occurred"
                },
                "status": {
                    "type": "string",
                    "example": "Error"
                },
                "trackingID": {
                    "type": "string",
                    "example": "12345678-1234-1234-1234-000000000000"
                }
            }
        },
        "base.ResponseOK": {
            "type": "object",
            "properties": {
                "status": {
                    "type": "string",
                    "example": "OK"
                },
                "trackingID": {
                    "type": "string",
                    "example": "12345678-1234-1234-1234-000000000000"
                }
            }
        },
        "base.ResponseOKWithID": {
            "type": "object",
            "properties": {
                "ID": {
                    "type": "string",
                    "example": "12345678-1234-1234-1234-000000000000"
                },
                "status": {
                    "type": "string",
                    "example": "OK"
                },
                "trackingID": {
                    "type": "string",
                    "example": "12345678-1234-1234-1234-000000000000"
                }
            }
        },
        "enum.Level": {
            "type": "string",
            "enum": [
                "None",
                "Beginner",
                "Elementary",
                "Intermediate"
            ],
            "x-enum-varnames": [
                "None",
                "Beginner",
                "Elementary",
                "Intermediate"
            ]
        },
        "enum.Role": {
            "type": "string",
            "enum": [
                "Admin",
                "Guest"
            ],
            "x-enum-varnames": [
                "Admin",
                "Guest"
            ]
        },
        "model.GetUserResponse": {
            "type": "object",
            "properties": {
                "status": {
                    "type": "string",
                    "example": "OK"
                },
                "trackingID": {
                    "type": "string",
                    "example": "12345678-1234-1234-1234-000000000000"
                },
                "user": {
                    "$ref": "#/definitions/model.UserObject"
                }
            }
        },
        "model.GetUsersResponse": {
            "type": "object",
            "properties": {
                "status": {
                    "type": "string",
                    "example": "OK"
                },
                "trackingID": {
                    "type": "string",
                    "example": "12345678-1234-1234-1234-000000000000"
                },
                "users": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.UserObject"
                    }
                }
            }
        },
        "model.LoginRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "model.LoginResponse": {
            "type": "object",
            "properties": {
                "refreshToken": {
                    "type": "string"
                },
                "status": {
                    "type": "string",
                    "example": "OK"
                },
                "token": {
                    "type": "string"
                },
                "trackingID": {
                    "type": "string",
                    "example": "12345678-1234-1234-1234-000000000000"
                }
            }
        },
        "model.RecreateJWTRequest": {
            "type": "object",
            "properties": {
                "refreshToken": {
                    "type": "string"
                }
            }
        },
        "model.RegisterRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "model.UpdateUserAllFieldRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "fullName": {
                    "type": "string"
                },
                "level": {
                    "$ref": "#/definitions/enum.Level"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "model.UpdateUserAuthorizationFieldsRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "new_password": {
                    "type": "string"
                },
                "old_password": {
                    "type": "string"
                }
            }
        },
        "model.UserObject": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "level": {
                    "$ref": "#/definitions/enum.Level"
                },
                "name": {
                    "type": "string"
                },
                "role": {
                    "$ref": "#/definitions/enum.Role"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "model.UsersByIdListRequest": {
            "type": "object",
            "properties": {
                "ids": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
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
