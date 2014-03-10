package model

import (
	"fmt"
	r "github.com/dancannon/gorethink"
	"github.com/materials-commons/base/schema"
	"testing"
)

var _ = fmt.Println

var (
	session, _ = r.Connect(map[string]interface{}{
		"address":  "localhost:30815",
		"database": "materialscommons",
	})
)

func TestGetUser(t *testing.T) {
	_, err := GetUser("nosuch@nosuch.com", session)
	if err == nil {
		t.Fatalf("Found non-existant user nosuch@nosuch.com")
	}

	u, err := GetUser("gtarcea@umich.edu", session)
	if err != nil {
		t.Fatalf("Didn't find existing user gtarcea@umich.edu: %s", err.Error())
	}

	if u.ApiKey != "472abe203cd411e3a280ac162d80f1bf" {
		t.Fatalf("ApiKey does not match, got %s", u.ApiKey)
	}
}

func TestGetUserModel(t *testing.T) {
	m := &Model{
		schema: schema.User{},
		table:  "users",
	}

	var user schema.User
	err := m.Qs(session).ByID("gtarcea@umich.edu", &user)
	if err != nil {
		t.Errorf("Lookup by Id failed: %s", err)
	}

	if user.Id != "gtarcea@umich.edu" {
		t.Errorf("Unexpected user return %#v", user)
	}

	var users []schema.User
	err = m.Qs(session).Rows(m.Table(), &users)
	if err != nil {
		t.Errorf("Lookup all users failed: %s", err)
	}

	if len(users) == 0 {
		t.Errorf("No users returned when looking up all users")
	}
}

func TestGetRows(t *testing.T) {
	var users []schema.User
	rql := r.Table("users")
	err := GetRows(rql, session, &users)
	if err != nil {
		t.Errorf("GetRows all users failed: %s", err)
	}

	if len(users) == 0 {
		t.Errorf("Users length == 0")
	}

	err = GetRows(rql, session, users)
	if err == nil {
		t.Errorf("Unexpected nil error when passing in bad parameter")
	}
}
