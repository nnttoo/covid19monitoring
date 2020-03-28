package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"runtime"
)

type myajax struct {
	w          http.ResponseWriter
	r          *http.Request
	myscrapper *Myscraper
}

func (i *myajax) write(str string) {
	i.w.Write([]byte(str))
}
func (i *myajax) formValue(name string) string {
	return i.r.FormValue(name)
}

func openbrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}

}

func (i *myajax) start() {
	ajaxtype := i.formValue("atype")

	if ajaxtype == "getlistcountry" {
		i.w.Write(i.myscrapper.GetListCountry())
		return
	}

	if ajaxtype == "getbycountry" {
		country := i.formValue("country")
		i.w.Write(i.myscrapper.getInfobyCountry(country))
		return
	}
	if ajaxtype == "openbrowser" {
		url := i.formValue("url")
		openbrowser(url)
		i.write("")
		return
	}
}
