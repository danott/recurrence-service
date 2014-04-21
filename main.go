package main

import (
	"net/http"
	"time"

	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
	r "github.com/danott/recurrence"
)

func main() {
	m := martini.Classic()
	m.Use(render.Renderer())

	m.Post("/", func(ren render.Render, req *http.Request) {
		start := time.Now()
		end := time.Now().AddDate(2, 0, 0)
		timeRange := r.TimeRange{start, end}
		schedule := r.ScheduleFromJSON(req.Body)

		dates := make([]string, 0)
		for o := range schedule.Occurrences(timeRange) {
			dates = append(dates, o.Format("2006-01-02"))
		}

		ren.JSON(200, dates)
	})

	m.Post("/marshal-test", func(ren render.Render, req *http.Request) {
		ren.JSON(200, r.ScheduleFromJSON(req.Body))
	})

	m.Get("/example", func(ren render.Render) {
		s := r.Difference{
			r.Intersection{
				r.Friday,
				r.Union{
					r.June,
					r.July,
					r.August,
				},
			},
			r.Year(1000),
		}

		ren.JSON(200, s)
	})

	m.Run()
}
