package service

import (
	"errors"
	"fmt"

	random "github.com/igorlopushko/ignite.homework/api/internal/rand"
	"github.com/igorlopushko/ignite.homework/api/model"
	"github.com/sirupsen/logrus"
)

// A IAlienService interface which defines managing aliens behavior.
type IAlienService interface {
	GenerateAliens(m map[string]*model.City, aliensCount int) (map[int]*model.Alien, error)
	GetRandomAlien(aliens map[int]*model.Alien) (*model.Alien, error)
}

// A AlienSvc is an implementation of the IAlienService.
type AlienSvc struct {
	MaxStepsCount int
}

// Generates a map of aliens required for the game.
func (s AlienSvc) GenerateAliens(m map[string]*model.City, aliensCount int) (map[int]*model.Alien, error) {
	r := make(map[int]*model.Alien)
	if aliensCount <= 0 {
		return r, nil
	}

	for i := 0; i < aliensCount; i++ {
		c, err := getAlienStartLocation(m)
		if err != nil {
			return nil, err
		}

		r[i] = &model.Alien{
			ID:              i,
			Name:            fmt.Sprintf("Alien%d", i),
			StepsCount:      0,
			CurrentLocation: c.Name,
		}

		c.AlienID = i
	}

	return r, nil
}

// Selects random alien from the map to operate with.
func (s AlienSvc) GetRandomAlien(aliens map[int]*model.Alien) (*model.Alien, error) {
	// check number of trapped and made max steps available aliens
	invalidAliensCount := 0
	for _, v := range aliens {
		if v.Trapped || v.StepsCount >= s.MaxStepsCount {
			invalidAliensCount++
		}
	}
	if invalidAliensCount == len(aliens) {
		msg := "all aliens are invalid and can't make a step"
		logrus.Error(msg)
		return nil, errors.New(msg)
	}

	// find random alien
	for {
		k, err := random.GetRandomNumber(len(aliens))
		if err != nil {
			return nil, err
		}

		for _, a := range aliens {
			if k == 0 {
				if !a.Trapped && a.StepsCount < s.MaxStepsCount {
					return a, nil
				}

				break
			}
			k--
		}
	}
}

func getAlienStartLocation(m map[string]*model.City) (*model.City, error) {
	for {
		k, err := random.GetRandomNumber(len(m))
		if err != nil {
			return nil, err
		}

		for _, c := range m {
			if k == 0 {
				if c.AlienID == -1 {
					return c, nil
				}
				break
			}
			k--
		}
	}
}
