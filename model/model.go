package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderMenu struct {
	MenuId  string  `json:"menu_id"`
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

type Review struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id"`
	Score       float64            `json:"score"`
	Description string             `json:"description"`
	Timestamp   time.Time          `json:"timestamp,omitempty"`
	Vendor_id   string             `json:"vendor_id" bson:"vendor_id"`
	User_id     string             `json:"user_id"`
}

type ReviewResponseBody struct {
	Data Review `json:"data"`
}

type User struct {
	ID              primitive.ObjectID `json:"id,omitempty"`
	Name            string             `json:"name,omitempty"`
	Email           string             `json:"email"`
	Password        string             `json:"password,omitempty"`
	PasswordConfirm string             `json:"passwordConfirm,omitempty"`
	Role            string             `json:"role,omitempty"`
	CreatedAt       time.Time          `json:"created_at,omitempty"`
	UpdatedAt       time.Time          `json:"updated_at,omitempty"`
	Token           string             `json:"token,omitempty"`
}

type UserResponseBody struct {
	Data struct {
		User `json:"user"`
	} `json:"data,omitempty"`
	Status  string `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
	Token   string `json:"access_token,omitempty"`
}
