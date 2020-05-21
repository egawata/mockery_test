package repos

import (
	"github.com/egawata/mockery_test/models"
)

//go:generate mockery -name NecoRepo -output ../internal/mocks -case snake
type NecoRepo interface {
	GetNeco(id int) *models.Neco
}
