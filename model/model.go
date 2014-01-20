package model

import (
	"fmt"
	r "github.com/dancannon/gorethink"
	"github.com/materials-commons/contrib/schema"
	"reflect"
)

type Model struct {
	schema interface{}
	table  string
}

type Query struct {
	*Model
	Rql r.RqlTerm
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
		Rql: r.Table(m.table),
	}
}

func (q *Query) Row(query r.RqlTerm) (interface{}, error) {
	result := reflect.New(reflect.TypeOf(q.schema))
	err := GetRow(query, q.Session, result.Interface())
	return result.Interface(), err
}

func (m *Model) Table() r.RqlTerm {
	return r.Table(m.table)
}

func (q *Query) All(query r.RqlTerm, results interface{}) error {
	resultsType := reflect.ValueOf(results)
	var _ = resultsType
	//if resultsType.Kind() != reflect.Ptr || (resultsType.Elem().Kind() != reflect.Slice && resultsType.Elem().Kind() != reflect.Interface) {
	//	fmt.Printf("Bad type for results")
		//return fmt.Errorf("Bad type for results")
	//}
	/*
	rows, err := query.Run(q.Session)
	if err != nil {
		return err
	}
*/
	//slicev := resultsType.Interface()
	//fmt.Printf("Here %#v", slicev)
	elementType := reflect.TypeOf(q.schema)
	//fmt.Printf("%T\n", elementType)
	//resultsSlice := reflect.SliceOf(elementType.Elem())
	//fmt.Printf("%T\n", resultsSlice)

	rows, err := query.Run(q.Session)
	if err != nil {
		return err
	}
	defer rows.Close()
	i := 0
	for rows.Next() {
		var result = reflect.New(elementType)
		rows.Scan(result.Interface())
		//results = append(results, result.Interface())
		//results = append(results, result.Interface())
		i++
	}

	fmt.Printf("%#v", results)
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
		return fmt.Errorf("Unknown Id: %s", id)
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
		return fmt.Errorf("Bad query")
	default:
		err := result.Scan(obj)
		return err
	}
}

func Delete(table, id string, session *r.Session) error {
	_, err := r.Table(table).Get(id).Delete().RunWrite(session)
	return err
}
