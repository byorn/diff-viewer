package main

import (
	"diffviewer/diff"
	"fmt"
)

func main() {
	fd := diff.CompareFolders("sample1", "sample2")
	fmt.Println(fd.String())
}
