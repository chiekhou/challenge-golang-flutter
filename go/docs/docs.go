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
        "/Signup": {
            "post": {
                "description": "Create a new user with the provided information",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Allow you to register as a new User",
                "parameters": [
                    {
                        "description": "User data",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.SignupRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "User created",
                        "schema": {
                            "$ref": "#/definitions/requests.SignupRequest"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "404": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "409": {
                        "description": "Conflict",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            }
        },
        "/forgotten_password": {
            "post": {
                "description": "Envoie un mail de reset de mot de passe",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Envoie un mail à un user existant afin de réinitialiser son mot de passe",
                "parameters": [
                    {
                        "description": "User data",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.EmailRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Connexion réussie",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "404": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "409": {
                        "description": "Conflict",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "description": "login to the app",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Allow you to log and have an JWT Token",
                "parameters": [
                    {
                        "description": "User data",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Connexion réussie",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "404": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "409": {
                        "description": "Conflict",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            }
        },
        "/logout": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Delete your token session",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Logout",
                "parameters": [
                    {
                        "type": "string",
                        "default": "Bearer \u003cAdd access token here\u003e",
                        "description": "Insert your access token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Déconnexion réussie",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "404": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "409": {
                        "description": "Conflict",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            }
        },
        "/profile": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Retourne les informations du profil de l'utilisateur connecté",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Récupère le profil de l'utilisateur actuellement connecté",
                "parameters": [
                    {
                        "type": "string",
                        "default": "Bearer \u003cAdd access token here\u003e",
                        "description": "Insert your access token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            }
        },
        "/reset_password": {
            "put": {
                "description": "Permet à l'utilisateur de réinitialiser son mot de passe en utilisant un jeton de réinitialisation valide.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Réinitialiser le mot de passe",
                "parameters": [
                    {
                        "description": "Données pour réinitialiser le mot de passe",
                        "name": "requests.ResetPasswordRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.ResetPasswordRequest"
                        }
                    }
                ],
                "responses": {
                    "204": {
                        "description": "Mot de passe réinitialisé avec succès"
                    },
                    "400": {
                        "description": "Token invalide ou expiré",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "Erreur interne du serveur",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            }
        },
        "/send_invitation": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Envoie un mail d'invitation afin de de rejoindre un groupen de voyage",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Voyage"
                ],
                "summary": "Rejoindre un groupe de voyage",
                "parameters": [
                    {
                        "type": "string",
                        "default": "Bearer \u003cAdd access token here\u003e",
                        "description": "Insert your access token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Invitation envoyée",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "404": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "409": {
                        "description": "Conflict",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "gin.H": {
            "type": "object",
            "additionalProperties": {}
        },
        "requests.EmailRequest": {
            "type": "object",
            "required": [
                "email"
            ],
            "properties": {
                "email": {
                    "type": "string"
                }
            }
        },
        "requests.LoginRequest": {
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
        "requests.ResetPasswordRequest": {
            "type": "object",
            "required": [
                "new_password",
                "token"
            ],
            "properties": {
                "new_password": {
                    "type": "string"
                },
                "token": {
                    "type": "string"
                }
            }
        },
        "requests.SignupRequest": {
            "type": "object",
            "required": [
                "address",
                "email",
                "first_name",
                "last_name",
                "password",
                "username"
            ],
            "properties": {
                "address": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "username": {
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
