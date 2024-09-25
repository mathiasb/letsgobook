package main

import (
	"crypto/tls"
	"database/sql"
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golangcollege/sessions"
	"github.com/mathiasb/snippetbox/pkg/models/mysql"
)

type application struct {
	fail          *log.Logger
	info          *log.Logger
	session       *sessions.Session
	snippets      *mysql.SnippetModel
	templateCache map[string]*template.Template
}

func main() {
	addr := flag.String("addr", ":3000", "HTTP network address")
	dsn := flag.String("dsn", "web:pass@/snippetbox?parseTime=true", "MySQL Data Source Name")
	secret := flag.String("secret", "s6Ndh+pPbnzHbS*+9Pk8qGWhTzbpa@ge", "Secret key")
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

	session := sessions.New([]byte(*secret))
	session.Lifetime = 12 * time.Hour
	session.Secure = true

	app := application{
		fail:          fail,
		info:          info,
		session:       session,
		snippets:      &mysql.SnippetModel{DB: db},
		templateCache: templateCache,
	}

	tlsConfig := &tls.Config{
		PreferServerCipherSuites: true,
		CurvePreferences:         []tls.CurveID{tls.X25519, tls.CurveP256},
	}

	srv := http.Server{
		Addr:     *addr,
		ErrorLog: fail,
		Handler:  app.routes(),
		TLSConfig: tlsConfig,
		IdleTimeout: time.Minute,
		ReadTimeout: 5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	info.Printf("Starting server on %s\n", srv.Addr)
	err = srv.ListenAndServeTLS("./tls/cert.pem", "./tls/key.pem")
	info.Println("Stopped server")
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
