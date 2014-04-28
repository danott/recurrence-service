package main

import (
	"net/http"
	"time"

	"github.com/codegangsta/martini-contrib/binding"
	r "github.com/danott/recurrence"
)

type PreviewParams struct {
	TimeRange r.TimeRange      `json:"timeRange"`
	Schedule  r.ScheduleStruct `json:"schedule"`
}

func (rp PreviewParams) Validate(errors *binding.Errors, req *http.Request) {
	if rp.Schedule.Schedule == nil {
		errors.Fields["schedule"] = "Required."
	}

	rp.TimeRange = timeRangeApplyDefaults(rp.TimeRange)
}

func timeRangeApplyDefaults(tr r.TimeRange) r.TimeRange {
	if tr.Start.IsZero() {
		tr.Start = time.Now()
	}

	if tr.End.IsZero() {
		tr.End = tr.Start.AddDate(1, 0, 0)
	}

	return tr
}
