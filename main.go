package main

import (
	"time"

	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/binding"
	"github.com/codegangsta/martini-contrib/render"
)

func main() {
	m := martini.Classic()
	m.Use(render.Renderer())

	m.Post("/", binding.Json(RecurrenceParams{}), binding.ErrorHandler, func(ren render.Render, params RecurrenceParams) {
		var dates []time.Time

		for o := range params.Schedule.Occurrences(params.TimeRange) {
			dates = append(dates, o)
		}

		ren.JSON(200, dates)
	})
	m.Run()
}
