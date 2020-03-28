package main

import (
	"log"
	"net"
	"net/http"

	rice "github.com/GeertJohan/go.rice"
	"github.com/zserge/lorca"
)

func runServer(ln net.Listener) {
	ricebox := rice.MustFindBox("../views/")
	fs := http.FileServer(ricebox.HTTPBox())
	myhttp := http.NewServeMux()
	myhttp.Handle("/", http.StripPrefix("", fs))

	myscrapperVar := &Myscraper{}

	myhttp.HandleFunc("/ajax", func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			err := recover()
			if err != nil {
				log.Println(err)
			}
			r.Body.Close()

		}()

		myajaxVar := &myajax{
			w:          w,
			r:          r,
			myscrapper: myscrapperVar,
		}
		myajaxVar.start()

	})

	http.Serve(ln, myhttp)
}

func main() {

	ln, _ := net.Listen("tcp", "127.0.0.1:0")
	go runServer(ln)
	url := "http://" + ln.Addr().String()
	log.Println(url)

	// Create UI with basic HTML passed via data URI
	ui, err := lorca.New(url, "", 300, 420)
	if err != nil {
		log.Fatal(err)
	}
	defer ui.Close()
	// Wait until UI window is closed
	<-ui.Done()
}
