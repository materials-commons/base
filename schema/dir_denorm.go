package schema

import (
	"time"
)

// FileEntry is a denormalized instance of a datafile used in the datadirs_denorm table.
type FileEntry struct {
	ID        string    `gorethink:"id"`
	Name      string    `gorethink:"name"`
	Owner     string    `gorethink:"owner"`
	Birthtime time.Time `gorethink:"birthtime"`
	Checksum  string    `gorethink:"checksum"`
	Size      int64     `gorethink:"size"`
}

// DataDirDenorm is a denormalized instance of a datadir used in the datadirs_denorm table.
type DataDirDenorm struct {
	ID        string      `gorethink:"id"`
	Name      string      `gorethink:"name"`
	Owner     string      `gorethink:"owner"`
	Birthtime time.Time   `gorethink:"birthtime"`
	DataFiles []FileEntry `gorethink:"datafiles"`
}
