package main

import (
	"log"
)

func GetLastChar(str string) string {
	return str[len(str)-1:]
}

func LogSeparator(leng ...int) {
	var lengt int
	if len(leng) == 0 {
		lengt = 8
	} else {
		lengt = leng[0]
	}
	str := ""
	for i := 0; i <= lengt; i++ {
		str = str + "="
	}
	log.Println(str)
}
