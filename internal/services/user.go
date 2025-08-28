package services

import (
	"order_orbit/internal/entities"
	"order_orbit/internal/storage"
)

type UserService struct {
	storage *storage.UserStorage
}

func NewUserService(storage *storage.UserStorage) *UserService {
	return &UserService{
		storage: storage,
	}
}

func (s *UserService) FetchAll() []entities.User {
	return s.storage.All()
}

func (s *UserService) FetchById(id string) entities.User {
	return s.storage.Find(id)
}

func (s *UserService) Create(attributes map[string]any) entities.User {
	return s.storage.Create(attributes)
}

func (s *UserService) Delete(id string) {
	s.storage.Delete(id)
}

func (s *UserService) Update(id string, attributes map[string]any) entities.User {
	return s.storage.Update(id, attributes)
}
