package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
	"github.com/mgutz/ansi"
)

var (
	httpaddr = flag.String("http", ":8099", "http server address")
)

type App struct {
	Storage
	API http.Handler
}

func NewApp() *App {
	var a = new(App)
	a.Storage = NewStorage("sqlite3", "./checkepisode.db")
	a.API = NewAPI(a)
	return a
}

func main() {
	flag.Parse()

	var app = NewApp()

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, ansi.Color(fmt.Sprintf("%s", r.Method), "black+h"), r.URL)
		if r.URL.String() == "/" {
			http.ServeFile(w, r, "./static/index.html")
			return
		} else if r.URL.String()[0:5] == "/api/" {
			app.API.ServeHTTP(w, r)
			return
		}
		http.NotFound(w, r)
	})

	log.Println("serving at", *httpaddr)
	if err := http.ListenAndServe(*httpaddr, nil); err != nil {
		log.Fatal(err)
	}
}
