package migrations

import (
	migrate "github.com/eminetto/mongo-migrate"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

func init() {
	migrate.Register(func(db *mgo.Database) error { //Up
		return db.C("user").Update(
			bson.M{"email": "admin@eduardofg.dev"},
			bson.M{"$set": bson.M{"name": "Code:Nation Admin User"}})

	}, func(db *mgo.Database) error { //Down
		return db.C("user").Update(
			bson.M{"email": "admin@eduardofg.dev"},
			bson.M{"$set": bson.M{"name": "Admin"}})
	})
}
