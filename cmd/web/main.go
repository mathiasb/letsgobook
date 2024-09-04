package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type application struct {
	fail *log.Logger
	info *log.Logger
}

func main() {
	addr := flag.String("addr", ":3000", "HTTP network address")
	flag.Parse()
	info := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	fail := log.New(os.Stderr, "FAIL\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := application{
		fail: fail,
		info: info,
	}
	mux := setupRoutes(&app)

	srv := http.Server{
		Addr:     *addr,
		ErrorLog: fail,
		Handler:  mux,
	}
	info.Printf("Starting server on %s\n", srv.Addr)
	err := srv.ListenAndServe()

	fail.Fatal(err)
}

func setupRoutes(app *application) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet", app.showSnippet)
	mux.HandleFunc("/foo/", app.fooHandler)
	mux.HandleFunc("/snippet/create", app.createSnippet)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	return mux
}
