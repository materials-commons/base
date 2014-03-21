package protocol

import (
	"github.com/materials-commons/base/mc"
	"time"
)

// MessageType identifies the type of message. This is prepended to the encoded
// buffer and is used by the receiver to identify the type of message it expects
// to decode.
type MessageType uint8

const (
	// LoginRequestMessage LoginRequest message
	LoginRequestMessage = iota

	// LoginResponseMessage LoginResponse message
	LoginResponseMessage

	// LogoutRequestMessage LogoutRequest message
	LogoutRequestMessage

	// CreateProjectRequestMessage CreateProjectRequest message
	CreateProjectRequestMessage

	// CreateProjectResponseMessage CreateProjectResponse message
	CreateProjectResponseMessage

	// CreateDirectoryRequestMessage CreateDirectoryRequest message
	CreateDirectoryRequestMessage

	// CreateDirectoryResponseMessage CreateDirectoryResponse message
	CreateDirectoryResponseMessage

	// CreateFileRequestMessage CreateFileRequest message
	CreateFileRequestMessage

	// CreateFileResponseMessage CreateFileResponse message
	CreateFileResponseMessage

	// DirectoryStatRequestMessage DirectoryStatRequest message
	DirectoryStatRequestMessage

	// DirectoryStatResponseMessage DirectoryStatResponse message
	DirectoryStatResponseMessage
)

// Status is the status of the request. All response type messages include a request status.
// This denotes success or failure of the request. On failure a message with additional details
// may be included.
type Status struct {
	Status  mc.ErrorCode
	Message string
}

// LoginRequest is sent to login to the server and prepare to issue commands.
type LoginRequest struct {
	User   string
	APIKey string
}

// LoginResponse will inform the client whether the login request was successful.
type LoginResponse struct {
	Status
}

// LogoutRequest is sent to end a session and disconnect. There client doesn't wait for
// a response from the server. It terminates the connection after sending this request.
type LogoutRequest struct{}

// CreateProjectRequest is sent to create a project.
type CreateProjectRequest struct {
	// The name of the project to create
	Name string

	// If shared is set to true then the server will check if a project matching
	// this name exists in any of the projects user has access to. If so it will
	// use that project.
	Shared bool
}

// CreateProjectResponse is the response to a CreateProjectRequest. A request to create
// a project will not create a project if a matching project already exists. In that case
// it will indicate this in the Status field.
type CreateProjectResponse struct {
	// The status of the request. There are two error codes for success:
	//    ErrorCodeSuccess - Project was created
	//    ErrorCodeExists  - Project already exists
	// All other error codes are failures.
	Status Status

	// The internal ProjectID for the created or existing project.
	ProjectID string

	// The internal id for the directory that the project is stored in.
	DirectoryID string
}

// CreateDirectoryRequest is sent to create a directory in a project.
type CreateDirectoryRequest struct {
	// The project to create the directory in.
	ProjectID string

	// The directory path, relative to the project to create. All members of the path
	// except the leaf must exist.
	Path string
}

// CreateDirectoryResponse is the response a CreateDirectoryRequest. A request to create
// a directory will not create a directory if a matching directory already exists. In
// that case it will indicate this in the Status field.
type CreateDirectoryResponse struct {
	// The status of the request. There are two error codes for success:
	//    ErrorCodeSuccess - Directory was created
	//    ErrorCodeExists  - Directory already exists
	// All other error codes are failures.
	Status

	// The internal id of the directory.
	DirectoryID string
}

// CreateFileRequest is sent to create a new file on the server. It is expected that
// after this request is sent that an attempt will be made to upload the file. Until
// this upload succeeds newer versions of the file cannot be created.
type CreateFileRequest struct {
	// The id of the project to create the file in.
	ProjectID string

	// The id of the directory in the project to create the file in.
	DirectoryID string

	// The name of the file.
	Name string

	// The files MD5 hash
	Checksum string

	// The size of the file
	Size int64

	// If CreateNewVersion is set to true then a new file version will
	// be created if the previous file has already been successfully
	// uploaded.
	CreateNewVersion bool
}

// CreateFileResponse is the response to a CreateFileRequest.
type CreateFileResponse struct {
	// Status of the request. There are three error codes for success:
	//    ErrorCodeSuccess - File was created
	//    ErrorCodeExists  - File already exists
	//    ErrorCodeNew     - A new version of the file was created
	// All other error codes are failures.
	Status

	// The internal id of the file.
	FileID string
}

// DirectoryStatRequest requests the server to send back its current
// view of the given directory.
type DirectoryStatRequest struct {
	ProjectID   string
	DirectoryID string
}

// StatEntryType identifies the type of stat entry as either a file
// or a directory.
type StatEntryType uint8

const (
	// StatTypeDirectory entry is a directory
	StatTypeDirectory StatEntryType = iota

	// StatTypeFile entry is a file
	StatTypeFile
)

// StatEntry describes a single file or directory
type StatEntry struct {
	// The type being described. If Type is StatTypeDirectory, then
	// Checksum, Size, and UploadedSize are not defined.
	Type StatEntryType

	// Name of entry
	Name string

	// The internal ID of the entry
	ID string

	// Owner of entry
	Owner string

	// The computed MD5 Hash
	Checksum string

	// The size of the file
	Size int64

	// The actual uploaded size of the file
	UploadedSize int64

	// Datetime the entry was created on the server.
	Birthtime time.Time
}

// DirectoryStatResponse is the response for a DirectoryStatRequest. It returns
// a list of all the known entries in the directory.
type DirectoryStatResponse struct {
	// Status of the request
	Status

	// ProjectID passed in the request
	ProjectID string

	// DirectoryID passed in the request
	DirectoryID string

	// A list of all the entries (file and directory) that the server
	// knows about for this directory.
	Entries []StatEntry
}
