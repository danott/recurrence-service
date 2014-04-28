package main

import (
	"github.com/codegangsta/martini"
	"github.com/danott/recurrence"
)

var (
	examples = map[string]recurrence.Schedule{
		"thanksgiving": recurrence.Intersection{recurrence.Week(4), recurrence.Thursday, recurrence.November},
		"complicated": recurrence.Exclusion{
			Schedule: recurrence.Union{
				recurrence.Monday,
				recurrence.Friday,
				recurrence.OrdinalWeekday(recurrence.Last, recurrence.Wednesday),
			},
			Exclude: recurrence.Intersection{
				recurrence.Monday,
				recurrence.January,
				recurrence.Year(2015),
			},
		},
		"weekends": recurrence.Union{recurrence.Saturday, recurrence.Sunday},
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
