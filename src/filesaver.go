package main

import (
	"io/ioutil"
	"log"
	"os"
)

var pathsavedcountri = "./lastcountry"

func getSelectedCountry() string {
	dat, err := ioutil.ReadFile(pathsavedcountri)
	if err != nil {
		return "Indonesia"
	}

	return string(dat)
}
func saveSelectedCountr(country string) {
	defer func() {
		err := recover()
		if err != nil {
			log.Println(err)
		}
	}()

	ioutil.WriteFile(pathsavedcountri, []byte(country), os.ModePerm)
}
