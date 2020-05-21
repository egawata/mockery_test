package service

import (
	"github.com/egawata/mockery_test/internal/mocks"
	"github.com/egawata/mockery_test/models"
	"testing"
)

func TestGetNeco(t *testing.T) {
	neco := models.Neco{123, "Hanii"}

	var m mocks.NecoRepo
	m.On("GetNeco", neco.ID).Return(&neco)

	s := NecoService{&m}
	res := s.GetNeco(123)
	t.Logf("res = %#v", res)
	if res.Name != "Hanii" {
		t.Errorf("unexpected name. want Hanii got %v", res.Name)
	}
}
