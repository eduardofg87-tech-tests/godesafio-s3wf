package user

import (
	"strings"
	"github.com/eduardofg87-tech-tests/godesafio-s3wf/backend/pkg/entity"
)

//IRepo in memory repo
type IRepo struct {
	m map[string]*entity.User
}

//NewInmemRepository create new repository
func NewInmemRepository() *IRepo {
	var m = map[string]*entity.User{}
	return &IRepo{
		m: m,
	}
}

//Store a User
func (r *IRepo) Store(a *entity.User) (entity.ID, error) {
	r.m[a.ID.String()] = a
	return a.ID, nil
}

//Find a User
func (r *IRepo) Find(id entity.ID) (*entity.User, error) {
	if r.m[id.String()] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[id.String()], nil
}

//Search Users
func (r *IRepo) Search(query string) ([]*entity.User, error) {
	var d []*entity.User
	for _, j := range r.m {
		if strings.Contains(strings.ToLower(j.Name), query) {
			d = append(d, j)
		}
	}
	if len(d) == 0 {
		return nil, entity.ErrNotFound
	}

	return d, nil
}

//FindAll Users
func (r *IRepo) FindAll() ([]*entity.User, error) {
	var d []*entity.User
	for _, j := range r.m {
		d = append(d, j)
	}
	return d, nil
}

//Delete a User
func (r *IRepo) Delete(id entity.ID) error {
	if r.m[id.String()] == nil {
		return entity.ErrNotFound
	}
	r.m[id.String()] = nil
	return nil
}
