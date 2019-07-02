package user

import "github.com/eduardofg87-tech-tests/godesafio-s3wf/backend/pkg/entity"

//Reader interface
type Reader interface {
	Find(id entity.ID) (*entity.User, error)
	Search(query string) ([]*entity.User, error)
	FindAll() ([]*entity.User, error)
}

//Writer User writer
type Writer interface {
	Store(b *entity.User) (entity.ID, error)
	Delete(id entity.ID) error
}

//Repository repository interface
type Repository interface {
	Reader
	Writer
}

//UseCase use case interface
type UseCase interface {
	Reader
	Writer
}
