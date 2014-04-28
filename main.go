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
	m.Get("/schedules/bootstrap", ScheduleBootstrap)
	m.Get("/schedules/:sha", ScheduleShow)
	m.Post("/schedules", binding.Json(recurrence.AnySchedule{}), ScheduleCreate)
	m.Post("/schedules/preview", binding.Json(recurrence.AnySchedule{}), SchedulePreview)
	m.Delete("/schedules/:sha", ScheduleDelete)

	m.Run()
}
