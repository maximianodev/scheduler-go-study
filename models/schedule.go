package models

type Schedule struct {
	ID        string `json:"id"`
	UserEmail string `json:"user_id"`
	DateHour  string `json:"date_hour"`
}

var Schedules = []Schedule{
	{ID: "1", UserEmail: "john.doe@example.com", DateHour: "2023-04-10 10:00"},
	{ID: "2", UserEmail: "john.doe@example.com", DateHour: "2023-04-10 11:00"},
	{ID: "3", UserEmail: "jane.doe@example.com", DateHour: "2023-04-20 12:00"},
}
