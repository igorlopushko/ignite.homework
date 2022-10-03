package service

import (
	"errors"

	random "github.com/igorlopushko/ignite.homework/api/internal/rand"
	"github.com/igorlopushko/ignite.homework/api/model"
)

// A INavigationService interface which defines managing navigation across the world map.
type INavigationService interface {
	GetRandomDirection(m map[string]*model.City, n string) (string, error)
}

// A NavigationSvc is an implementation of the INavigationService.
type NavigationSvc struct {
}

// Defines a string of random direction to navigate across the map.
func (s NavigationSvc) GetRandomDirection(m map[string]*model.City, n string) (string, error) {
	// check if there are any available directions
	availableDirections := getAvailableDirections(m, n)
	if len(availableDirections) == 0 {
		return "", errors.New("no available directions")
	}

	k, err := random.GetRandomNumber(len(availableDirections))
	if err != nil {
		return "", err
	}

	return availableDirections[k], nil
}

func getAvailableDirections(m map[string]*model.City, n string) []string {
	// TODO: update directions????
	d := make([]string, 0)
	if _, ok := m[n]; !ok {
		return d
	}

	c := m[n]
	if c.NorthDirection != "" {
		if _, ok := m[c.NorthDirection]; ok {
			d = append(d, c.NorthDirection)
		}
	}
	if c.EastDirection != "" {
		if _, ok := m[c.EastDirection]; ok {
			d = append(d, c.EastDirection)
		}
	}
	if c.SouthDirection != "" {
		if _, ok := m[c.SouthDirection]; ok {
			d = append(d, c.SouthDirection)
		}
	}
	if c.WestDirection != "" {
		if _, ok := m[c.WestDirection]; ok {
			d = append(d, c.WestDirection)
		}
	}

	return d
}
