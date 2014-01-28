package main

import (
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/binding"
	"github.com/codegangsta/martini-contrib/render"
	"net/http"
	"strconv"
)

func main() {

	m := martini.Classic()
	m.Use(render.Renderer())
	m.MapTo(db, (*DB)(nil))

	m.Get("/issues", func(r render.Render) {
		r.JSON(http.StatusOK, db.GetAll())
	})

	m.Get("/issues/:id", func(r render.Render, params martini.Params) {
		id, err := strconv.Atoi(params["id"])
		issue := db.Get(id)
		if err != nil || issue == nil {
			r.JSON(http.StatusNotFound, nil)
		}
		r.JSON(http.StatusOK, issue)
	})

	m.Post("/issues", binding.Form(Issue{}), func(w, issue Issue, r render.Render, params martini.Params) {
		_, err := db.Add(&issue)
		if err != nil {
			r.JSON(http.StatusConflict, nil)
			return
		}
		r.JSON(http.StatusCreated, issue)
	})

	m.Put("/issues/:id", binding.Form(Issue{}), func(issue Issue, r render.Render, params martini.Params) {
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			r.JSON(http.StatusInternalServerError, nil)
			return
		}
		issue.Id = id
		err = db.Update(&issue)
		if err != nil {
			r.JSON(http.StatusNotFound, nil)
			return
		}
		r.JSON(http.StatusOK, issue)
	})

	m.Delete("/issues/:id", func(r render.Render, params martini.Params) {
		id, err := strconv.Atoi(params["id"])
		issue := db.Get(id)
		if err != nil || issue == nil {
			r.JSON(http.StatusNotFound, nil)
			return
		}
		db.Delete(id)
		r.JSON(http.StatusNoContent, nil)
	})

	m.Run()
}
