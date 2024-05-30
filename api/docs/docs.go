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
                "description": "Checks the user data and returns a jwt token on correct Login",
                "tags": [
                    "Auth"
                ],
                "summary": "Logs in the user",
                "parameters": [
                    {
                        "description": "Login details",
                        "name": "details",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.PostUser"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/services.LoginResult"
                        }
                    }
                }
            }
        },
        "/auth/register": {
            "post": {
                "description": "Checks the user data and adds it to the repo",
                "tags": [
                    "Auth"
                ],
                "summary": "Registers the user",
                "parameters": [
                    {
                        "description": "Register details",
                        "name": "details",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.PostUser"
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
        "/task": {
            "get": {
                "description": "Fetches All of the user's tasks",
                "tags": [
                    "Tasks"
                ],
                "summary": "Fetch All tasks",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authenticator",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dto.GetTask"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Creates a new user task",
                "tags": [
                    "Tasks"
                ],
                "summary": "Creates a task",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authenticator",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "new task data",
                        "name": "task",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreateTask"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/dto.GetTask"
                        }
                    }
                }
            }
        },
        "/task/{taskId}": {
            "get": {
                "description": "Finds a task by it's task id",
                "tags": [
                    "Tasks"
                ],
                "summary": "Find a task by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authenticator",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Task ID",
                        "name": "taskId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.GetTask"
                        }
                    }
                }
            },
            "delete": {
                "description": "Deletes the task",
                "tags": [
                    "Tasks"
                ],
                "summary": "Delete task",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authenticator",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Task ID",
                        "name": "taskId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            },
            "patch": {
                "description": "Toggles the task's complete status",
                "tags": [
                    "Tasks"
                ],
                "summary": "Toggle complete status",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authenticator",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Task ID",
                        "name": "taskId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.GetTask"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.CreateTask": {
            "type": "object",
            "required": [
                "title"
            ],
            "properties": {
                "text": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "dto.GetTask": {
            "type": "object",
            "properties": {
                "completed": {
                    "type": "boolean"
                },
                "id": {
                    "type": "string"
                },
                "text": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "dto.PostUser": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string",
                    "minLength": 8
                },
                "username": {
                    "type": "string",
                    "minLength": 4
                }
            }
        },
        "services.LoginResult": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:9090",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "TODOapp api",
	Description:      "A siple TODO task service",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
