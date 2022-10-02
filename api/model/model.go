package model

type City struct {
	Name           string
	NorthDirection string
	EastDirection  string
	SouthDirection string
	WestDirection  string
	AlienID        int
}

type Alien struct {
	ID              int
	Name            string
	CurrentLocation string
	StepsCount      int
	Trapped         bool
}
