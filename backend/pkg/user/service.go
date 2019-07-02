package user

import (
	"strings"
	"time"

	"github.com/eduardofg87-tech-tests/godesafio-s3wf/backend/pkg/entity"
)

//Service service interface
type Service struct {
	repo Repository
}

//NewService create new service
func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

//Store an user
func (s *Service) Store(b *entity.User) (entity.ID, error) {
	b.ID = entity.NewID()
	b.CreatedAt = time.Now()
	return s.repo.Store(b)
}

//Find an user
func (s *Service) Find(id entity.ID) (*entity.User, error) {
	return s.repo.Find(id)
}

//Search users
func (s *Service) Search(query string) ([]*entity.User, error) {
	return s.repo.Search(strings.ToLower(query))
}

//FindAll users
func (s *Service) FindAll() ([]*entity.User, error) {
	return s.repo.FindAll()
}

//Delete an user
func (s *Service) Delete(id entity.ID) error {
	b, err := s.Find(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}
