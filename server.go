package main

import (
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/binding"
	"github.com/codegangsta/martini-contrib/render"
	"strconv"
)

type Issue struct {
	Title string `json:"title" form:"title"`
	Body  string `json:"body" form:"body"`
}

var (
	data []Issue
)

func main() {
	data = append(data, Issue{Title: "test", Body: "Test"})
	data = append(data, Issue{Title: "test1", Body: "Test2"})

	m := martini.Classic()
	m.Use(render.Renderer())

	m.Get("/issues", func(r render.Render) {
		r.JSON(200, data)
	})

	m.Get("/issues/:id", func(r render.Render, params martini.Params) {
		if params["id"] == "" {
			r.JSON(404, nil)
		}
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			r.JSON(500, nil)
		}
		r.JSON(200, data[id])
	})

	m.Post("/issues", binding.Form(Issue{}), func(issue Issue, r render.Render, params martini.Params) {
		data = append(data, issue)
		r.JSON(200, issue)
	})

	m.Put("/issues/:id", binding.Form(Issue{}), func(issue Issue, r render.Render, params martini.Params) {
		if params["id"] == "" {
			r.JSON(404, nil)
		}
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			r.JSON(500, nil)
		}
		i := &data[id]
		i.Title = issue.Title
		i.Body = issue.Body
		r.JSON(200, i)
	})

	m.Delete("/issues/:id", func(r render.Render, params martini.Params) {
		if params["id"] == "" {
			r.JSON(404, nil)
		}
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			r.JSON(500, nil)
		}
		data = append(data[:id], data[id+1:]...)
		r.JSON(200, nil)
	})

	m.Run()
}
