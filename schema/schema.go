package schema

import (
	"time"
)

// UserGroup models user groups and access permissions to user data.
type UserGroup struct {
	ID          string    `gorethink:"id,omitempty"`
	Owner       string    `gorethink:"owner"`
	Name        string    `gorethink:"name"`
	Description string    `gorethink:"description"`
	Birthtime   time.Time `gorethink:"birthtime"`
	MTime       time.Time `gorethink:"mtime"`
	Access      string    `gorethink:"access"`
	Users       []string  `gorethink:"users"`
}

// NewUserGroup creates a new UserGroup instance.
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

// DataFile models a user datafile. A datafile is an abstract representation of
// a real file plus the attributes that we need in our model for access, and other
// metadata.
type DataFile struct {
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

// NewDataFile creates a new DataFile instance.
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

// Project models a users project. A project is an instance of a users workspace
// where they conduct their research. A project can be shared.
type Project struct {
	ID          string    `gorethink:"id,omitempty"`
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

// NewProject creates a new Project instance.
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

// DataDir models a directory of user files. A datadir is an abstract representation
// of a users file system directory plus the metadata needed by the system.
type DataDir struct {
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

// NewDataDir creates a new DataDir instance.
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

// User models a user in the system.
type User struct {
	ID          string    `gorethink:"id,omitempty"`
	Name        string    `gorethink:"name"`
	Email       string    `gorethink:"email"`
	Fullname    string    `gorethink:"fullname"`
	Password    string    `gorethink:"password"`
	APIKey      string    `gorethink:"apikey"`
	Birthtime   time.Time `gorethink:"birthtime"`
	MTime       time.Time `gorethink:"mtime"`
	Avatar      string    `gorethink:"avatar"`
	Description string    `gorethink:"description"`
	Affiliation string    `gorethink:"affiliation"`
	HomePage    string    `gorethink:"homepage"`
	Notes       []string  `gorethink:"notes"`
}

// NewUser creates a new User instance.
func NewUser(name, email, password, apikey string) User {
	now := time.Now()
	return User{
		ID:        email,
		Name:      name,
		Email:     email,
		Password:  password,
		APIKey:    apikey,
		Birthtime: now,
		MTime:     now,
	}
}

// Project2DataDir is a join table that maps projects to their datadirs.
type Project2DataDir struct {
	ID        string `gorethink:"id,omitempty" db:"-"`
	ProjectID string `gorethink:"project_id" db:"project_id"`
	DataDirID string `gorethink:"datadir_id" db:"datadir_id"`
}

// DataFileEntry is a denormalized instance of a datafile used in the datadirs_denorm table.
type DataFileEntry struct {
	ID        string    `gorethink:"id"`
	Name      string    `gorethink:"name"`
	Owner     string    `gorethink:"owner"`
	Birthtime time.Time `gorethink:"birthtime"`
	Checksum  string    `gorethink:"checksum"`
	Size      int64     `gorethink:"size"`
}

// DataDirDenorm is a denormalized instance of a datadir used in the datadirs_denorm table.
type DataDirDenorm struct {
	ID        string          `gorethink:"id"`
	Name      string          `gorethink:"name"`
	Owner     string          `gorethink:"owner"`
	Birthtime time.Time       `gorethink:"birthtime"`
	DataFiles []DataFileEntry `gorethink:"datafiles"`
}
