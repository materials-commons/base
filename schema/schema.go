package schema

import (
	"time"
)

type UserGroup struct {
	Id          string    `gorethink:"id,omitempty"`
	Owner       string    `gorethink:"owner"`
	Name        string    `gorethink:"name"`
	Description string    `gorethink:"description"`
	Birthtime   time.Time `gorethink:"birthtime"`
	MTime       time.Time `gorethink:"mtime"`
	Access      string    `gorethink:"access"`
	Users       []string  `gorethink:"users"`
}

func NewUserGroup(owner, name string) UserGroup {
	now := time.Now()
	return UserGroup{
		Owner:       owner,
		Name:        name,
		Description: name,
		Access:      "private",
		Birthtime:   now,
		MTime:       now,
	}
}

type DataFile struct {
	Id              string    `gorethink:"id,omitempty"`
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
	Notes           []string  `gorethink:"description" db:"-"`
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

func NewDataFile(name, access, owner string) DataFile {
	now := time.Now()
	return DataFile{
		Name:        name,
		Access:      access,
		Owner:       owner,
		Description: name,
		Birthtime:   now,
		MTime:       now,
		ATime:       now,
	}
}

type Project struct {
	Id          string    `gorethink:"id,omitempty"`
	Name        string    `gorethink:"name"`
	Description string    `gorethink:"description"`
	DataDir     string    `gorethink:"datadir" db:"-"`
	Owner       string    `gorethink:"owner" db:"-"`
	Birthtime   time.Time `gorethink:"birthtime"`
	MTime       time.Time `gorethink:"mtime"`
	Notes       []string  `gorethink:"notes" db:"-"`
	Tags        []string  `gorethink:"tags" db:"-"`
	Reviews     []string  `gorethink:"reviews" db:"-"`
	MyTags      []string  `gorethink:"mytags" db:"-"`
}

func NewProject(name, datadir, owner string) Project {
	now := time.Now()
	return Project{
		Name:      name,
		DataDir:   datadir,
		Owner:     owner,
		Birthtime: now,
		MTime:     now,
	}
}

type DataDir struct {
	Id              string    `gorethink:"id,omitempty"`
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

func NewDataDir(name, access, owner, parent string) DataDir {
	now := time.Now()
	return DataDir{
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

type User struct {
	Id          string    `gorethink:"id,omitempty"`
	Name        string    `gorethink:"name"`
	Email       string    `gorethink:"email"`
	Fullname    string    `gorethink:"fullname"`
	Password    string    `gorethink:"password"`
	ApiKey      string    `gorethink:"apikey"`
	Birthtime   time.Time `gorethink:"birthtime"`
	MTime       time.Time `gorethink:"mtime"`
	Avatar      string    `gorethink:"avatar"`
	Description string    `gorethink:"description"`
	Affiliation string    `gorethink:"affiliation"`
	HomePage    string    `gorethink:"homepage"`
	Notes       []string  `gorethink:"notes"`
}

func NewUser(name, email, password, apikey string) User {
	now := time.Now()
	return User{
		Name:      name,
		Email:     email,
		Password:  password,
		ApiKey:    apikey,
		Birthtime: now,
		MTime:     now,
	}
}

// Join table structures
type Project2DataDir struct {
	Id        string `gorethink:"id,omitempty" db:"-"`
	ProjectID string `gorethink:"project_id" db:"project_id"`
	DataDirID string `gorethink:"datadir_id" db:"datadir_id"`
}
