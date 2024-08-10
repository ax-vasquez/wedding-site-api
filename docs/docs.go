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
        "/entree": {
            "post": {
                "description": "Create a new entree and return the new record's data to the caller",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "entrees"
                ],
                "summary": "create entree",
                "parameters": [
                    {
                        "description": "The input entree data (only ` + "`" + `option_name` + "`" + ` is required)",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Entree"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/controllers.V1_API_RESPONSE_ENTREE"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controllers.V1_API_RESPONSE_ENTREE"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controllers.V1_API_RESPONSE_ENTREE"
                        }
                    }
                }
            },
            "delete": {
                "description": "Deletes an entree and returns a response to indicate success or failure",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "entrees"
                ],
                "summary": "deletes an entree",
                "parameters": [
                    {
                        "type": "string",
                        "format": "uuid",
                        "description": "Entree ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "202": {
                        "description": "Accepted",
                        "schema": {
                            "$ref": "#/definitions/controllers.V1_API_RESPONSE_ENTREE"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controllers.V1_API_RESPONSE_ENTREE"
                        }
                    }
                }
            }
        },
        "/entrees": {
            "get": {
                "description": "Gets the selected entree for the given user ID (empty array if no selection has been made), or a list of all available entrees if no user ID is provided",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "entrees"
                ],
                "summary": "gets one or all entrees",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.V1_API_RESPONSE_ENTREE"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controllers.V1_API_RESPONSE_ENTREE"
                        }
                    }
                }
            }
        },
        "/horsdoeuvres": {
            "get": {
                "description": "Gets the selected hors doeuvres for the given user ID (empty array if no selection has been made), or a list of all available entrees if no user ID is provided",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "hors doeuvres"
                ],
                "summary": "gets one or all hors doeuvres",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.V1_API_RESPONSE_HORS_DOEUVRES"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controllers.V1_API_RESPONSE_HORS_DOEUVRES"
                        }
                    }
                }
            },
            "post": {
                "description": "Creates an hors doeuvres and return the new record's data to the caller",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "hors doeuvres"
                ],
                "summary": "creates an hors doeuvres",
                "parameters": [
                    {
                        "description": "The input hors doeuvres data (only ` + "`" + `option_name` + "`" + ` is required)",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.HorsDoeuvres"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/controllers.V1_API_RESPONSE_HORS_DOEUVRES"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controllers.V1_API_RESPONSE_HORS_DOEUVRES"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controllers.V1_API_RESPONSE_HORS_DOEUVRES"
                        }
                    }
                }
            },
            "delete": {
                "description": "Deletes an hors doeuvres and returns a response to indicate success or failure",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "hors doeuvres"
                ],
                "summary": "deletes an hors doeuvres",
                "parameters": [
                    {
                        "type": "string",
                        "format": "uuid",
                        "description": "Hors Doeuvres ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "202": {
                        "description": "Accepted",
                        "schema": {
                            "$ref": "#/definitions/controllers.V1_API_RESPONSE_HORS_DOEUVRES"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controllers.V1_API_RESPONSE_HORS_DOEUVRES"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controllers.V1_API_RESPONSE_HORS_DOEUVRES"
                        }
                    }
                }
            }
        },
        "/user": {
            "post": {
                "description": "Creates a user with the given input and returns an array of user objects, containing the newly-created user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "creates a user",
                "parameters": [
                    {
                        "description": "The input user data (only ` + "`" + `first_name` + "`" + `, ` + "`" + `last_name` + "`" + ` and ` + "`" + `email` + "`" + ` are required)",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/controllers.V1_API_RESPONSE_USERS"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controllers.V1_API_RESPONSE_USERS"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controllers.V1_API_RESPONSE_USERS"
                        }
                    }
                }
            },
            "delete": {
                "description": "Deletes an user and returns a response to indicate success or failure",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "deletes a user",
                "parameters": [
                    {
                        "type": "string",
                        "format": "uuid",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "202": {
                        "description": "Accepted",
                        "schema": {
                            "$ref": "#/definitions/controllers.V1_API_RESPONSE_USERS"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controllers.V1_API_RESPONSE_USERS"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controllers.V1_API_RESPONSE_USERS"
                        }
                    }
                }
            },
            "patch": {
                "description": "Updates a user with the given input",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "updates a user",
                "parameters": [
                    {
                        "description": "The input user update data (only ` + "`" + `id` + "`" + ` is required, but is not useful without setting other fields to update)",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                ],
                "responses": {
                    "202": {
                        "description": "Accepted",
                        "schema": {
                            "$ref": "#/definitions/controllers.V1_API_RESPONSE_USERS"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controllers.V1_API_RESPONSE_USERS"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controllers.V1_API_RESPONSE_USERS"
                        }
                    }
                }
            }
        },
        "/user/{user_id}/entrees": {
            "get": {
                "description": "Gets the selected entree for the given user ID (empty array if no selection has been made), or a list of all available entrees if no user ID is provided",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "entrees"
                ],
                "summary": "gets one or all entrees",
                "parameters": [
                    {
                        "type": "string",
                        "format": "uuid",
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
                            "$ref": "#/definitions/controllers.V1_API_RESPONSE_ENTREE"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controllers.V1_API_RESPONSE_ENTREE"
                        }
                    }
                }
            }
        },
        "/user/{user_id}/horsdoeuvres": {
            "get": {
                "description": "Gets the selected hors doeuvres for the given user ID (empty array if no selection has been made), or a list of all available entrees if no user ID is provided",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "hors doeuvres"
                ],
                "summary": "gets one or all hors doeuvres",
                "parameters": [
                    {
                        "type": "string",
                        "format": "uuid",
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
                            "$ref": "#/definitions/controllers.V1_API_RESPONSE_HORS_DOEUVRES"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controllers.V1_API_RESPONSE_HORS_DOEUVRES"
                        }
                    }
                }
            }
        },
        "/users": {
            "get": {
                "description": "Gets user(s) by the ID(s) in the request query string, ` + "`" + `?ids=` + "`" + `",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "gets user(s)",
                "parameters": [
                    {
                        "type": "string",
                        "format": "uuid",
                        "description": "user search by id (UUID)",
                        "name": "ids",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.V1_API_RESPONSE_USERS"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controllers.V1_API_RESPONSE_USERS"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controllers.EntreeData": {
            "type": "object",
            "properties": {
                "entrees": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Entree"
                    }
                }
            }
        },
        "controllers.HorsDoeuvresData": {
            "type": "object",
            "properties": {
                "hors_doeuvres": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.HorsDoeuvres"
                    }
                }
            }
        },
        "controllers.UserData": {
            "type": "object",
            "properties": {
                "users": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.User"
                    }
                }
            }
        },
        "controllers.V1_API_RESPONSE_ENTREE": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/controllers.EntreeData"
                },
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "controllers.V1_API_RESPONSE_HORS_DOEUVRES": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/controllers.HorsDoeuvresData"
                },
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "controllers.V1_API_RESPONSE_USERS": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/controllers.UserData"
                },
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "gin.H": {
            "type": "object",
            "additionalProperties": {}
        },
        "gorm.DeletedAt": {
            "type": "object",
            "properties": {
                "time": {
                    "type": "string"
                },
                "valid": {
                    "description": "Valid is true if Time is not NULL",
                    "type": "boolean"
                }
            }
        },
        "models.Entree": {
            "type": "object",
            "required": [
                "option_name"
            ],
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "deleted_at": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "id": {
                    "type": "string"
                },
                "option_name": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "models.HorsDoeuvres": {
            "type": "object",
            "required": [
                "option_name"
            ],
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "deleted_at": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "id": {
                    "type": "string"
                },
                "option_name": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "models.User": {
            "type": "object",
            "required": [
                "email",
                "first_name",
                "last_name"
            ],
            "properties": {
                "can_invite_others": {
                    "type": "boolean"
                },
                "createdAt": {
                    "description": "We override Gorm's CreatedAt field so we can set the gorm:\"\u003c-:create\" directive,\nwhich prevents this field from being altered once the record is created",
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "deleted_at": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "email": {
                    "type": "string"
                },
                "entreeSelection": {
                    "$ref": "#/definitions/models.Entree"
                },
                "entree_selection_id": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "horsDoeuvresSelection": {
                    "$ref": "#/definitions/models.HorsDoeuvres"
                },
                "hors_doeuvres_selection_id": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "is_admin": {
                    "type": "boolean"
                },
                "is_going": {
                    "type": "boolean"
                },
                "last_name": {
                    "type": "string"
                },
                "updated_at": {
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
