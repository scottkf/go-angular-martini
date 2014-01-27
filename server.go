package main

import (
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/binding"
	"github.com/codegangsta/martini-contrib/render"
	"net/http"
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
		r.JSON(http.StatusOK, data)
	})

	m.Get("/issues/:id", func(r render.Render, params martini.Params) {
		id, err := Validate(params["id"])
		if err > 0 {
			r.JSON(err, nil)
			return
		}
		r.JSON(http.StatusOK, data[id])
	})

	m.Post("/issues", binding.Form(Issue{}), func(issue Issue, r render.Render, params martini.Params) {
		data = append(data, issue)
		r.JSON(http.StatusOK, issue)
	})

	m.Put("/issues/:id", binding.Form(Issue{}), func(issue Issue, r render.Render, params martini.Params) {
		id, err := Validate(params["id"])
		if err > 0 {
			r.JSON(err, nil)
			return
		}
		i := &data[id]
		i.Title = issue.Title
		i.Body = issue.Body
		r.JSON(http.StatusOK, i)
	})

	m.Delete("/issues/:id", func(r render.Render, params martini.Params) {
		id, err := Validate(params["id"])
		if err > 0 {
			r.JSON(err, nil)
			return
		}
		data = append(data[:id], data[id+1:]...)
		r.JSON(http.StatusOK, nil)
	})

	m.Run()
}

func Validate(id string) (int, int) {
	if id == "" {
		return -1, http.StatusNotFound
	}
	i, err := strconv.Atoi(id)
	if err != nil {
		return -1, http.StatusInternalServerError
	}
	if i >= len(data) {
		return -1, http.StatusBadRequest
	}
	return i, -1
}
