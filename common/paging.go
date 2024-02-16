package common

import (
	"fmt"
	"time"
)

type Paging struct {
	Page  int   `json:"page" form:"page"`
	Limit int   `json:"limit" form:"limit"`
	Total int64 `json:"total" form:"total"`
}

func (p *Paging) Process() {
	if p.Limit == 0 {
		p.Limit = 10
	}

	if p.Page <= 0 {
		p.Page = 1
	}
}

type PagingWithTime struct {
	Paging
	StartDate   int64     `json:"start_date" form:"start_date"`
	EndDate     int64     `json:"end_date" form:"end_date"`
	Start       time.Time `json:"-" form:"-"`
	End         time.Time `json:"-" form:"-"`
	CustomStart time.Time `json:"-" form:"-"`
	CustomEnd   time.Time `json:"-" form:"-"`
}

func (p *PagingWithTime) Process() bool {
	if p.Limit == 0 {
		p.Limit = 15
	}

	if p.Page <= 0 {
		p.Page = 1
	}
	if p.StartDate == 0 && p.EndDate == 0 {
		p.End = time.Now()
		p.Start = time.Now().AddDate(0, 0, -14)
		return true
	}
	p.Start = time.Unix(p.StartDate, 0)
	p.End = time.Unix(p.EndDate+86399, 0)
	if p.End.After(time.Now()) {
		p.End = time.Now()
	}
	if p.Start.After(p.End) {
		return false
	}
	t := time.Now()
	k := t
	fmt.Println("Is after: ", k.After(t))
	return true
}

func (p *PagingWithTime) ProcessLastTime() {
	subDays := GetSubDay(p.Start, p.End)
	p.Start = Date(p.Start)
	p.End = EndDate(p.End)
	if subDays >= 0 {
		p.CustomEnd = p.End.AddDate(0, 0, -(int(subDays) + 1))
		p.CustomStart = p.Start.AddDate(0, 0, -(int(subDays) + 1))
	}
}

func GetSubDay(start, end time.Time) int64 {
	t1 := Date(start)
	t2 := Date(end)
	days := t2.Sub(t1).Hours() / 24
	return int64(days)
}

func Date(now time.Time) time.Time {
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
}

func EndDate(now time.Time) time.Time {
	return time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, int(time.Second-time.Nanosecond), time.UTC)
}
