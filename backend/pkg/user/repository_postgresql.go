package user

import (
	"github.com/eduardofg87-tech-tests/godesafio-s3wf/backend/pkg/entity"
	
)

//PostgreSQLRepository PostgreSQLdb repo
type PostgreSQLRepository struct {
	pool *mgosession.Pool
	db   string
}

//NewPostgreSQLRepository create new repository
func NewPostgreSQLRepository(p *mgosession.Pool, db string) *PostgreSQLRepository {
	return &PostgreSQLRepository{
		pool: p,
		db:   db,
	}
}

//Find a user
func (r *PostgreSQLRepository) Find(id entity.ID) (*entity.Bookmark, error) {
	result := entity.Bookmark{}
	session := r.pool.Session(nil)
	coll := session.DB(r.db).C("user")
	err := coll.Find(bson.M{"_id": id}).One(&result)
	switch err {
	case nil:
		return &result, nil
	case mgo.ErrNotFound:
		return nil, entity.ErrNotFound
	default:
		return nil, err
	}
}

//Store a user
func (r *PostgreSQLRepository) Store(b *entity.Bookmark) (entity.ID, error) {
	session := r.pool.Session(nil)
	coll := session.DB(r.db).C("user")
	err := coll.Insert(b)
	if err != nil {
		return entity.ID(0), err
	}
	return b.ID, nil
}

//FindAll users
func (r *PostgreSQLRepository) FindAll() ([]*entity.Bookmark, error) {
	var d []*entity.Bookmark
	session := r.pool.Session(nil)
	coll := session.DB(r.db).C("user")
	err := coll.Find(nil).Sort("name").All(&d)
	switch err {
	case nil:
		return d, nil
	case mgo.ErrNotFound:
		return nil, entity.ErrNotFound
	default:
		return nil, err
	}
}

//Search users
func (r *PostgreSQLRepository) Search(query string) ([]*entity.Bookmark, error) {
	var d []*entity.Bookmark
	session := r.pool.Session(nil)
	coll := session.DB(r.db).C("user")
	err := coll.Find(bson.M{"name": &bson.RegEx{Pattern: query, Options: "i"}}).Limit(10).Sort("name").All(&d)
	switch err {
	case nil:
		return d, nil
	case mgo.ErrNotFound:
		return nil, entity.ErrNotFound
	default:
		return nil, err
	}
}

//Delete a user
func (r *PostgreSQLRepository) Delete(id entity.ID) error {
	session := r.pool.Session(nil)
	coll := session.DB(r.db).C("user")
	return coll.Remove(bson.M{"_id": id})
}
