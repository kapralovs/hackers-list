package localstorage

import (
	"sync"

	"github.com/kapralovs/hackers-list/internal/models"
)

type LocalStorage struct {
	hackers map[string]*models.Hacker
	mutex   *sync.Mutex
}

// func NewLocalStorage() *LocalStorage {
// 	return &LocalStorage{
// 		hackers: make(map[string]*models.Hacker),
// 		mutex:   new(sync.Mutex),
// 	}
// }

// func (ls *LocalStorage) GetHackersList() ([]*models.Hacker, error) {
// 	hackers := []*models.Hacker{}
// 	return hackers, nil
// }
