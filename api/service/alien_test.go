package service

import (
	"testing"

	"github.com/igorlopushko/ignite.homework/api/model"
)

func TestGenerateAliens_PathZeroCount_ReturnsEmptyMap(t *testing.T) {
	svc := AlienSvc{MaxStepsCount: 10}
	r, err := svc.GenerateAliens(nil, 0)

	if err != nil {
		t.Errorf("GenerateAliens() method does not have to return an error")
	}

	if len(r) > 0 {
		t.Errorf("GenerateAliens() method does not have to return map with values")
	}
}

func TestGenerateAliens_NonZeroCount_ReturnsAliensMap(t *testing.T) {
	i := 2
	m := make(map[string]*model.City)
	m["Bar"] = &model.City{Name: "Bar", AlienID: -1}
	m["Foo"] = &model.City{Name: "Foo", AlienID: -1}
	svc := AlienSvc{MaxStepsCount: 10}
	r, err := svc.GenerateAliens(m, i)

	if err != nil {
		t.Errorf("GenerateAliens() method does not have to return an error")
	}

	if len(r) != 2 {
		t.Errorf("GenerateAliens() method has to return a map of length 2")
	}
}

func TestGetRandomAlien_AllAliensAreTrappedOrMaxStepsMade_RerunsError(t *testing.T) {
	aliens := make(map[int]*model.Alien)
	aliens[0] = &model.Alien{ID: 0, Trapped: true, Name: "alien0"}
	aliens[0] = &model.Alien{ID: 1, Trapped: false, StepsCount: 10}
	svc := AlienSvc{MaxStepsCount: 10}
	_, err := svc.GetRandomAlien(aliens)

	if err == nil {
		t.Errorf("GetRandomAlien() method has to return an error")
	}
}

func TestGetRandomAlien_ValidAliens_ReturnsRandomAlien(t *testing.T) {
	aliens := make(map[int]*model.Alien)
	aliens[0] = &model.Alien{ID: 0, Trapped: false, Name: "alien0"}
	aliens[1] = &model.Alien{ID: 1, Trapped: false, StepsCount: 10}
	svc := AlienSvc{MaxStepsCount: 10}
	r, err := svc.GetRandomAlien(aliens)

	if err != nil {
		t.Errorf("GetRandomAlien() method does not have to return an error")
	}

	if r.Name != "alien0" {
		t.Errorf("GetRandomAlien() method has to return alien with name 'alien0'")
	}
}
