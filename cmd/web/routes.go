package main

import (
	"net/http"

	"github.com/bmizerany/pat"
)

func (app *application) routes() http.Handler {
	patRouter := pat.New()
	patRouter.Get("/", http.HandlerFunc(app.home))
	patRouter.Get("/snippet/create", http.HandlerFunc(app.createSnippetForm))
	patRouter.Post("/snippet/create", http.HandlerFunc(app.createSnippet))
	patRouter.Get("/snippet/:id", http.HandlerFunc(app.showSnippet))

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	patRouter.Get("/static/", http.StripPrefix("/static", fileServer))

	return app.recoverPanicMiddleware(
		app.logRequestMiddleware(
			secureHeaderMiddleware(
				patRouter)))
}
