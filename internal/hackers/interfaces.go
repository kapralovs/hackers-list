package hackers

import "github.com/kapralovs/hackers-list/internal/models"

type Usecase interface {
	GetHackersList() ([]*models.Hacker, error)
}
type Repository interface {
	GetHackersList() ([]*models.Hacker, error)
}
