package models

import (
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Struct adalah kumpulan definisi variabel (atau property) dan atau fungsi (atau method), yang dibungkus sebagai tipe data baru dengan nama tertentu atau bisa disebut class
type User struct {
	Id        int       `gorm:"type:int;primaryKey;autoIncrement" json:"id"`
	Role      string    `gorm:"type:varchar(255);" json:"role"`
	Name      string    `gorm:"type:varchar(255);" json:"name"`
	Email     string    `gorm:"type:varchar(50);" json:"email,omitempty"`
	Password  string    `gorm:"type:varchar(255);" json:"password"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Tasks     []Task    `gorm:"constraint:OnDelete:CASCADE" json:"tasks,omitempty"` // has many;

	// omitempty = opsi yang menentukan sebuah object-key akan di hilangkan jika sebuah object-key memiliki nilai kosongopsi yang menentukan sebuah object-key akan di hilangkan jika sebuah object-key memiliki nilai kosong
}

func (u *User) AfterDelete(tx *gorm.DB) (err error) {
	tx.Clauses(clause.Returning{}).Where("user_id = ?", u.Id).Delete(&Task{})
	return
}
