package main

import (
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/binding"
	"github.com/codegangsta/martini-contrib/render"
	"github.com/danott/recurrence"
)

func main() {
	m := martini.Classic()
	m.Use(render.Renderer())
	m.Use(database())

	m.Get("/schedules", ScheduleIndex)
	m.Post("/schedules", binding.Json(recurrence.AnySchedule{}), ScheduleCreate)
	m.Get("/schedules/:sha", ScheduleShow)
	m.Delete("/schedules/:sha", ScheduleDelete)
	m.Post("/schedules/preview", binding.Json(previewParams{}), SchedulePreview)

	m.Run()
}
