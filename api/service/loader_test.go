package service

import "testing"

func TestLoad_WrongFileName_ReturnsError(t *testing.T) {
	svc := &FileLoaderSrv{}
	_, err := svc.Load("wrong.txt")

	if err == nil {
		t.Errorf("Load() method has to return an error")
	}
}

func TestLoad_CorrectFileName_ReturnsMap(t *testing.T) {
	svc := &FileLoaderSrv{}
	m, err := svc.Load("../../tests/map_test.txt")

	if err != nil {
		t.Errorf("Load() method does not have to return an error")
	}

	if len(m) != 5 {
		t.Errorf("Load() method has to return map with '5' values, contains '%d'", len(m))
	}
}
