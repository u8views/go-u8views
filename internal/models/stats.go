package models

type ProfileViewsStats struct {
	DayCount   int64
	WeekCount  int64
	MonthCount int64
	TotalCount int64
}

type DayWeekMonthViewsStats struct {
	DayCount   int64
	WeekCount  int64
	MonthCount int64
}
