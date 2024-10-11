package clients

import (
	"github.com/maximianodev/scheduler/models"
)

var schedules []models.Schedule = []models.Schedule{
	{ID: "1", UserID: "1", DateHour: "2023-04-10 10:00"},
	{ID: "2", UserID: "2", DateHour: "2023-04-10 11:00"},
	{ID: "3", UserID: "2", DateHour: "2023-04-20 12:00"},
}

var users []models.User = []models.User{
	{ID: "1", Name: "John Doe", Email: "john.doe@example.com", Phone: "123-456-7890"},
	{ID: "2", Name: "Jane Doe", Email: "jane.doe@example.com", Phone: "987-654-3210"},
}

func GetAllSchedules() *[]models.Schedule {
	return &schedules
}

func GetAllUsers() *[]models.User {
	return &users
}

func GetUserById(id string) *models.User {
	for _, user := range users {
		if user.ID == id {
			return &user
		}
	}

	return nil
}
