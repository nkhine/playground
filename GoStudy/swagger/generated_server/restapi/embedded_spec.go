// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Swagger 2.0 study - a fridge API",
    "title": "Fridge API",
    "version": "0.0.1"
  },
  "host": "localhost:8080",
  "basePath": "/",
  "paths": {
    "/products": {
      "get": {
        "description": "List all products in the fridge",
        "produces": [
          "application/json"
        ],
        "parameters": [
          {
            "type": "boolean",
            "default": false,
            "description": "Sort products alphabetically A to Z",
            "name": "sort",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "List products success",
            "schema": {
              "$ref": "#/definitions/Products"
            }
          }
        }
      },
      "post": {
        "description": "Insert more of a product into the fridge",
        "consumes": [
          "application/json"
        ],
        "parameters": [
          {
            "name": "product",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/Product"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Insert product success"
          }
        }
      }
    },
    "/products/{name}": {
      "get": {
        "description": "Get single product information",
        "produces": [
          "application/json"
        ],
        "parameters": [
          {
            "type": "string",
            "description": "Name of the product to get information about",
            "name": "name",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Get products success",
            "schema": {
              "$ref": "#/definitions/Product"
            }
          },
          "404": {
            "description": "No such product in the fridge"
          }
        }
      },
      "put": {
        "description": "Withdraw given amount of given product from the fridge",
        "consumes": [
          "application/json"
        ],
        "parameters": [
          {
            "type": "string",
            "description": "Name of the product to withdraw from the fridge",
            "name": "name",
            "in": "path",
            "required": true
          },
          {
            "description": "Amount of the product to withdraw from the fridge",
            "name": "product",
            "in": "body",
            "schema": {
              "type": "object",
              "properties": {
                "quantity": {
                  "type": "number",
                  "default": 1,
                  "example": 1
                }
              }
            }
          }
        ],
        "responses": {
          "204": {
            "description": "Product successfuly withdrawn from the fridge"
          },
          "404": {
            "description": "No such product in the fridge"
          }
        }
      }
    }
  },
  "definitions": {
    "Product": {
      "description": "Single product",
      "type": "object",
      "properties": {
        "name": {
          "$ref": "#/definitions/ProductName"
        },
        "quantity": {
          "type": "number",
          "title": "Quantity of the product; pieces/liters/packages/etc",
          "example": 0.5
        }
      }
    },
    "ProductName": {
      "type": "string",
      "minLength": 1,
      "example": "Milk"
    },
    "Products": {
      "description": "List of products",
      "type": "array",
      "items": {
        "$ref": "#/definitions/Product"
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Swagger 2.0 study - a fridge API",
    "title": "Fridge API",
    "version": "0.0.1"
  },
  "host": "localhost:8080",
  "basePath": "/",
  "paths": {
    "/products": {
      "get": {
        "description": "List all products in the fridge",
        "produces": [
          "application/json"
        ],
        "parameters": [
          {
            "type": "boolean",
            "default": false,
            "description": "Sort products alphabetically A to Z",
            "name": "sort",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "List products success",
            "schema": {
              "$ref": "#/definitions/Products"
            }
          }
        }
      },
      "post": {
        "description": "Insert more of a product into the fridge",
        "consumes": [
          "application/json"
        ],
        "parameters": [
          {
            "name": "product",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/Product"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Insert product success"
          }
        }
      }
    },
    "/products/{name}": {
      "get": {
        "description": "Get single product information",
        "produces": [
          "application/json"
        ],
        "parameters": [
          {
            "type": "string",
            "description": "Name of the product to get information about",
            "name": "name",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Get products success",
            "schema": {
              "$ref": "#/definitions/Product"
            }
          },
          "404": {
            "description": "No such product in the fridge"
          }
        }
      },
      "put": {
        "description": "Withdraw given amount of given product from the fridge",
        "consumes": [
          "application/json"
        ],
        "parameters": [
          {
            "type": "string",
            "description": "Name of the product to withdraw from the fridge",
            "name": "name",
            "in": "path",
            "required": true
          },
          {
            "description": "Amount of the product to withdraw from the fridge",
            "name": "product",
            "in": "body",
            "schema": {
              "type": "object",
              "properties": {
                "quantity": {
                  "type": "number",
                  "default": 1,
                  "minimum": 0,
                  "example": 1
                }
              }
            }
          }
        ],
        "responses": {
          "204": {
            "description": "Product successfuly withdrawn from the fridge"
          },
          "404": {
            "description": "No such product in the fridge"
          }
        }
      }
    }
  },
  "definitions": {
    "Product": {
      "description": "Single product",
      "type": "object",
      "properties": {
        "name": {
          "$ref": "#/definitions/ProductName"
        },
        "quantity": {
          "type": "number",
          "title": "Quantity of the product; pieces/liters/packages/etc",
          "minimum": 0,
          "example": 0.5
        }
      }
    },
    "ProductName": {
      "type": "string",
      "minLength": 1,
      "example": "Milk"
    },
    "Products": {
      "description": "List of products",
      "type": "array",
      "items": {
        "$ref": "#/definitions/Product"
      }
    }
  }
}`))
}