package diff

import (
	"bufio"
	"os"
)

// CompareFiles compares two files line by line
func CompareFiles(oldFile, newFile string) []ContentDiff {
	var diffs []ContentDiff
	oldLines := readLines(oldFile)
	newLines := readLines(newFile)

	maxLines := max(len(oldLines), len(newLines))

	for i := 0; i < maxLines; i++ {
		oldLine := getLine(oldLines, i)
		newLine := getLine(newLines, i)

		if oldLine != newLine {
			if oldLine == "" {
				diffs = append(diffs, ContentDiff{LineNumber: i + 1, NewLine: newLine, Change: Added})
			} else if newLine == "" {
				diffs = append(diffs, ContentDiff{LineNumber: i + 1, OldLine: oldLine, Change: Removed})
			} else {
				diffs = append(diffs, ContentDiff{LineNumber: i + 1, OldLine: oldLine, NewLine: newLine, Change: Modified})
			}
		}
	}

	return diffs
}

func readLines(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		return nil
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func getLine(lines []string, index int) string {
	if index >= 0 && index < len(lines) {
		return lines[index]
	}
	return ""
}
