package model

import (
	"fmt"
	r "github.com/dancannon/gorethink"
	"github.com/materials-commons/contrib/schema"
)

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
