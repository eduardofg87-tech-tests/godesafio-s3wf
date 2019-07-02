package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
	"github.com/GuilhermeCaruso/bellt"
	"github.com/eduardofg87-tech-tests/godesafio-s3wf/backend/api/handler"
	"github.com/eduardofg87-tech-tests/godesafio-s3wf/backend/config"
	"github.com/eduardofg87-tech-tests/godesafio-s3wf/backend/pkg/bookmark"
	"github.com/eduardofg87-tech-tests/godesafio-s3wf/backend/pkg/middleware"
)

func main() {
	router := bellt.NewRouter()

    // API v1
	// | /ping                 |  POST  |
	// | /auth                 |  POST  |
	// | /user                 |  POST  |
	// | /user                 |  PUT   |
	// | /user                 | DELETE |
	// | /user/{:uuid}         |  GET   |
	

	router.HandleGroup("/v1",
		router.SubHandleFunc("/ping", meetupevents.HandlerAllEvents, "POST"),
	)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Println("erro ao criar o server backend", err)
	}
}
