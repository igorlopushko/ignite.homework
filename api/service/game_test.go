package service

import (
	"testing"

	"github.com/igorlopushko/ignite.homework/api/internal/mock"
	"github.com/igorlopushko/ignite.homework/api/model"
)

func TestRun_AliensCountGreaterThenCities_ReturnsError(t *testing.T) {
	aliensCount := 2
	cities := make(map[string]*model.City)

	svc := &Game{
		Cities: cities,
	}

	err := svc.Run(aliensCount)

	if err == nil {
		t.Errorf("Run() method has to return an error")
	}
}

func TestRun_AliensSevReturnsErrorOnAlienGeneration_ReturnsError(t *testing.T) {
	aliensCount := 2

	cities := make(map[string]*model.City)
	cities["Bar"] = &model.City{Name: "Bar"}
	cities["Foo"] = &model.City{Name: "Foo"}

	alienSvc := mock.AlienMockSvc{MockCase: "case1"}

	svc := &Game{
		Cities:   cities,
		AlienSvc: alienSvc,
	}

	err := svc.Run(aliensCount)

	if err == nil {
		t.Errorf("Run() method has to return an error")
	}
}

func TestRun_NoCitiesAndAliensLeft_ReturnsError(t *testing.T) {
	aliensCount := -2
	cities := make(map[string]*model.City)
	alienSvc := mock.AlienMockSvc{MockCase: "case2"}

	svc := &Game{
		Cities:   cities,
		AlienSvc: alienSvc,
	}

	err := svc.Run(aliensCount)

	if err == nil {
		t.Errorf("Run() method has to return an error")
	}
}

func TestRun_GetRandomAlienReturnsError_ReturnsError(t *testing.T) {
	aliensCount := 1

	cities := make(map[string]*model.City)
	cities["Bar"] = &model.City{Name: "Bar"}
	cities["Foo"] = &model.City{Name: "Foo"}
	cities["Bee"] = &model.City{Name: "Bee"}

	alienSvc := mock.AlienMockSvc{MockCase: "case3"}

	svc := &Game{
		Cities:   cities,
		AlienSvc: alienSvc,
	}

	err := svc.Run(aliensCount)

	if err == nil {
		t.Errorf("Run() method has to return an error")
	}
}
