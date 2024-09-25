package main

import (
	"net/http"

	"github.com/bmizerany/pat"
)

func (app *application) routes() http.Handler {
	s := app.session
	patRouter := pat.New()
	patRouter.Get("/", s.Enable(http.HandlerFunc(app.home)))
	patRouter.Get("/snippet/create", s.Enable(http.HandlerFunc(app.createSnippetForm)))
	patRouter.Post("/snippet/create", s.Enable(http.HandlerFunc(app.createSnippet)))
	patRouter.Get("/snippet/:id", s.Enable(http.HandlerFunc(app.showSnippet)))

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	patRouter.Get("/static/", http.StripPrefix("/static", fileServer))

	return app.recoverPanicMiddleware(
		app.logRequestMiddleware(
			secureHeaderMiddleware(
				patRouter)))
}
