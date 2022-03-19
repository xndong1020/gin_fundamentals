// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/albums": {
            "get": {
                "description": "Get Albums list with pagination",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Album"
                ],
                "operationId": "get-albums-list",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 0,
                        "description": "pagination current page",
                        "name": "page",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "default": 0,
                        "description": "pagination page_size",
                        "name": "page_size",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.AlbumResponse"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create new album",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Album"
                ],
                "operationId": "create-new-album",
                "parameters": [
                    {
                        "description": "album data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entities.Album"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entities.Album"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            }
        },
        "/albums/{id}": {
            "get": {
                "description": "Get Album By Id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Album"
                ],
                "operationId": "get-albums-by-id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "album Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.AlbumResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete Album By Id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Album"
                ],
                "operationId": "delete-albums-by-id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "album Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "entities.Album": {
            "type": "object",
            "required": [
                "artist",
                "id",
                "price",
                "title"
            ],
            "properties": {
                "artist": {
                    "type": "string"
                },
                "has_read": {
                    "type": "boolean"
                },
                "id": {
                    "type": "integer",
                    "minimum": 1
                },
                "price": {
                    "type": "number",
                    "minimum": 0
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "models.AlbumResponse": {
            "type": "object",
            "properties": {
                "artist": {
                    "type": "string"
                },
                "content": {
                    "type": "string"
                },
                "has_read": {
                    "type": "boolean"
                },
                "id": {
                    "type": "integer"
                },
                "price": {
                    "type": "number"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "models.Error": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Swagger ACY API",
	Description:      "This is a web api server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}