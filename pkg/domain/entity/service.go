package entity

import (
	"sync"
)

var serviceOnce sync.Once

type Service struct {
	repository *Repository
}

func NewItemService(repository *Repository) (service *Service) {
	serviceOnce.Do(func() {
		service = &Service{repository}
	})
	return
}

/*
	TODO: Define services.

	Example:
		func (s *Service) SERVICE() ([]DTO, error) {
			value, err := s.repository.BEHAVE()
			if err != nil {
				return nil, err
			}

			return dto.Mapper(value), nil
		}
*/
