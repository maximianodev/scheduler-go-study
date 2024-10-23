package clients

import (
	"github.com/maximianodev/scheduler/models"
)

func ListAllUsers() *[]models.User {
	return &models.Users
}

func FindUserById(id string) *models.User {
	for _, user := range models.Users {
		if user.ID == id {
			return &user
		}
	}

	return nil
}

func ListAllSchedules() *[]models.Schedule {
	return &models.Schedules
}

func ListScheduleByEmail(email string) *[]models.Schedule {
	var schedules []models.Schedule

	for _, schedule := range models.Schedules {
		if schedule.UserEmail == email {
			schedules = append(schedules, schedule)
		}
	}

	return &schedules
}
