package usecase

import "github.com/kapralovs/hackers-list/internal/hackers"

type Usecase struct {
	repo hackers.Repository
}

func New(r hackers.Repository) *Usecase {
	return &Usecase{
		repo: r,
	}
}
