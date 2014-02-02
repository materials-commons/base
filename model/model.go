package model

import (
	"fmt"
	r "github.com/dancannon/gorethink"
	"github.com/materials-commons/base/mc"
	"github.com/materials-commons/base/schema"
	"reflect"
)

type Model struct {
	schema interface{}
	table  string
}

type Query struct {
	*Model
	Rql     r.RqlTerm
	Session *r.Session
}

func (q *Query) ById(id string) (interface{}, error) {
	result := reflect.New(reflect.TypeOf(q.schema))
	err := GetItem(id, q.table, q.Session, result.Interface())
	return result.Interface(), err
}

func (m *Model) Q(session *r.Session) *Query {
	return &Query{
		Model:   m,
		Session: session,
		Rql:     r.Table(m.table),
	}
}

func (q *Query) Row(query r.RqlTerm) (interface{}, error) {
	result := reflect.New(reflect.TypeOf(q.schema))
	err := GetRow(query, q.Session, result.Interface())
	return result.Elem(), err
}

func (m *Model) Table() r.RqlTerm {
	return r.Table(m.table)
}

func (q *Query) All(query r.RqlTerm, results interface{}) error {
	elementType := reflect.TypeOf(q.schema)
	resultsValue := reflect.ValueOf(results)
	if resultsValue.Kind() != reflect.Ptr || (resultsValue.Elem().Kind() != reflect.Slice && resultsValue.Elem().Kind() != reflect.Interface) {
		return fmt.Errorf("Bad type for results")
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

func (q *Query) Update() error {
	return nil
}

func (q *Query) Insert() error {
	return nil
}

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

func GetDataFile(id string, session *r.Session) (*schema.DataFile, error) {
	var df schema.DataFile
	if err := GetItem(id, "datafiles", session, &df); err != nil {
		return nil, err
	}
	return &df, nil
}

func GetUser(id string, session *r.Session) (*schema.User, error) {
	var u schema.User
	if err := GetItem(id, "users", session, &u); err != nil {
		return nil, err
	}
	return &u, nil
}

func GetProject(id string, session *r.Session) (*schema.Project, error) {
	var p schema.Project
	if err := GetItem(id, "projects", session, &p); err != nil {
		return nil, err
	}
	return &p, nil
}

func GetDataDir(id string, session *r.Session) (*schema.DataDir, error) {
	var d schema.DataDir
	if err := GetItem(id, "datadirs", session, &d); err != nil {
		return nil, err
	}
	return &d, nil
}

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

func GetRows(query r.RqlTerm, session *r.Session, results interface{}) error {
	resultsValue := reflect.ValueOf(results)
	if resultsValue.Kind() != reflect.Ptr || (resultsValue.Elem().Kind() != reflect.Slice && resultsValue.Elem().Kind() != reflect.Interface) {
		return fmt.Errorf("Bad type for results")
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

func Delete(table, id string, session *r.Session) error {
	_, err := r.Table(table).Get(id).Delete().RunWrite(session)
	return err
}
