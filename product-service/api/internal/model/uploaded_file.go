package model

// FileType represents the status of the uploaded file type
type FileType string

const (
	// FileTypeImage means the file type is image
	FileTypeImage FileType = "IMAGE"
	// FileTypeVideo means the file type is video
	FileTypeVideo FileType = "VIDEO"
	// FileTypeDocument means the file type is document
	FileTypeDocument FileType = "DOCUMENT"
)

// String returns string type of custom type
func (c FileType) String() string {
	return string(c)
}
