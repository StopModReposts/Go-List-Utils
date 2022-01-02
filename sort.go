package main

import (
	"log"
	"sort"
	"strings"
)

func SortCommand() {
	f := GetListFile()
	t := GetList(f)
	sort.SliceStable(t, func(i, j int) bool {
		compared := strings.Compare(t[i].Domain, t[j].Domain)
		if compared == -1 {
			return true
		}
		if compared == 1 {
			return false
		}
		return false
	})
	SaveList(f, t)
	log.Println("Sorting done")
}
