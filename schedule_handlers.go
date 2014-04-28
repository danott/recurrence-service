package main

import (
	"net/http"
	"time"

	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
	"github.com/danott/recurrence"
)

func ScheduleIndex(ren render.Render, store simpleStore) {
	ren.JSON(200, store)
}

func ScheduleCreate(ren render.Render, schedule recurrence.AnySchedule, store simpleStore) {
	id := scheduleId(schedule)
	store[id] = schedule
	ren.JSON(200, id)
}

func ScheduleShow(ren render.Render, params martini.Params, store simpleStore, req *http.Request) {
	timeRange := timeRangeFromQueryParams(req)
	schedule, ok := store[params["sha"]]

	if !ok {
		ren.Error(404)
		return
	}

	var dates []time.Time

	for o := range schedule.Occurrences(timeRange) {
		dates = append(dates, o)
	}

	ren.JSON(200, dates)
}

func ScheduleDelete(res http.ResponseWriter, params martini.Params, store simpleStore) {
	delete(store, params["sha"])
	res.WriteHeader(200)
}

func SchedulePreview(ren render.Render, params previewParams) {
	params.TimeRange = timeRangeApplyDefaults(params.TimeRange)
	var dates []time.Time

	for o := range params.Schedule.Occurrences(params.TimeRange) {
		dates = append(dates, o)
	}

	ren.JSON(200, dates)
}