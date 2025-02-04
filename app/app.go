package app

import (
	"fmt"
	"os"
	"strings"
)

const psep = string(os.PathSeparator)

func CollectFat() {
	rootPath, topLen := getArgs()
	rootDir := newFatDir("", "")
	dirArr := make([]*FatDir, 0)

	countRecursive(rootPath, "", rootDir, &dirArr, 0)
	sort(dirArr)
	convertSizes(dirArr)
	dirArr = removeNestedDirs(dirArr)

	arrSize := min(len(dirArr), topLen)
	var longestPath int
	for _, dir := range dirArr[:arrSize] {
		longestPath = max(longestPath, len(dir.fullFath))
	}
	for i, dir := range dirArr[:arrSize] {
		printTop(dir, i, longestPath)
	}
}

func countRecursive(path, name string, parent *FatDir, dirArr *[]*FatDir, addDash int) {
	var (
		entries  []os.DirEntry
		err      error
		fileInfo os.FileInfo

		fatDir    = newFatDir(path, name)
		entryPath = path + strings.Repeat(psep, addDash) + name
	)

	if entries, err = os.ReadDir(entryPath); err != nil {
		fmt.Printf("Failed to open directory entry under '%s', skipping...\n", entryPath)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			countRecursive(entryPath, entry.Name(), fatDir, dirArr, 1)
		} else {
			if fileInfo, err = os.Stat(entryPath + psep + entry.Name()); err != nil {
				fmt.Printf("Failed to read information from file '%s', skipping...\n", entry.Name())
				continue
			}
			fatDir.size += float32(fileInfo.Size())
		}
	}
	*dirArr = append(*dirArr, fatDir)
	parent.size += fatDir.size
}
