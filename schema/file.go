package schema

import (
	"time"
)

// File models a user file. A datafile is an abstract representation of a real file
// plus the attributes that we need in our model for access, and other metadata.
type File struct {
	ID          string    `gorethink:"id,omitempty"`
	Current     bool      `gorethink:"current"`
	Name        string    `gorethink:"name"`
	Birthtime   time.Time `gorethink:"birthtime"`
	MTime       time.Time `gorethink:"mtime"`
	ATime       time.Time `gorethink:"atime"`
	Description string    `gorethink:"description"`
	Notes       []string  `gorethink:"notes"`
	Owner       string    `gorethink:"owner"`
	Checksum    string    `gorethink:"checksum"`
	Size        int64     `gorethink:"size"`
	MediaType   string    `gorethink:"mediatype"`
	Parent      string    `gorethink:"parent"`
	UsesID      string    `gorethink:"usesid"`
	DataDirs    []string  `gorethink:"datadirs"`
}

// FileID returns the id to use for the file. Because files can be duplicates, all
// duplicates are stored under a single ID. UsesID is set to the ID that an entry
// points to when it is a duplicate.
func (f *File) FileID() string {
	if f.UsesID != "" {
		return f.UsesID
	}

	return f.ID
}

// NewFile creates a new File instance.
func NewFile(name, owner string) File {
	now := time.Now()
	return File{
		Name:        name,
		Owner:       owner,
		Description: "",
		Birthtime:   now,
		MTime:       now,
		ATime:       now,
		Current:     true,
	}
}
