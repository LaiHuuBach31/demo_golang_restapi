package Models

import "time"

type Category struct {
	Id       int       `json:"id" gorm:"type:INT(10) UNSIGNED NOT NULL AUTO_INCREMENT; primaryKey"`
	Name     string    `json:"name"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
}
