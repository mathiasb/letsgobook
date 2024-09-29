package main

import (
	"net/http"

	"github.com/bmizerany/pat"
)

func (app *application) routes() http.Handler {
	s := app.session
	patRouter := pat.New()
	patRouter.Get("/", s.Enable(
		noSurf(
			app.authenticate(
				http.HandlerFunc(app.home)))))
	patRouter.Get("/snippet/create", s.Enable(
		noSurf(
			app.authenticate(
				app.requireAuthentication(
					http.HandlerFunc(app.createSnippetForm))))))
	patRouter.Post("/snippet/create", s.Enable(
		noSurf(
			app.authenticate(
				app.requireAuthentication(
					http.HandlerFunc(app.createSnippet))))))
	patRouter.Get("/snippet/:id", s.Enable(
		noSurf(
			app.authenticate(
				http.HandlerFunc(app.showSnippet)))))
	patRouter.Get("/user/signup", s.Enable(
		noSurf(
			app.authenticate(
				http.HandlerFunc(app.signupUserForm)))))
	patRouter.Post("/user/signup", s.Enable(
		noSurf(
			app.authenticate(
				http.HandlerFunc(app.signupUser)))))
	patRouter.Get("/user/login", s.Enable(
		noSurf(
			app.authenticate(
				http.HandlerFunc(app.loginUserForm)))))
	patRouter.Post("/user/login", s.Enable(
		noSurf(
			app.authenticate(
				http.HandlerFunc(app.loginUser)))))
	patRouter.Post("/user/logout", s.Enable(
		noSurf(
			app.authenticate(
				app.requireAuthentication(
					http.HandlerFunc(app.logoutUser))))))

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	patRouter.Get("/static/", http.StripPrefix("/static", fileServer))

	return app.recoverPanicMiddleware(
		app.logRequestMiddleware(
			secureHeaderMiddleware(
				patRouter)))
}
