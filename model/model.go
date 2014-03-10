package model

import (
	"fmt"
	r "github.com/dancannon/gorethink"
	"github.com/materials-commons/base/db"
	"github.com/materials-commons/base/mc"
	"github.com/materials-commons/base/schema"
	"reflect"
)

// Model holds the schema definition and the table for the schema.
type Model struct {
	schema interface{}
	table  string
}

// Query holds the model and database references, such as the query to run
// and the database session.
type Query struct {
	*Model
	Rql     r.RqlTerm
	Session *r.Session
}

// ByID retrieves an entry by its id field.
func (q *Query) ByID(id string, obj interface{}) error {
	err := GetItem(id, q.table, q.Session, obj)
	return err
}

// Q constructs a Query and fills in its Session by calling db.RSession().
func (m *Model) Q() *Query {
	session, err := db.RSession()
	if err != nil {
		panic("Unable to connect to database")
	}
	return m.Qs(session)
}

// Qs constructs a query and accepts a database Session to use.
func (m *Model) Qs(session *r.Session) *Query {
	return &Query{
		Model:   m,
		Session: session,
		Rql:     r.Table(m.table),
	}
}

// Row returns a single item. It takes an arbitrary query.
func (q *Query) Row(query r.RqlTerm, obj interface{}) error {
	err := GetRow(query, q.Session, obj)
	return err
}

// Table returns the RqlTerm for the table. It abstracts away having to know the particular
// table for a given model.
func (m *Model) Table() r.RqlTerm {
	return r.Table(m.table)
}

// T is a shortcut for Table.
func (m *Model) T() r.RqlTerm {
	return r.Table(m.table)
}

// Rows returns a list of items from the database. It takes an arbitrary query.
func (q *Query) Rows(query r.RqlTerm, results interface{}) error {
	elementType := reflect.TypeOf(q.schema)
	resultsValue := reflect.ValueOf(results)
	if resultsValue.Kind() != reflect.Ptr || (resultsValue.Elem().Kind() != reflect.Slice && resultsValue.Elem().Kind() != reflect.Interface) {
		return fmt.Errorf("bad type for results")
	}

	sliceValue := resultsValue.Elem()

	if resultsValue.Elem().Kind() == reflect.Interface {
		sliceValue = sliceValue.Elem().Slice(0, sliceValue.Cap())
	} else {
		sliceValue = sliceValue.Slice(0, sliceValue.Cap())
	}

	rows, err := query.Run(q.Session)
	if err != nil {
		return err
	}
	defer rows.Close()

	i := 0
	for rows.Next() {
		var result = reflect.New(elementType)
		rows.Scan(result.Interface())
		if sliceValue.Len() == i {
			sliceValue = reflect.Append(sliceValue, result.Elem())
			sliceValue = sliceValue.Slice(0, sliceValue.Cap())
		} else {
			sliceValue.Index(i).Set(result.Elem())
		}
		i++
	}

	resultsValue.Elem().Set(sliceValue.Slice(0, i))
	return nil
}

// Update updates an existing database model entry.
func (q *Query) Update() error {
	return nil
}

// Insert inserts a new item into the model.
func (q *Query) Insert() error {
	return nil
}

// Delete deletes an existing database model entry.
func (q *Query) Delete() error {
	return nil
}

/*
func (q *Query) Exec() (id string, err error) {
	rw, err := q.RunWrite(q.session)
	switch {
	case err != nil:
		return err

	}
}
*/

/* ************************************************************** */

// The following are older functions that will be removed.

// MatchingUserGroups finds user groups matching on query.
func MatchingUserGroups(query r.RqlTerm, session *r.Session) ([]schema.UserGroup, error) {
	var results []schema.UserGroup
	rows, err := query.Run(session)
	if err != nil {
		return results, err
	}

	defer rows.Close()

	for rows.Next() {
		var ug schema.UserGroup
		rows.Scan(&ug)
		results = append(results, ug)
	}

	return results, nil
}

// GetDataFile retrieves an existing datafile by id.
func GetDataFile(id string, session *r.Session) (*schema.DataFile, error) {
	var df schema.DataFile
	if err := GetItem(id, "datafiles", session, &df); err != nil {
		return nil, err
	}
	return &df, nil
}

// GetUser retrieves and existing user by id.
func GetUser(id string, session *r.Session) (*schema.User, error) {
	var u schema.User
	if err := GetItem(id, "users", session, &u); err != nil {
		return nil, err
	}
	return &u, nil
}

// GetProject retrieves an existing project by id.
func GetProject(id string, session *r.Session) (*schema.Project, error) {
	var p schema.Project
	if err := GetItem(id, "projects", session, &p); err != nil {
		return nil, err
	}
	return &p, nil
}

// GetDataDir retrieves an existing datadir by id.
func GetDataDir(id string, session *r.Session) (*schema.DataDir, error) {
	var d schema.DataDir
	if err := GetItem(id, "datadirs", session, &d); err != nil {
		return nil, err
	}
	return &d, nil
}

// GetItem retrieves an item by id in the given table.
func GetItem(id, table string, session *r.Session, obj interface{}) error {
	result, err := r.Table(table).Get(id).RunRow(session)
	switch {
	case err != nil:
		return err
	case result.IsNil():
		return mc.ErrNotFound
	default:
		err := result.Scan(obj)
		return err
	}
}

// GetRow runs a query and returns a single item.
func GetRow(query r.RqlTerm, session *r.Session, obj interface{}) error {
	result, err := query.RunRow(session)
	switch {
	case err != nil:
		return err
	case result.IsNil():
		return mc.ErrNotFound
	default:
		err := result.Scan(obj)
		return err
	}
}

// GetRows runs a query an returns a list of results.
func GetRows(query r.RqlTerm, session *r.Session, results interface{}) error {
	resultsValue := reflect.ValueOf(results)
	if resultsValue.Kind() != reflect.Ptr || (resultsValue.Elem().Kind() != reflect.Slice && resultsValue.Elem().Kind() != reflect.Interface) {
		return fmt.Errorf("bad type for results")
	}

	sliceValue := resultsValue.Elem()

	if resultsValue.Elem().Kind() == reflect.Interface {
		sliceValue = sliceValue.Elem().Slice(0, sliceValue.Cap())
	} else {
		sliceValue = sliceValue.Slice(0, sliceValue.Cap())
	}
	elementType := sliceValue.Type().Elem()
	rows, err := query.Run(session)
	if err != nil {
		return err
	}
	defer rows.Close()
	i := 0
	for rows.Next() {
		var result = reflect.New(elementType)
		rows.Scan(result.Interface())
		if sliceValue.Len() == i {
			sliceValue = reflect.Append(sliceValue, result.Elem())
			sliceValue = sliceValue.Slice(0, sliceValue.Cap())
		} else {
			sliceValue.Index(i).Set(result.Elem())
		}
		i++
	}

	resultsValue.Elem().Set(sliceValue.Slice(0, i))
	return nil
}

// Delete deletes an item by id in the given table.
func Delete(table, id string, session *r.Session) error {
	_, err := r.Table(table).Get(id).Delete().RunWrite(session)
	return err
}
