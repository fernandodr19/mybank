// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package swagger

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {
            "name": "MIT",
            "url": "https://opensource.org/licenses/MIT"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/transactions": {
            "post": {
                "description": "Process a transaction for a given account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Transactions"
                ],
                "summary": "Process a transaction",
                "parameters": [
                    {
                        "description": "Body",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/transactions.TransactionRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/transactions.TransactionResponse"
                        }
                    },
                    "400": {
                        "description": "Could not parse request"
                    },
                    "404": {
                        "description": "Account not found"
                    },
                    "422": {
                        "description": "Could not process transaction due to lack of balance or available credit"
                    },
                    "500": {
                        "description": "Internal server error"
                    }
                }
            }
        }
    },
    "definitions": {
        "transactions.TransactionRequest": {
            "type": "object",
            "properties": {
                "account_id": {
                    "type": "string"
                },
                "amount": {
                    "type": "integer"
                },
                "operation_type_id": {
                    "type": "integer"
                }
            }
        },
        "transactions.TransactionResponse": {
            "type": "object",
            "properties": {
                "transaction_id": {
                    "type": "string"
                }
            }
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
	Version:     "1.0",
	Host:        "localhost:3000",
	BasePath:    "/api/v1",
	Schemes:     []string{"http"},
	Title:       "Swagger Mybank API",
	Description: "Documentation Mybank API",
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
