package service

import (
	"testing"

	"github.com/igorlopushko/ignite.homework/api/model"
)

func TestGenerateAliens_PathZeroCount_ReturnsEmptyMap(t *testing.T) {
	svc := AlienSvc{MaxStepsCount: 10}
	r, err := svc.GenerateAliens(nil, 0)

	if err != nil {
		t.Errorf("GenerateAliens() method does not have to return error")
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
		t.Errorf("GenerateAliens() method does not have to return error")
	}

	if len(r) != 2 {
		t.Errorf("GenerateAliens() method has to return a map of length 2")
	}
}
