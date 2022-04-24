package model

import (
	v "github.com/go-ozzo/ozzo-validation/v4"
)

type Location struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

func (c Location) Validate() error {
	return v.ValidateStruct(&c)
}
