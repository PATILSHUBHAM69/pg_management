package models

type owner struct {
	ID          int
	First_Name  string
	Last_Name   string
	Email       string
	PhoneNumber string
	PG_Type     string
	PropertyID  uint
	City        string
	Pin_Code    int
	State       string
	Location    string
}

// type owner struct {
// 	ID          uint   `json:"id" validate:"required"`
// 	First_Name  string `json:"firstname" validate:"required"`
// 	Last_Name   string `json:"lastname" validate:"required"`
// 	Email       string `json:"email" validate:"required"`
// 	PhoneNumber string `json:"phonenumber" validate:"required"`
// 	PG_Type     string `json:"pgtype" validate:"required"`
// 	PropertyID  uint   `json:"propertyid" validate:"required"`
// 	City        string `json:"city" validate:"required"`
// 	Pin_Code    int    `json:"pincode" validate:"required"`
// 	State       string `json:"state" validate:"required"`
// 	Location    string `json:"location" validate:"required"`
// }
