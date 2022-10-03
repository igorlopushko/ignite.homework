package service

import (
	"testing"

	"github.com/igorlopushko/ignite.homework/api/model"
)

func TestGetRandomDirection_NoAvailableDirections_ReturnsError(t *testing.T) {
	m := make(map[string]*model.City)
	m["Foo"] = &model.City{Name: "Foo", AlienID: -1}

	svc := NavigationSvc{}
	_, err := svc.GetRandomDirection(m, "Foo")

	if err == nil {
		t.Errorf("GetRandomDirection() method has to return an error")
	}
}

func TestGetRandomDirection_AllDirectionsAreAvailable_ReturnsRandomDirection(t *testing.T) {
	m := make(map[string]*model.City)
	m["Foo"] = &model.City{
		Name:           "Foo",
		AlienID:        -1,
		NorthDirection: "Bar",
		SouthDirection: "Bee",
		EastDirection:  "Qu-ux",
		WestDirection:  "Baz"}
	m["Bar"] = &model.City{Name: "Bar"}
	m["Bee"] = &model.City{Name: "Bee"}
	m["Qu-ux"] = &model.City{Name: "Qu-ux"}
	m["Baz"] = &model.City{Name: "Baz"}

	svc := NavigationSvc{}
	r, err := svc.GetRandomDirection(m, "Foo")

	if err != nil {
		t.Errorf("GetRandomDirection() method does not have to return an error")
	}

	if r == "" {
		t.Errorf("GetRandomDirection() method has to return valid direction, but returned an empty string")
	}
}
