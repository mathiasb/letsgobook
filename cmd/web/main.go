package main

import (
	"database/sql"
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/mathiasb/snippetbox/pkg/models/mysql"
)

type application struct {
	fail          *log.Logger
	info          *log.Logger
	snippets      *mysql.SnippetModel
	templateCache map[string]*template.Template
}

func main() {
	addr := flag.String("addr", ":3000", "HTTP network address")
	dsn := flag.String("dsn", "web:pass@/snippetbox?parseTime=true", "MySQL Data Source Name")
	flag.Parse()

	info := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	fail := log.New(os.Stderr, "FAIL\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDb(*dsn)
	if err != nil {
		fail.Fatal(err)
	}
	defer db.Close()

	templateCache, err := newTemplateCache("./ui/html/")
	if err != nil {
		fail.Fatal(err)
	}

	// application dependencies
	app := application{
		fail:          fail,
		info:          info,
		snippets:      &mysql.SnippetModel{DB: db},
		templateCache: templateCache,
	}

	srv := http.Server{
		Addr:     *addr,
		ErrorLog: fail,
		Handler:  app.routes(),
	}

	info.Printf("Starting server on %s\n", srv.Addr)
	err = srv.ListenAndServe()

	fail.Fatal(err)
}

func openDb(dsn string) (*sql.DB, error) {
	if db, err := sql.Open("mysql", dsn); err != nil {
		return nil, err
	} else if err = db.Ping(); err != nil {
		return nil, err
	} else {
		return db, nil
	}
}
