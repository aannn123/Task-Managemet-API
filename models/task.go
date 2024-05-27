package models

import "time"

// Struct adalah kumpulan definisi variabel (atau property) dan atau fungsi (atau method), yang dibungkus sebagai tipe data baru dengan nama tertentu atau bisa disebut class
type Task struct {
	Id           int       `gorm:"type:int;primaryKey;autoIncrement" json:"id"`
	UserId       int       `gorm:"type:int;" json:"userId"`
	Title        string    `gorm:"type:varchar(255);" json:"title"`
	Description  string    `gorm:"type:text;" json:"description"`
	Status       string    `gorm:"type:varchar(50);" json:"status"`
	Reason       string    `gorm:"type:text;default:" json:"reason"`
	Revision     int8      `gorm:"type:int;default:0" json:"revision"`
	DueDate      string    `gorm:"type:varchar(50);" json:"dueDate"`
	SubmitDate   string    `gorm:"type:varchar(50);" json:"submitDate"`
	RejectedDate string    `gorm:"type:varchar(50);" json:"rejectedDate"`
	ApprovedDate string    `gorm:"type:varchar(50);" json:"approvedDate"`
	Attachment   string    `gorm:"type:varchar(255);" json:"attachment"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	User         User      `gorm:"foreignKey:UserId" json:"user,omitempty"` // belong to
}
