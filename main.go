package main

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"time"

	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/binding"
	"github.com/codegangsta/martini-contrib/render"
	r "github.com/danott/recurrence"
)

func main() {
	m := martini.Classic()

	m.Use(render.Renderer())
	m.Use(DB())

	m.Get("/schedules", ScheduleIndex)
	m.Get("/schedules/:sha", ScheduleShow)
	m.Delete("/schedules/:sha", ScheduleDelete)
	m.Post("/schedules", binding.Json(r.ScheduleStruct{}), ScheduleCreate)
	m.Post("/schedules/preview", binding.Json(PreviewParams{}), SchedulePreview)

	m.Run()
}

type simpleStore map[string]r.Schedule

// DB Returns a martini.Handler
func DB() martini.Handler {
	db := make(simpleStore)
	sched := r.Weekday(1)
	db[ScheduleId(sched)] = sched

	return func(c martini.Context) {
		c.Map(db)
		c.Next()
	}
}

func SchedulePreview(ren render.Render, params PreviewParams) {
	var dates []time.Time

	for o := range params.Schedule.Occurrences(params.TimeRange) {
		dates = append(dates, o)
	}

	ren.JSON(200, dates)
}

func ScheduleIndex(ren render.Render, store simpleStore) {
	ren.JSON(200, store)
}

func ScheduleCreate(ren render.Render, schedule r.ScheduleStruct, store simpleStore) {
	scheduleId := ScheduleId(schedule)
	store[scheduleId] = schedule
	ren.JSON(200, scheduleId)
}

func ScheduleDelete(res http.ResponseWriter, params martini.Params, store simpleStore) {
	delete(store, params["sha"])
	res.WriteHeader(200)
}

func ScheduleShow(ren render.Render, params martini.Params, store simpleStore, req *http.Request) {
	var timeRange r.TimeRange
	timeRangeString := `{"start":"` + req.URL.Query().Get("start") + `","end":"` + req.URL.Query().Get("end") + `"}`
	json.Unmarshal([]byte(timeRangeString), &timeRange)
	timeRange = timeRangeApplyDefaults(timeRange)

	schedule, ok := store[params["sha"]]

	if !ok {
		ren.Error(404)
	} else {
		var dates []time.Time

		for o := range schedule.Occurrences(timeRange) {
			dates = append(dates, o)
		}

		ren.JSON(200, dates)
	}
}

func ScheduleId(schedule r.Schedule) string {
	raw, _ := json.Marshal(schedule)
	hash := sha1.New()
	hash.Write(raw)
	md := hash.Sum(nil)
	return hex.EncodeToString(md)
}
