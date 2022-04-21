package service

import (
	"in-memory_location_server/model"
)

type LocationPayload struct {
	OrderId string           `json:"order_id"`
	History []model.Location `json:"history"`
}
