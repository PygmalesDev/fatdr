package app

import (
	"fmt"
	"slices"
	"strings"
)

func printTop(dir *FatDir, counter int, longestPath int) {
	distance := strings.Repeat(" ", longestPath-len(dir.fullFath))
	fmt.Printf("%2v. %s%s  %6.2f %s\n", counter+1, dir.fullFath, distance, dir.size, dir.sizePrefix)
}

func sort(dirArr []*FatDir) {
	cmpFunc := func(dir1, dir2 *FatDir) int {
		if dir1.size > dir2.size {
			return -1
		} else if dir1.size == dir2.size {
			return 0
		} else {
			return 1
		}
	}
	slices.SortFunc(dirArr, cmpFunc)
}

// removeNestedDirs removes the directories that only contain one nested directory
func removeNestedDirs(dirArr []*FatDir) []*FatDir {
	newArr := make([]*FatDir, 0)

	newArr = append(newArr, dirArr[0])
	i := 0
	for _, dir := range dirArr {
		if newArr[i].size == dir.size {
			if len(dir.fullFath) > len(newArr[i].fullFath) {
				newArr[i] = dir
			}
			continue
		}
		newArr = append(newArr, dir)
		i++
	}
	return newArr
}

func convertSizes(dirArr []*FatDir) {
	for _, dir := range dirArr {
		dir.sizePrefix = "B"
		if dir.size >= 1_000 && dir.size < 1_000_000 {
			dir.size /= 1_000
			dir.sizePrefix = "KB"
		}
		if dir.size >= 1_000_000 && dir.size < 1_000_000_000 {
			dir.size /= 1_000_000
			dir.sizePrefix = "MB"
		}
		if dir.size >= 1_000_000_000 {
			dir.size /= 1_000_000_000
			dir.sizePrefix = "GB"
		}
		dir.fullFath = dir.path + psep + dir.name
	}
}
