// Package model is implemented to represent domain models.
package model

// A City represents the city on the map with all possible directions.
type City struct {
	Name           string
	NorthDirection string
	EastDirection  string
	SouthDirection string
	WestDirection  string
	AlienID        int
}

// An Alien represents alien object essential data.
type Alien struct {
	ID              int
	Name            string
	CurrentLocation string
	StepsCount      int
	Trapped         bool
}
