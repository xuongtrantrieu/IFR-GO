package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"apps"
	"apps/user"
	"db"
	"utils/response"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

var routes = flag.Bool("routes", false, "Generate router documentation")

func main() {
	flag.Parse()

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("root.")) })
	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("pong")) })
	r.Get("/panic", func(w http.ResponseWriter, r *http.Request) { panic("test") })

	r.Mount("/user", userRouter())

	log.Fatal(http.ListenAndServe(":8000", r))
}

func userRouter() chi.Router {
	router := chi.NewRouter()
	// r.Use(adminOnly)
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {

	})
	router.Post("/", func(w http.ResponseWriter, r *http.Request) {
		user := &user.User{}
		err := render.Bind(r, user)
		if err != nil {
			response.ResponseERR(w, r, err.Error())
		} else {
			if err := apps.Save(db.InsertIntoCollection, user); err != nil {
				response.ResponseERR(w, r, err.Error())
			} else {
				response.ResponseOK(w, r, user)
			}
		}
	})

	return router
}

func adminOnly(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		isAdmin, ok := r.Context().Value("acl.admin").(bool)
		fmt.Println(r.Context().Value("user"))
		if !ok || !isAdmin {
			http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}
