package main

import (
	"net/http"
	"time"

	"github.com/codegangsta/martini-contrib/binding"
	r "github.com/danott/recurrence"
)

type RecurrenceParams struct {
	TimeRange r.TimeRange      `json:"timeRange"`
	Schedule  r.ScheduleStruct `json:"schedule"`
}

func (rp RecurrenceParams) Validate(errors *binding.Errors, req *http.Request) {
	if rp.Schedule.Schedule == nil {
		errors.Fields["schedule"] = "Required."
	}

	if rp.TimeRange.Start.IsZero() {
		rp.TimeRange.Start = time.Now()
	}

	if rp.TimeRange.End.IsZero() {
		rp.TimeRange.End = rp.TimeRange.Start.AddDate(1, 0, 0)
	}
}
