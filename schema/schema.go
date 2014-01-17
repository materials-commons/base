package schema

import (
	"strings"
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
	Access          string    `gorethink:"access"`
	MarkedForReview bool      `gorethink:"marked_for_review"`
	Reviews         []string  `gorethink:"reviews"`
	Birthtime       time.Time `gorethink:"birthtime"`
	MTime           time.Time `gorethink:"mtime"`
	ATime           time.Time `gorethink:"atime"`
	Tags            []string  `gorethink:"tags"`
	MyTags          []string  `gorethink:"mytags"`
	Description     string    `gorethink:"description"`
	Notes           []string  `gorethink:"description"`
	Owner           string    `gorethink:"owner"`
	Process         string    `gorethink:"process"`
	Machine         string    `gorethink:"machine"`
	Checksum        string    `gorethink:"checksum"`
	Size            int64     `gorethink:"size"`
	Location        string    `gorethink:"location"`
	MediaType       string    `gorethink:"mediatype"`
	Conditions      []string  `gorethink:"conditions"`
	Text            string    `gorethink:"text"`
	MetaTags        []string  `gorethink:"metatags"`
	DataDirs        []string  `gorethink:"datadirs"`
	Parent          string    `gorethink:"parent"`
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
	DataDir     string    `gorethink:"datadir"`
	Owner       string    `gorethink:"owner"`
	Birthtime   time.Time `gorethink:"birthtime"`
	MTime       time.Time `gorethink:"mtime"`
	Notes       []string  `gorethink:"notes"`
	Tags        []string  `gorethink:"tags"`
	Reviews     []string  `gorethink:"reviews"`
	MyTags      []string  `gorethink:"mytags"`
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
	Access          string    `gorethink:"access"`
	Owner           string    `gorethink:"owner"`
	MarkedForReview bool      `gorethink:"marked_for_review"`
	Name            string    `gorethink:"name"`
	DataFiles       []string  `gorethink:"datafiles"`
	DataParams      []string  `gorethink:"dataparams"`
	Users           []string  `gorethink:"users"`
	Tags            []string  `gorethink:"tags"`
	MyTags          []string  `gorethink:"mytags"`
	Parent          string    `gorethink:"parent"`
	Reviews         []string  `gorethink:"reviews"`
	Birthtime       time.Time `gorethink:"birthtime"`
	MTime           time.Time `gorethink:"mtime"`
	ATime           time.Time `gorethink:"atime"`
}

func NewDataDir(name, access, owner, parent string) DataDir {
	now := time.Now()
	return DataDir{
		Id:        owner + "$" + strings.Replace(name, "/", "_", -1),
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
