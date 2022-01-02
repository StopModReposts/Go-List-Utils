package main

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

func GetList(path string) []IllegalSite {
	file1, err1 := ioutil.ReadFile(path)
	HandleError(err1)

	t := []IllegalSite{}
	err2 := yaml.Unmarshal(file1, &t)
	HandleError(err2)
	return t
}

func SaveList(path string, data []IllegalSite) {
	d, err := yaml.Marshal(data)
	HandleError(err)

	err2 := ioutil.WriteFile(path, d, 0)
	HandleError(err2)

	log.Println("printed to", path)
}
