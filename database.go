package main

import (
	"github.com/codegangsta/martini"
	"github.com/danott/recurrence"
)

var (
	examples = map[string]recurrence.Schedule{
		"thanksgiving": recurrence.Intersection{recurrence.Week(4), recurrence.Thursday, recurrence.November},
		"weekends":     recurrence.Union{recurrence.Saturday, recurrence.Sunday},
	}
)

type simpleStore map[string]recurrence.Schedule

func database() martini.Handler {
	db := make(simpleStore)

	return func(c martini.Context) {
		c.Map(db)
		c.Next()
	}
}
