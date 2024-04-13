// Package models provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.1.0 DO NOT EDIT.
package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Cart defines model for Cart.
type Cart struct {
	Items *[]struct {
		Quantity *int  `json:"quantity,omitempty"`
		Yarn     *Yarn `json:"yarn,omitempty"`
	} `json:"items,omitempty"`
}

// Yarn defines model for Yarn.
type Yarn struct {
	Available *bool               `json:"available,omitempty"`
	ProductId *primitive.ObjectID `bson:"_id,omitempty" json:"product_id,omitempty"`
	UnitColor *string             `json:"unit_color,omitempty"`
	UnitName  *string             `json:"unit_name,omitempty"`
	UnitPrice *float32            `json:"unit_price,omitempty"`
}

// PatchCartProductIdJSONBody defines parameters for PatchCartProductId.
type PatchCartProductIdJSONBody struct {
	Amount *int `json:"amount,omitempty"`
}

// PostCartJSONRequestBody defines body for PostCart for application/json ContentType.
type PostCartJSONRequestBody = Yarn

// PatchCartProductIdJSONRequestBody defines body for PatchCartProductId for application/json ContentType.
type PatchCartProductIdJSONRequestBody PatchCartProductIdJSONBody
