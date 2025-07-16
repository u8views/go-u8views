package models

type ProfileViewsStats struct {
	HourCount  int64
	DayCount   int64
	WeekCount  int64
	MonthCount int64
	TotalCount int64
}

type DayWeekMonthViewsStats struct {
	HourCount  int64
	DayCount   int64
	WeekCount  int64
	MonthCount int64
}

func (s *DayWeekMonthViewsStats) Inc() {
	s.HourCount += 1
	s.DayCount += 1
	s.WeekCount += 1
	s.MonthCount += 1
}

type TimeCount struct {
	Time  int64 `json:"time"`
	Count int64 `json:"count"`
}
