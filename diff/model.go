package diff

type FileContent struct {
	Lines []string
}

// ChangeType represents the type of change in the file comparison.
type ChangeType string

const (
	Added     ChangeType = "Added"
	Removed   ChangeType = "Removed"
	Modified  ChangeType = "Modified"
	Unchanged ChangeType = "Unchanged"
)

// ContentDiff represents changes in the content of a file.
type ContentDiff struct {
	LineNumber int        `json:"line_number"`        // Line number where the change occurred
	OldLine    string     `json:"old_line,omitempty"` // Original line (before change)
	NewLine    string     `json:"new_line,omitempty"` // New line (after change)
	Change     ChangeType `json:"change"`             // Type of change
}

// FileDiff represents differences between two versions of a file.
type FileDiff struct {
	FileName string        `json:"file_name"` // Name of the file being compared
	Status   ChangeType    `json:"status"`    // Status of the file (Added, Removed, Modified)
	Content  []ContentDiff `json:"content"`   // List of content changes
}

// FolderDiff represents differences between two directories.
type FolderDiff struct {
	FolderName string       `json:"folder_name"` // Name of the folder
	Status     ChangeType   `json:"status"`      // Status of the folder (Added, Removed, Modified)
	Files      []FileDiff   `json:"files"`       // List of file differences
	SubFolders []FolderDiff `json:"sub_folders"` // List of sub-folder differences
}
