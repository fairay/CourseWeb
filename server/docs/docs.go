// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
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
        "/accounts": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Accounts"
                ],
                "summary": "Creates a new account",
                "parameters": [
                    {
                        "description": "Account data",
                        "name": "account",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/objects.AccountDTO"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/objects.AccountDTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "Invalid"
                        }
                    }
                }
            }
        },
        "/accounts/login": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Accounts"
                ],
                "summary": "User's authorization",
                "parameters": [
                    {
                        "description": "Authentication data",
                        "name": "account",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/objects.AccountDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/objects.AccountDTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "Invalid"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "Authentication"
                        }
                    }
                }
            }
        },
        "/accounts/logout": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Accounts"
                ],
                "summary": "Logging out",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "Successful"
                        }
                    }
                }
            }
        },
        "/accounts/{login}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Accounts"
                ],
                "summary": "Retrieves account",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Account login",
                        "name": "login",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/objects.AccountDTO"
                        }
                    }
                }
            }
        },
        "/accounts/{login}/like": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Likes"
                ],
                "summary": "Retrieves all user's liked recipes",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Requested account",
                        "name": "login",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/objects.RecipeDTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "Invalid"
                        }
                    }
                }
            }
        },
        "/accounts/{login}/recipes": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Recipes"
                ],
                "summary": "Retrieves user's recipes",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Category title",
                        "name": "login",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/objects.RecipeDTO"
                            }
                        }
                    }
                }
            }
        },
        "/accounts/{login}/role": {
            "patch": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Accounts"
                ],
                "summary": "Change user's role",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Account login",
                        "name": "login",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Account data",
                        "name": "account",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/objects.AccountDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "Successful"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "Invalid"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "type": "Access"
                        }
                    }
                }
            }
        },
        "/categories": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Categories"
                ],
                "summary": "Retrieves all categories",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Search query",
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
                                "$ref": "#/definitions/objects.CategoryDTO"
                            }
                        }
                    }
                }
            }
        },
        "/categories/{title}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Categories"
                ],
                "summary": "Retrieves category by title",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Category title",
                        "name": "title",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/objects.CategoryDTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "Invalid"
                        }
                    }
                }
            }
        },
        "/categories/{title}/recipes": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Categories"
                ],
                "summary": "Retrieves all recipes at this category",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Searched category",
                        "name": "title",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/objects.RecipeDTO"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "Invalid"
                        }
                    }
                }
            }
        },
        "/recipes": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Recipes"
                ],
                "summary": "Retrieves all recipes",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/objects.RecipeDTO"
                            }
                        }
                    }
                }
            },
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Recipes"
                ],
                "summary": "Creates a new recipe",
                "parameters": [
                    {
                        "description": "Recipe data",
                        "name": "recipe",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/objects.RecipeDTO"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/objects.RecipeDTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "Invalid"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "Authentication"
                        }
                    }
                }
            }
        },
        "/recipes/{id}": {
            "delete": {
                "tags": [
                    "Recipes"
                ],
                "summary": "Deletes the recipe",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Recipe's id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "Successful"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "Invalid"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "Authentication"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "type": "Access"
                        }
                    }
                }
            }
        },
        "/recipes/{id}/categories": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Categories"
                ],
                "summary": "Retrieves all recipe's categories",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Recipe's id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/objects.CategoryDTO"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "Invalid"
                        }
                    }
                }
            },
            "put": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Categories"
                ],
                "summary": "Adds category to mentioned recipe",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Recipe's id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Category",
                        "name": "recipe",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/objects.CategoryDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "Successful"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "Invalid"
                        }
                    }
                }
            },
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Categories"
                ],
                "summary": "Sets categories to mentioned recipe",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Recipe's id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Categories",
                        "name": "categories",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/objects.CategoryDTO"
                            }
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "type": "Successful"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "Invalid"
                        }
                    }
                }
            },
            "delete": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Categories"
                ],
                "summary": "Removes category",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Recipe's id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Category",
                        "name": "recipe",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/objects.CategoryDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "Successful"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "Invalid"
                        }
                    }
                }
            }
        },
        "/recipes/{id}/ingredients": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Ingredients"
                ],
                "summary": "Retrieves all recipe's ingredients",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Recipe's id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/objects.IngredientDTO"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "Invalid"
                        }
                    }
                }
            },
            "put": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Ingredients"
                ],
                "summary": "Adds ingredient",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Recipe's id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Ingredient",
                        "name": "recipe",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/objects.IngredientDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            },
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Ingredients"
                ],
                "summary": "Posts all recipe's ingredients",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Recipe's id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Ingredients",
                        "name": "recipes",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/objects.IngredientDTO"
                            }
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "type": "Successful"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "Invalid"
                        }
                    }
                }
            },
            "delete": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Ingredients"
                ],
                "summary": "Removes ingredient",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Recipe's id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Ingredient",
                        "name": "recipe",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/objects.IngredientDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/recipes/{id}/like": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Likes"
                ],
                "summary": "Retrieves all users who liked mentioned recipe",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Recipe id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/objects.AccountDTO"
                            }
                        }
                    }
                }
            },
            "put": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Likes"
                ],
                "summary": "Adds like to recipe from authorized user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Recipe id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "type": "Successful"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "Invalid"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "Authentication"
                        }
                    }
                }
            },
            "delete": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Likes"
                ],
                "summary": "Deletes like to recipe from authorized user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Recipe id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "Successful"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "Invalid"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "Authentication"
                        }
                    }
                }
            }
        },
        "/recipes/{id}/like/amount": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Likes"
                ],
                "summary": "Retrieves the recipe's amount of likes",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Recipe id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "Invalid"
                        }
                    }
                }
            }
        },
        "/recipes/{id}/steps": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Steps"
                ],
                "summary": "Retrieves all steps",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Recipe's id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/objects.StepDTO"
                            }
                        }
                    }
                }
            },
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Steps"
                ],
                "summary": "Posts step",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Recipe's id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Step",
                        "name": "step",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/objects.StepDTO"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/objects.StepDTO"
                        }
                    }
                }
            }
        },
        "/recipes/{id}/steps/{step}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Steps"
                ],
                "summary": "Retrieves step",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Recipe's id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Step's number",
                        "name": "step",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/objects.StepDTO"
                        }
                    }
                }
            },
            "delete": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Steps"
                ],
                "summary": "Removes step",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Recipe's id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Step's number",
                        "name": "step",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            },
            "patch": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Steps"
                ],
                "summary": "Updates step",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Recipe's id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Step's number",
                        "name": "step",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/objects.StepDTO"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "objects.AccountDTO": {
            "type": "object",
            "properties": {
                "login": {
                    "type": "string",
                    "example": "admin"
                },
                "password": {
                    "type": "string",
                    "example": "0"
                },
                "role": {
                    "type": "string",
                    "example": "admin"
                }
            }
        },
        "objects.CategoryDTO": {
            "type": "object",
            "properties": {
                "title": {
                    "type": "string"
                }
            }
        },
        "objects.IngredientDTO": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "objects.RecipeDTO": {
            "type": "object",
            "properties": {
                "author": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "duration": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "portion_num": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "objects.StepDTO": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "num": {
                    "type": "integer"
                },
                "recipe": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BasicAuth": {
            "type": "basic"
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "0.1",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "Recipes API",
	Description: "API for culinary recipes (BMSTU Web development cource project)",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
