package model

import (
	"gorm.io/gorm"
)

type TodoRequest struct {
	Id int `json:"id"`
	User_id string `json:"userId"`
	Status int `json:"status"`
	Title string `json:"title"`
	Memo string `json:"memo"`
	Reminder_date string `json:"reminderDate"`
	Deadline_date string `json:"deadlineDate"`
	Repeat_type int `json:"repeatType"`
	Genre int `json:"Genre"`
	Important bool `json:"important"`
	Priority int `json:"priority"`
	Image_file_path string `json:"imageFilePath"`
}

type Todo struct {
	gorm.Model
	
	Id int `gorm:"primaryKey" json:"id"`
	User_id string `json:"user_id"`
	Status int `json:"status"`
	Title string `json:"title"`
	Memo string `json:"memo"`
	Reminder_date string `json:"reminder_date"`
	Deadline_date string `json:"deadline_date"`
	Repeat_type int `json:"repeat_type"`
	Genre int `json:"Genre"`
	Important bool `json:"important"`
	Priority int `json:"priority"`
	Image_file_path string `json:"image_file_path"`
}