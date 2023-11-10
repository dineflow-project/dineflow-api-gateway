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
	User_id     string             `json:"user_id,omitempty"`
}

type Notification struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	RecipientID string             `json:"recipient_id" bson:"recipient_id"`
	OrderID     string             `json:"order_id" bson:"order_id"`
	IsRead      bool               `json:"is_read" bson:"is_read"`
	Type        string             `json:"type" bson:"type"`
	Timestamp   primitive.DateTime `json:"timestamp" bson:"timestamp"`
}
type UnreadNotificationResponseBody struct {
	Count int `json:"count" bson:"count"`
}

type NotificationResponseBody struct {
	Data []Notification `json:"notifications"`
}

type ReviewResponseBody struct {
	Data Review `json:"data"`
}

// type User struct {
// 	ID              primitive.ObjectID `json:"id,omitempty"`
// 	Name            string             `json:"name,omitempty"`
// 	Email           string             `json:"email"`
// 	Password        string             `json:"password,omitempty"`
// 	PasswordConfirm string             `json:"passwordConfirm,omitempty"`
// 	Role            string             `json:"role,omitempty"`
// 	CreatedAt       time.Time          `json:"created_at,omitempty"`
// 	UpdatedAt       time.Time          `json:"updated_at,omitempty"`
// 	Token           string             `json:"token,omitempty"`
// }

// type UserResponseBody struct {
// 	Data struct {
// 		User `json:"user"`
// 	} `json:"data,omitempty"`
// 	Status  string `json:"status,omitempty"`
// 	Message string `json:"message,omitempty"`
// 	Token   string `json:"access_token,omitempty"`
// }

type Menu struct {
	ID          int     `json:"id"`
	VendorID    int     `json:"vendor_id"`
	Name        string  `json:"name"`
	Price       float32 `json:"price"`
	ImagePath   string  `json:"image_path"`
	Description string  `json:"description"`
	IsAvailable int     `json:"is_available"`
}

type Status string

const (
	OPEN  Status = "Open"
	CLOSE Status = "Close"
)

type Vendor struct {
	ID               int    `json:"id"`
	CanteenID        int    `json:"canteen_id"`
	Name             string `json:"name"`
	OwnerID          string `json:"owner_id"`
	OpeningTimestamp string `json:"opening_timestamp"`
	ClosingTimestamp string `json:"closing_timestamp"`
	Status           Status `json:"status"`
	Image_path       string `json:"image_path"`
}

type Canteen struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Image_path string `json:"image_path"`
}
