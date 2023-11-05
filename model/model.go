package model

import (
	"time"
)

type OrderMenu struct {
	MenuId  string  `json:"id"`
	Amount  int     `json:"amount"`
	Price   float32 `json:"price"`
	Request string  `json:"request,omitempty"`
}

type Order struct {
	ID         string       `json:"id,omitempty" bson:"_id,omitempty"`
	Status     string       `json:"status,omitempty"`
	OrderMenus []*OrderMenu `json:"order_menus,omitempty"`
	VendorId   string       `json:"vendor_id,omitempty"`
	Price      float32      `json:"price,omitempty"`
	UserId     string       `json:"user_id,omitempty"`
	CreateAt   time.Time    `json:"created_at,omitempty"`
	UpdatedAt  time.Time    `json:"updated_at,omitempty"`
}
