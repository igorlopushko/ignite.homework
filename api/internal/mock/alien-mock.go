// Package mock is implemented to represent mock objects for unit-testing.
package mock

import (
	"errors"

	"github.com/igorlopushko/ignite.homework/api/model"
)

const (
	case1 string = "case1"
	case2 string = "case2"
	case3 string = "case3"
	case4 string = "case4"
)

// A AlienMockSvc represents mock behavior for the alien service
type AlienMockSvc struct {
	MockCase string
}

// Returns values for the mock test cases.
func (s AlienMockSvc) GenerateAliens(m map[string]*model.City, aliensCount int) (map[int]*model.Alien, error) {
	switch s.MockCase {
	case case1:
		return nil, errors.New("error in case 1")
	case case2:
		aliens := make(map[int]*model.Alien)
		return aliens, nil
	case case3:
		aliens := make(map[int]*model.Alien)
		aliens[0] = &model.Alien{ID: 0, Name: "Alien0", Trapped: false}
		return aliens, nil
	case case4:
		aliens := make(map[int]*model.Alien)
		aliens[0] = &model.Alien{ID: 0, Name: "Alien0", Trapped: false}
		return aliens, nil
	}

	return nil, errors.New("error in undefined case")
}

// Returns values for the mock test cases.
func (s AlienMockSvc) GetRandomAlien(aliens map[int]*model.Alien) (*model.Alien, error) {
	switch s.MockCase {
	case case1:
		return nil, errors.New("error in case 1")
	case case2:
		return nil, errors.New("error in case 2")
	case case3:
		return nil, errors.New("error in case 3")
	case case4:
		a := &model.Alien{ID: 0, Name: "Alien0", Trapped: false}
		return a, nil
	}

	return nil, errors.New("error in undefined case")
}
