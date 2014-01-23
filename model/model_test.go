package model

import (
	"fmt"
	r "github.com/dancannon/gorethink"
	"github.com/materials-commons/contrib/schema"
	"testing"
)

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

	u, err := m.Q(session).ById("gtarcea@umich.edu")
	fmt.Println("err =", err)
	fmt.Printf("%#v\n", u)

	if true {
		return
	}

	var users []schema.User
	err = m.Q(session).All(m.Table(), &users)
	for _, user := range users {
		fmt.Printf("\n\nuser = %#v\n", user)
	}
	//fmt.Println(err)
	//fmt.Printf("%#v\n", users)
}

func TestArray(t *testing.T) {
	items := make([]int, 3)
	fillIt(items)
	fmt.Printf("%#v", items)
}

func fillIt(results []int) {
	fmt.Println("cap", cap(results))
	fmt.Println("len", len(results))
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		if i >= len(results) {
			fmt.Println("Doing append")
			results = append(results, i)
			fmt.Println("Past append")
			fmt.Println("new cap", cap(results))
			fmt.Println("new len", len(results))
		} else {
			fmt.Println("Assigning to results", i)
			results[i] = i
		}
		//results = append(results, i)
	}
}
