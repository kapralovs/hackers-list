package usecase

import (
	"github.com/kapralovs/hackers-list/internal/hackers"
	"github.com/kapralovs/hackers-list/internal/models"
)

type Usecase struct {
	repo hackers.Repository
}

func New(r hackers.Repository) *Usecase {
	return &Usecase{
		repo: r,
	}
}

// func (uc *Usecase) GetHackersList() ([]*models.Hacker, error) {
// 	hackers, err := uc.repo.GetHackersList()
// 	if err != nil {

// 	}
// 	return hackers, nil
// }

func (uc *Usecase) GetHackersList() []*models.Hacker {
	hackers := uc.repo.GetHackersList()
	return hackers
}
