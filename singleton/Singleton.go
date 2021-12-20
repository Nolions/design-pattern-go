package singleton

import "sync"

type Service struct {
}

var instance *Service
var once sync.Once

func NewService() *Service {
	once.Do(func() {
		instance = &Service{}
	})

	return instance
}
