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
        "/auth/login": {
            "post": {
                "description": "Authenticates a user with email/username and password, and returns a JWT token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "User login",
                "parameters": [
                    {
                        "description": "User login data",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schemas.LoginSchema"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.SuccessBody"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responses.FailureBody"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/responses.FailureBody"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/responses.FailureBody"
                        }
                    }
                }
            }
        },
        "/auth/register": {
            "post": {
                "description": "Creates a new user with the provided details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Register a new customer",
                "parameters": [
                    {
                        "description": "User registration data",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schemas.RegisterSchema"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/conversation": {
            "post": {
                "security": [
                    {
                        "AuthorizationToken": []
                    }
                ],
                "description": "Creates a new direct or group chat",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Conversation"
                ],
                "summary": "Start Conversation",
                "parameters": [
                    {
                        "description": "Conversation Data",
                        "name": "conversation",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schemas.StartConversationSchema"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/responses.SuccessBody"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responses.FailureBody"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/responses.FailureBody"
                        }
                    }
                }
            }
        },
        "/conversation/{id}": {
            "get": {
                "security": [
                    {
                        "AuthorizationToken": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Conversation"
                ],
                "summary": "Get Conversation Details",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Conversation ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Conversation"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responses.FailureBody"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/responses.FailureBody"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/responses.FailureBody"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "AuthorizationToken": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Conversation"
                ],
                "summary": "Delete Conversation",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Conversation ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.SuccessBody"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responses.FailureBody"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/responses.FailureBody"
                        }
                    }
                }
            }
        },
        "/conversation/{id}/member/{user_id}": {
            "post": {
                "security": [
                    {
                        "AuthorizationToken": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Conversation"
                ],
                "summary": "Add Member",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Conversation ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.SuccessBody"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responses.FailureBody"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/responses.FailureBody"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "AuthorizationToken": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Conversation"
                ],
                "summary": "Remove Member",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Conversation ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.SuccessBody"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responses.FailureBody"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/responses.FailureBody"
                        }
                    }
                }
            }
        },
        "/conversations/member/{user_id}": {
            "get": {
                "security": [
                    {
                        "AuthorizationToken": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Conversation"
                ],
                "summary": "Get User Conversations",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Page number (default: 1)",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Items per page (default: 10)",
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
                                "$ref": "#/definitions/models.ConversationWithMembers"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responses.FailureBody"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/responses.FailureBody"
                        }
                    }
                }
            }
        },
        "/user/{id}": {
            "get": {
                "security": [
                    {
                        "AuthorizationToken": []
                    }
                ],
                "description": "Fetches user details by user ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Get User Details",
                "parameters": [
                    {
                        "type": "string",
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
                            "$ref": "#/definitions/responses.SuccessBody"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responses.FailureBody"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/responses.FailureBody"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/responses.FailureBody"
                        }
                    }
                }
            },
            "patch": {
                "security": [
                    {
                        "AuthorizationToken": []
                    }
                ],
                "description": "Updates user details by user ID. Only the provided fields will be updated.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Update User Details",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "User update data",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schemas.UpdateUserSchema"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User details updated successfully",
                        "schema": {
                            "$ref": "#/definitions/responses.SuccessBody"
                        }
                    },
                    "400": {
                        "description": "Invalid request data",
                        "schema": {
                            "$ref": "#/definitions/responses.FailureBody"
                        }
                    },
                    "404": {
                        "description": "User not found",
                        "schema": {
                            "$ref": "#/definitions/responses.FailureBody"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/responses.FailureBody"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "enums.UserStatus": {
            "type": "string",
            "enum": [
                "active",
                "inactive",
                "banned"
            ],
            "x-enum-varnames": [
                "UserActive",
                "UserInactive",
                "UserBanned"
            ]
        },
        "models.Conversation": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "groupPhotoPath": {
                    "type": "string"
                },
                "groupPhotoUrl": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "isGroup": {
                    "type": "boolean"
                },
                "name": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "models.ConversationWithMembers": {
            "type": "object",
            "properties": {
                "conversation": {
                    "$ref": "#/definitions/models.Conversation"
                },
                "members": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.User"
                    }
                }
            }
        },
        "models.User": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "dateOfBirth": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "firstName": {
                    "type": "string"
                },
                "gender": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "isOwner": {
                    "type": "boolean"
                },
                "isStaff": {
                    "type": "boolean"
                },
                "lastName": {
                    "type": "string"
                },
                "lastSeen": {
                    "type": "string"
                },
                "mobileNumber": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "profilePhotoPath": {
                    "type": "string"
                },
                "profilePhotoUrl": {
                    "type": "string"
                },
                "status": {
                    "$ref": "#/definitions/enums.UserStatus"
                },
                "updatedAt": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "responses.FailureBody": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "responses.SuccessBody": {
            "type": "object",
            "properties": {
                "data": {},
                "success": {
                    "type": "boolean"
                }
            }
        },
        "schemas.LoginSchema": {
            "type": "object",
            "required": [
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string",
                    "maxLength": 32,
                    "minLength": 8
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "schemas.RegisterSchema": {
            "type": "object",
            "required": [
                "date_of_birth",
                "email",
                "first_name",
                "gender",
                "last_name",
                "mobile_number",
                "password",
                "username"
            ],
            "properties": {
                "date_of_birth": {
                    "description": "Keep as string",
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 2
                },
                "gender": {
                    "type": "string",
                    "enum": [
                        "male",
                        "female",
                        "other"
                    ]
                },
                "last_name": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 2
                },
                "mobile_number": {
                    "type": "string"
                },
                "password": {
                    "type": "string",
                    "maxLength": 32,
                    "minLength": 8
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "schemas.StartConversationSchema": {
            "type": "object",
            "required": [
                "members"
            ],
            "properties": {
                "is_group": {
                    "type": "boolean"
                },
                "members": {
                    "type": "array",
                    "maxItems": 1024,
                    "minItems": 1,
                    "items": {
                        "type": "string"
                    }
                },
                "name": {
                    "type": "string",
                    "maxLength": 100,
                    "minLength": 1
                }
            }
        },
        "schemas.UpdateUserSchema": {
            "type": "object",
            "properties": {
                "date_of_birth": {
                    "description": "Keep as string",
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 2
                },
                "gender": {
                    "type": "string",
                    "enum": [
                        "male",
                        "female",
                        "other"
                    ]
                },
                "last_name": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 2
                },
                "mobile_number": {
                    "type": "string"
                },
                "password": {
                    "type": "string",
                    "maxLength": 32,
                    "minLength": 8
                },
                "username": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "AuthorizationToken": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
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
