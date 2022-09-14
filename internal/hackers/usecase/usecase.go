package usecase

import "github.com/kapralovs/hackers-list/internal/hackers"

type HackerUseCase struct {
	hackerRepo hackers.Repository
}
