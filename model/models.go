package model

import (
	"github.com/materials-commons/base/schema"
)

// Groups is a default model for the usergroups table.
var Groups = &Model{
	schema: schema.UserGroup{},
	table:  "usergroups",
}
