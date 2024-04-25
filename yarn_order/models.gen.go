// Package main provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.1.0 DO NOT EDIT.
package main

import (
	openapi_types "github.com/oapi-codegen/runtime/types"
)

// Order defines model for order.
type Order struct {
	CartId  *string             `json:"cart_id,omitempty"`
	OrderId *openapi_types.UUID `json:"order_id,omitempty"`
	UserId  *string             `json:"user_id,omitempty"`
}
