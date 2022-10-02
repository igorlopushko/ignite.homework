package service

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"

	"github.com/igorlopushko/ignite.homework/api/config"
	"github.com/igorlopushko/ignite.homework/api/model"
	"github.com/sirupsen/logrus"
)

const (
	gameOverMsg = "game over"
)

type Game struct {
	Cities map[string]*model.City
	Aliens map[int]*model.Alien
}

func (g *Game) Run(aliensCount int) error {
	var err error

	// ASSUMPTION: there could not be more aliens than number of cities
	if aliensCount > len(g.Cities) {
		logrus.Error("aliens count is greater than number of available cities")
		return errors.New(gameOverMsg)
	}

	g.Aliens, err = g.generateAliens(aliensCount)
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
		a, err := g.getRandomAlien()
		if err != nil {
			return errors.New(gameOverMsg)
		}

		// make step
		g.step(a)
	}
}

func (g *Game) generateAliens(aliensCount int) (map[int]*model.Alien, error) {
	r := make(map[int]*model.Alien)

	for i := 0; i < aliensCount; i++ {
		c, err := g.getAlienStartLocation()
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

func (g *Game) getAlienStartLocation() (*model.City, error) {
	for {
		k, err := getRandomNumber(len(g.Cities))
		if err != nil {
			return nil, err
		}

		for _, c := range g.Cities {
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

func (g *Game) getRandomAlien() (*model.Alien, error) {
	// check number of trapped and made max steps available aliens
	invalidAliensCount := 0
	for _, v := range g.Aliens {
		if v.Trapped || v.StepsCount >= config.App.AlienMaxStepsNumber {
			invalidAliensCount++
		}
	}
	if invalidAliensCount == len(g.Aliens) {
		msg := "all aliens are invalid and can't make a step"
		logrus.Error(msg)
		return nil, errors.New(msg)
	}

	// find random alien
	for {
		k, err := getRandomNumber(len(g.Aliens))
		if err != nil {
			return nil, err
		}

		for _, a := range g.Aliens {
			if k == 0 {
				if !a.Trapped && a.StepsCount < config.App.AlienMaxStepsNumber {
					return a, nil
				}

				break
			}
			k--
		}
	}
}

func (g *Game) step(a *model.Alien) {
	d, err := g.getRandomDirection(a.CurrentLocation)
	if err != nil {
		a.Trapped = true
		logrus.Error(fmt.Printf("no available steps for '%s' alien", a.Name))
		return
	}

	// leave the city by alien
	g.Cities[a.CurrentLocation].AlienID = -1

	if g.Cities[d].AlienID != -1 {
		// if the destination city is already occupied delete city and both
		logrus.Warn(fmt.Printf("%s has been destroyed by %s and %s!", g.Cities[d].Name, a.Name, g.Cities[d].Name))

		delete(g.Aliens, a.ID)
		delete(g.Aliens, g.Cities[d].AlienID)
		delete(g.Cities, g.Cities[d].Name)
	} else {
		// acquire new city by alien
		a.CurrentLocation = g.Cities[d].Name

		// TODO: update directions????
		g.Cities[d].AlienID = a.ID
		a.StepsCount++
	}
}

func (g *Game) getRandomDirection(n string) (string, error) {
	// check if there are any available directions
	availableDirections := g.getAvailableDirections(n)
	if len(availableDirections) == 0 {
		return "", errors.New("no available directions")
	}

	k, err := getRandomNumber(len(availableDirections))
	if err != nil {
		return "", err
	}

	return availableDirections[k], nil
}

func (g *Game) getAvailableDirections(n string) []string {
	// TODO: update directions???
	d := make([]string, 0)
	if _, ok := g.Cities[n]; !ok {
		return d
	}

	c := g.Cities[n]
	if c.NorthDirection != "" {
		if _, ok := g.Cities[c.NorthDirection]; ok {
			d = append(d, c.NorthDirection)
		}
	}
	if c.EastDirection != "" {
		if _, ok := g.Cities[c.EastDirection]; ok {
			d = append(d, c.EastDirection)
		}
	}
	if c.SouthDirection != "" {
		if _, ok := g.Cities[c.SouthDirection]; ok {
			d = append(d, c.SouthDirection)
		}
	}
	if c.WestDirection != "" {
		if _, ok := g.Cities[c.WestDirection]; ok {
			d = append(d, c.WestDirection)
		}
	}

	return d
}

func getRandomNumber(l int) (int64, error) {
	r, err := rand.Int(rand.Reader, big.NewInt(int64(l)))
	if err != nil {
		return 0, err
	}
	return r.Int64(), nil
}
