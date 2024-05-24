package Models

import "time"

type Cashier struct {
	Id       uint      `json:"id"`
	Name     string    `json:"name"`
	Passcode string    `json:"passcode"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
}
