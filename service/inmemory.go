package service

import (
	"fmt"
	"time"

	"github.com/patrickmn/go-cache"

	"in-memory_location_server/model"
)

var dTTL = 600 * time.Second

type Repo struct {
	storage *cache.Cache
}

func NewRepo() Repo {
	return Repo{
		storage: cache.New(600*time.Second, 5*time.Minute),
	}
}

func (s *Repo) AddLocation(orderId string, location model.Location) {
	if data, ok := s.storage.Get(orderId); ok {
		if locationList, ok := data.([]model.Location); ok {
			locationList := append(locationList, location)
			s.storage.Set(orderId, locationList, dTTL)
			return
		}
	}
	s.storage.Set(orderId, []model.Location{location}, dTTL)
}

func (s *Repo) GetLocation(orderId string) ([]model.Location, error) {
	if data, ok := s.storage.Get(orderId); ok {
		if locationList, ok := data.([]model.Location); ok {
			return locationList, nil
		} else {
			return nil, fmt.Errorf("can not retrive the location list for order id %s, ", orderId)
		}
	}
	return nil, fmt.Errorf("can not retrive the location list for order id %s", orderId)
}

func (s *Repo) DeleteLocation(orderId string) error {
	if _, ok := s.storage.Get(orderId); !ok {
		return fmt.Errorf("the location list for order id %s does not exists", orderId)
	}
	s.storage.Delete(orderId)
	return nil
}
