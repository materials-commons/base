package model

import (
	"github.com/materials-commons/base/schema"
)

var Groups = &Model{
	schema: schema.UserGroup{},
	table:  "usergroups",
}
