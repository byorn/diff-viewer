package diff

import (
	"io/fs"
	"os"
	"path/filepath"
)

func TestFunc(iterator int) FolderDiff {
	fd := FolderDiff{
		FolderName: "sample1",
		Status:     Unchanged,
	}
	fd.Files = append(fd.Files, FileDiff{
		FileName: "sample1.txt",
		Status:   Modified,
		Content:  []ContentDiff{},
	})
	if iterator == 1 {
		return fd
	} else {
		fd.SubFolders = append(fd.SubFolders, TestFunc(1))
	}
	return fd
}

// CompareFolders recursively compares two folders
func CompareFolders(oldPath, newPath string) FolderDiff {
	folderDiff := FolderDiff{
		FolderName: filepath.Base(newPath),
		Status:     Unchanged,
	}

	oldFiles := getFileMap(oldPath)
	newFiles := getFileMap(newPath)

	// Compare files
	for fileName, oldFilePath := range oldFiles {
		if newFilePath, exists := newFiles[fileName]; exists {
			// File exists in both folders, check for modifications
			contentDiffs := CompareFiles(oldFilePath, newFilePath)
			if len(contentDiffs) > 0 {
				folderDiff.Files = append(folderDiff.Files, FileDiff{
					FileName: fileName,
					Status:   Modified,
					Content:  contentDiffs,
				})
			}
		} else {
			// File removed
			folderDiff.Files = append(folderDiff.Files, FileDiff{
				FileName: fileName,
				Status:   Removed,
			})
		}
	}

	// Check for added files
	for fileName, _ := range newFiles {
		if _, exists := oldFiles[fileName]; !exists {
			folderDiff.Files = append(folderDiff.Files, FileDiff{
				FileName: fileName,
				Status:   Added,
			})
		}
	}

	// Compare subfolders
	oldSubFolders := getFolderMap(oldPath)
	newSubFolders := getFolderMap(newPath)

	for subFolderName, oldSubFolderPath := range oldSubFolders {
		if newSubFolderPath, exists := newSubFolders[subFolderName]; exists {
			// Folder exists in both locations, compare recursively
			subDiff := CompareFolders(oldSubFolderPath, newSubFolderPath)
			if len(subDiff.Files) > 0 || len(subDiff.SubFolders) > 0 {
				folderDiff.SubFolders = append(folderDiff.SubFolders, subDiff)
			}
		} else {
			// Folder removed
			folderDiff.SubFolders = append(folderDiff.SubFolders, FolderDiff{
				FolderName: subFolderName,
				Status:     Removed,
			})
		}
	}

	// Check for added folders
	for subFolderName, _ := range newSubFolders {
		if _, exists := oldSubFolders[subFolderName]; !exists {

			newFiles := getFileMap(newPath + "/" + subFolderName)
			var newFilesDiff []FileDiff
			for fileName, _ := range newFiles {
				newFilesDiff = append(newFilesDiff, FileDiff{
					FileName: fileName,
					Status:   Added,
				})
			}
			folderDiff.SubFolders = append(folderDiff.SubFolders, FolderDiff{
				FolderName: subFolderName,
				Status:     Added,
				Files:      newFilesDiff,
			})
		}
	}

	return folderDiff
}

// Utility functions
func getFileMap(folder string) map[string]string {
	files := make(map[string]string)

	entries, err := os.ReadDir(folder)
	if err != nil {
		return files // Return empty map on error
	}

	for _, entry := range entries {
		if !entry.IsDir() { // Only add files, not directories
			files[entry.Name()] = filepath.Join(folder, entry.Name())
		}
	}

	return files
}

func getFolderMap(folder string) map[string]string {
	subFolders := make(map[string]string)
	_ = filepath.Walk(folder, func(path string, info fs.FileInfo, err error) error {
		if err == nil && info.IsDir() && path != folder {
			subFolders[info.Name()] = path
		}
		return nil
	})
	return subFolders
}
