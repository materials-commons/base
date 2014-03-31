package dir

// PatchType denotes the kind of patch operation
type PatchType int

const (
	// PatchCreate created item
	PatchCreate PatchType = iota

	// PatchDelete deleted item
	PatchDelete

	// PatchEdit item content was changed
	PatchEdit
)

// Patch is an instance of a difference when comparing two directories. It specifies
// the kind of change to apply. The list of patches implies changes to apply to
// the original directory to make it look like the new directory.
type Patch struct {
	File FileInfo
	Type PatchType
}
