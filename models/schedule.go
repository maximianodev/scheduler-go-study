package models

type Schedule struct {
	ID       string `json:"id"`
	UserID   string `json:"user_id"`
	DateHour string `json:"date_hour"`
}

var Schedules = []Schedule{
	{ID: "1", UserID: "1", DateHour: "2023-04-10 10:00"},
	{ID: "2", UserID: "2", DateHour: "2023-04-10 11:00"},
	{ID: "3", UserID: "2", DateHour: "2023-04-20 12:00"},
}
