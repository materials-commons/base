package schema

import (
	"time"
)

// Directory models a directory of user files. A dir is an abstract representation
// of a users file system directory plus the metadata needed by the system.
type Directory struct {
	ID              string    `gorethink:"id,omitempty"`
	Access          string    `gorethink:"access" db:"-"`
	Owner           string    `gorethink:"owner" db:"-"`
	MarkedForReview bool      `gorethink:"marked_for_review" db:"-"`
	Name            string    `gorethink:"name"`
	DataFiles       []string  `gorethink:"datafiles" db:"-"`
	DataParams      []string  `gorethink:"dataparams" db:"-"`
	Users           []string  `gorethink:"users" db:"-"`
	Tags            []string  `gorethink:"tags" db:"-"`
	MyTags          []string  `gorethink:"mytags" db:"-"`
	Parent          string    `gorethink:"parent"`
	Reviews         []string  `gorethink:"reviews" db:"-"`
	Birthtime       time.Time `gorethink:"birthtime"`
	MTime           time.Time `gorethink:"mtime"`
	ATime           time.Time `gorethink:"atime"`
}

// NewDirectory creates a new Directory instance.
func NewDirectory(name, access, owner, parent string) Directory {
	now := time.Now()
	return Directory{
		Owner:     owner,
		Name:      name,
		Parent:    parent,
		Access:    access,
		Users:     []string{owner},
		Birthtime: now,
		MTime:     now,
		ATime:     now,
	}
}
