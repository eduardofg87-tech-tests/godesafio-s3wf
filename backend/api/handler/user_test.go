package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/codegangsta/negroni"
	"github.com/eduardofg87-tech-tests/godesafio-s3wf/backend/pkg/user"
	"github.com/eduardofg87-tech-tests/godesafio-s3wf/backend/pkg/entity"
	"github.com/GuilhermeCaruso/bellt"
	"github.com/stretchr/testify/assert"
)

func TestUserIndex(t *testing.T) {
	repo := user.NewInmemRepository()
	service := user.NewService(repo)
	r := mux.NewRouter()
	n := negroni.New()
	MakeUserHandlers(r, *n, service)
	path, err := r.GetRoute("userIndex").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/user", path)
	b := &entity.User{
		Name:        "Elton Minetto",
		Description: "Minetto's page",
		Link:        "http://www.eltonminetto.net",
		Tags:        []string{"golang", "php", "linux", "mac"},
		Favorite:    true,
	}
	_, _ = service.Store(b)
	ts := httptest.NewServer(userIndex(service))
	defer ts.Close()
	res, err := http.Get(ts.URL)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func TestUserIndexNotFound(t *testing.T) {
	repo := user.NewInmemRepository()
	service := user.NewService(repo)
	ts := httptest.NewServer(userIndex(service))
	defer ts.Close()
	res, err := http.Get(ts.URL + "?name=github")
	assert.Nil(t, err)
	assert.Equal(t, http.StatusNotFound, res.StatusCode)
}

func TestUserSearch(t *testing.T) {
	repo := user.NewInmemRepository()
	service := user.NewService(repo)
	b := &entity.User{
		Name:        "Elton Minetto",
		Description: "Minetto's page",
		Link:        "http://www.eltonminetto.net",
		Tags:        []string{"golang", "php", "linux", "mac"},
		Favorite:    true,
	}
	_, _ = service.Store(b)
	ts := httptest.NewServer(userIndex(service))
	defer ts.Close()
	res, err := http.Get(ts.URL + "?name=minetto")
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func TestUserAdd(t *testing.T) {
	repo := user.NewInmemRepository()
	service := user.NewService(repo)
	r := mux.NewRouter()
	n := negroni.New()
	MakeUserHandlers(r, *n, service)
	path, err := r.GetRoute("userAdd").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/user", path)

	h := userAdd(service)

	ts := httptest.NewServer(h)
	defer ts.Close()
	payload := fmt.Sprintf(`{
  "name": "Github",
  "description": "Github site",
  "link": "http://github.com",
  "tags": [
    "git",
    "social"
  ]
}`)
	resp, _ := http.Post(ts.URL+"/v1/user", "application/json", strings.NewReader(payload))
	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	var b *entity.User
	json.NewDecoder(resp.Body).Decode(&b)
	assert.True(t, entity.IsValidID(b.ID.String()))
	assert.Equal(t, "http://github.com", b.Link)
	assert.False(t, b.CreatedAt.IsZero())
}

func TestUserFind(t *testing.T) {
	repo := user.NewInmemRepository()
	service := user.NewService(repo)
	r := mux.NewRouter()
	n := negroni.New()
	MakeUserHandlers(r, *n, service)
	path, err := r.GetRoute("userFind").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/user/{id}", path)
	b := &entity.User{
		Name:        "Elton Minetto",
		Description: "Minetto's page",
		Link:        "http://www.eltonminetto.net",
		Tags:        []string{"golang", "php", "linux", "mac"},
		Favorite:    true,
	}
	bID, _ := service.Store(b)
	handler := userFind(service)
	r.Handle("/v1/user/{id}", handler)
	ts := httptest.NewServer(r)
	defer ts.Close()
	res, err := http.Get(ts.URL + "/v1/user/" + bID.String())
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	var d *entity.User
	json.NewDecoder(res.Body).Decode(&d)
	assert.NotNil(t, d)
	assert.Equal(t, bID, d.ID)
}

func TestUserRemove(t *testing.T) {
	repo := user.NewInmemRepository()
	service := user.NewService(repo)
	r := mux.NewRouter()
	n := negroni.New()
	MakeUserHandlers(r, *n, service)
	path, err := r.GetRoute("userDelete").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/user/{id}", path)
	b := &entity.User{
		Name:        "Elton Minetto",
		Description: "Minetto's page",
		Link:        "http://www.eltonminetto.net",
		Tags:        []string{"golang", "php", "linux", "mac"},
		Favorite:    false,
	}
	bID, _ := service.Store(b)
	handler := userDelete(service)
	req, _ := http.NewRequest("DELETE", "/v1/user/"+bID.String(), nil)
	r.Handle("/v1/user/{id}", handler).Methods("DELETE", "OPTIONS")
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}
