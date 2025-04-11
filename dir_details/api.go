package dir_details

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
)

// generateDirTree generates a tree structure for the given directory
func generateDirTree(dirPath string) (FileNode, error) {
	root := FileNode{}
	info, err := os.Stat(dirPath)
	if err != nil {
		return root, err
	}

	root.Name = info.Name()
	root.IsDir = info.IsDir()

	// If it's a directory, recursively fetch its contents
	if info.IsDir() {
		entries, err := os.ReadDir(dirPath)
		if err != nil {
			return root, err
		}

		for _, entry := range entries {
			childPath := filepath.Join(dirPath, entry.Name())
			childNode, err := generateDirTree(childPath)
			if err != nil {
				return root, err
			}
			root.Children = append(root.Children, childNode)
		}
	}

	return root, nil
}

// DirTreeHandler handles the "/dir_tree" endpoint
func DirTreeHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the query parameter 'dir'
	dirPath := r.URL.Query().Get("dir")
	if dirPath == "" {
		http.Error(w, "Missing 'dir' query parameter", http.StatusBadRequest)
		return
	}

	// Generate the directory tree
	tree, err := generateDirTree(dirPath)
	if err != nil {
		http.Error(w, "Error reading directory: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Convert the tree to JSON and send the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tree)
}
