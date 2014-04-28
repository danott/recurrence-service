package main

import (
	"github.com/codegangsta/martini"
	"github.com/danott/recurrence"
)

type simpleStore map[string]recurrence.Schedule

// DB Returns a martini.Handler
func database() martini.Handler {
	db := make(simpleStore)

	return func(c martini.Context) {
		c.Map(db)
		c.Next()
	}
}
