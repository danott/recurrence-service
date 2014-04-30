package main

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"time"

	"github.com/danott/recurrence"
)

func timeRangeApplyDefaults(tr recurrence.TimeRange) recurrence.TimeRange {
	if tr.Start.IsZero() {
		tr.Start = time.Now()
	}

	if tr.End.IsZero() {
		tr.End = tr.Start.AddDate(1, 0, 0)
	}

	return tr
}

func timeRangeFromQueryParams(req *http.Request) recurrence.TimeRange {
	var timeRange recurrence.TimeRange
	timeRangeString := `{"start":"` + req.URL.Query().Get("start") + `","end":"` + req.URL.Query().Get("end") + `"}`
	json.Unmarshal([]byte(timeRangeString), &timeRange)
	return timeRangeApplyDefaults(timeRange)
}

func scheduleId(schedule recurrence.Schedule) string {
	raw, _ := json.Marshal(schedule)
	hash := sha1.New()
	hash.Write(raw)
	md := hash.Sum(nil)
	return hex.EncodeToString(md)
}
