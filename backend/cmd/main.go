package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/eduardofg87-tech-tests/godesafio-s3wf/backend/config"
	"github.com/eduardofg87-tech-tests/godesafio-s3wf/backend/pkg/bookmark"
	"github.com/eduardofg87-tech-tests/godesafio-s3wf/backend/pkg/entity"
	"github.com/juju/mgosession"
	mgo "gopkg.in/mgo.v2"
)

func handleParams() (string, error) {
	if len(os.Args) < 2 {
		return "", errors.New("Invalid query")
	}
	return os.Args[1], nil
}

func main() {
	query, err := handleParams()
	if err != nil {
		log.Fatal(err.Error())
	}

	session, err := mgo.Dial(config.MONGODB_HOST)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer session.Close()

	mPool := mgosession.NewPool(nil, session, config.MONGODB_CONNECTION_POOL)
	defer mPool.Close()

	bookmarkRepo := bookmark.NewMongoRepository(mPool, config.MONGODB_DATABASE)
	bookmarkService := bookmark.NewService(bookmarkRepo)
	all, err := bookmarkService.Search(query)
	if err != nil {
		log.Fatal(err)
	}
	if len(all) == 0 {
		log.Fatal(entity.ErrNotFound.Error())
	}
	for _, j := range all {
		fmt.Printf("%s %s %v \n", j.Name, j.Link, j.Tags)
	}
}
