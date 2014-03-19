package schema

import (
	"time"
)

// File models a user file. A datafile is an abstract representation of
// a real file plus the attributes that we need in our model for access, and other
// metadata.
type File struct {
	ID              string    `gorethink:"id,omitempty"`
	Name            string    `gorethink:"name"`
	Access          string    `gorethink:"access" db:"-"`
	MarkedForReview bool      `gorethink:"marked_for_review" db:"-"`
	Reviews         []string  `gorethink:"reviews" db:"-"`
	Birthtime       time.Time `gorethink:"birthtime"`
	MTime           time.Time `gorethink:"mtime"`
	ATime           time.Time `gorethink:"atime"`
	Tags            []string  `gorethink:"tags" db:"-"`
	MyTags          []string  `gorethink:"mytags" db:"-"`
	Description     string    `gorethink:"description" db:"-"`
	Notes           []string  `gorethink:"notes" db:"-"`
	Owner           string    `gorethink:"owner" db:"-"`
	Process         string    `gorethink:"process" db:"-"`
	Machine         string    `gorethink:"machine" db:"-"`
	Checksum        string    `gorethink:"checksum"`
	Size            int64     `gorethink:"size"`
	Location        string    `gorethink:"location" db:"-"`
	MediaType       string    `gorethink:"mediatype" db:"-"`
	Conditions      []string  `gorethink:"conditions" db:"-"`
	Text            string    `gorethink:"text" db:"-"`
	MetaTags        []string  `gorethink:"metatags" db:"-"`
	DataDirs        []string  `gorethink:"datadirs" db:"-"`
	Parent          string    `gorethink:"parent"`
	UsesID          string    `gorethink:"usesid" db:"-"`
}

// NewFile creates a new File instance.
func NewFile(name, access, owner string) File {
	now := time.Now()
	return File{
		Name:        name,
		Access:      access,
		Owner:       owner,
		Description: name,
		Birthtime:   now,
		MTime:       now,
		ATime:       now,
	}
}
