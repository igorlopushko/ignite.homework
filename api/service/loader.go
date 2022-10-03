package service

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/igorlopushko/ignite.homework/api/model"
)

// A Direction represents direction type behavior.
type Direction int

const (
	KeyValueStringSize           = 2
	North              Direction = iota
	East
	South
	West
)

// Returns string representation of the direction enum.
func (d Direction) String() string {
	switch d {
	case North:
		return "north"
	case East:
		return "east"
	case South:
		return "south"
	case West:
		return "west"
	}
	return "unknown"
}

// A ILoaderService interface determines the data loader functionality.
type ILoaderService interface {
	Load(path string) (map[string]*model.City, error)
}

// A FileLoaderSrv is a representation of the service which performs data load from the file.
type FileLoaderSrv struct {
}

// Loads map data from the file.
func (s FileLoaderSrv) Load(path string) (map[string]*model.City, error) {
	// ASSUMPTION: file contains only one line per one city.
	// If there is a duplicate line for the same city the data will be overridden

	// open file
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	// remember to close the file at the end of the program
	defer f.Close()

	r := make(map[string]*model.City)

	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		v := strings.Split(scanner.Text(), " ")
		if len(v) == 0 {
			return nil, errors.New("map line has wrong format")
		}

		// create new city or get existing one from map
		var c *model.City
		if _, ok := r[v[0]]; ok {
			c = r[v[0]]
		} else {
			c = &model.City{
				Name:    v[0],
				AlienID: -1,
			}
		}

		// add directions
		if len(v) > 1 {
			for i := 1; i < len(v); i++ {
				r, err = addDirections(v[i], r, c)
				if err != nil {
					return nil, err
				}
			}
		}

		r[v[0]] = c
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return r, nil
}

func addDirections(v string, m map[string]*model.City, c *model.City) (map[string]*model.City, error) {
	d := strings.Split(v, "=")
	if len(d) != KeyValueStringSize {
		return nil, errors.New("map line has wrong format")
	}

	switch d[0] {
	case North.String():
		c.NorthDirection = d[1]
		if _, ok := m[d[1]]; ok {
			m[d[1]].SouthDirection = c.Name
			return m, nil
		}

		m[d[1]] = &model.City{
			Name:           d[1],
			AlienID:        -1,
			SouthDirection: c.Name,
		}
	case East.String():
		c.EastDirection = d[1]
		if _, ok := m[d[1]]; ok {
			m[d[1]].WestDirection = c.Name
			return m, nil
		}

		m[d[1]] = &model.City{
			Name:          d[1],
			AlienID:       -1,
			WestDirection: c.Name,
		}
	case South.String():
		c.SouthDirection = d[1]
		if _, ok := m[d[1]]; ok {
			m[d[1]].NorthDirection = c.Name
			return m, nil
		}

		m[d[1]] = &model.City{
			Name:           d[1],
			AlienID:        -1,
			NorthDirection: c.Name,
		}
	case West.String():
		c.WestDirection = d[1]
		if _, ok := m[d[1]]; ok {
			m[d[1]].EastDirection = c.Name
			return m, nil
		}

		m[d[1]] = &model.City{
			Name:          d[1],
			AlienID:       -1,
			EastDirection: c.Name,
		}
	default:
		return nil, fmt.Errorf("unknown direction '%s'", d[1])
	}

	return m, nil
}
