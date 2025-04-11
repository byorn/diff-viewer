package dir_details

// FileNode represents a file or directory in a tree structure
type FileNode struct {
	Name     string     `json:"name"`
	IsDir    bool       `json:"isDir"`
	Children []FileNode `json:"children,omitempty"`
}
