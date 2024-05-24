package Models

import "time"

type Product struct {
	Id         int       `json:"id" gorm:"type:INT(10) UNSIGNED NOT NULL AUTO_INCREMENT; primaryKey"`
	Name       string    `json:"name"`
	Price      float32   `json:"price"`
	Image      string    `json:"image"`
	CreateAt   time.Time `json:"create_at"`
	UpdateAt   time.Time `json:"update_at"`
	CategoryId int       `json:"categoryId"`
}
