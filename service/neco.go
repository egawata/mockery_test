package service

import (
	"github.com/egawata/mockery_test/models"
	"github.com/egawata/mockery_test/repos"
)

type NecoService struct {
	dao repos.NecoRepo
}

func (s *NecoService) getNecoRepo() repos.NecoRepo {
	if s.dao == nil {
		s.dao = &DefaultNecoRepo{}
	}
	return s.dao
}

func (s *NecoService) GetNeco(id int) *models.Neco {
	return s.getNecoRepo().GetNeco(id)
}

type DefaultNecoRepo struct{}

func (s *DefaultNecoRepo) GetNeco(id int) *models.Neco {
	return &models.Neco{999, "Juri"}
}
