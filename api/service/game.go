// Package service is implemented to represent alien invasion simulation.
package service

import (
	"errors"
	"fmt"

	"github.com/igorlopushko/ignite.homework/api/model"
	"github.com/sirupsen/logrus"
)

const (
	gameOverMsg = "game over"
)

// A Game represents game object behavior.
type Game struct {
	Cities        map[string]*model.City
	Aliens        map[int]*model.Alien
	AlienSvc      IAlienService
	NavigationSvc INavigationService
}

// Executes alien invasion simulation process.
func (g *Game) Run(aliensCount int) error {
	var err error

	// ASSUMPTION: there could not be more aliens than number of cities
	if aliensCount > len(g.Cities) {
		logrus.Error("aliens count is greater than number of available cities")
		return errors.New(gameOverMsg)
	}

	g.Aliens, err = g.AlienSvc.GenerateAliens(g.Cities, aliensCount)
	if err != nil {
		return err
	}

	for {
		// if no cities or aliens left then exit
		if len(g.Cities) == 0 || len(g.Aliens) == 0 {
			logrus.Error("there are no either cities or aliens left")
			return errors.New(gameOverMsg)
		}

		// get random alien to make a step
		a, err := g.AlienSvc.GetRandomAlien(g.Aliens)
		if err != nil {
			return errors.New(gameOverMsg)
		}

		// make step
		g.step(a)
	}
}

func (g *Game) step(a *model.Alien) {
	d, err := g.NavigationSvc.GetRandomDirection(g.Cities, a.CurrentLocation)
	if err != nil {
		a.Trapped = true
		logrus.Error(fmt.Printf("no available steps for '%s' alien", a.Name))
		return
	}

	// leave the city by alien
	g.Cities[a.CurrentLocation].AlienID = -1

	if g.Cities[d].AlienID != -1 {
		// if the destination city is already occupied delete city and both
		logrus.Warn(fmt.Sprintf("%s has been destroyed by %s and %s!", g.Cities[d].Name, a.Name, g.Aliens[g.Cities[d].AlienID].Name))

		delete(g.Aliens, a.ID)
		delete(g.Aliens, g.Cities[d].AlienID)
		delete(g.Cities, g.Cities[d].Name)
	} else {
		// acquire new city by alien
		a.CurrentLocation = g.Cities[d].Name

		g.Cities[d].AlienID = a.ID
		a.StepsCount++
	}
}
