package model

type Register struct {
	FullName    string `json:"fullname" `
	Email       string `json:"email" binding:"required"`
	Password    string `json:"password" binding:"required"`
	Numberphone string `json:"Numberphone"`
	UserID      string `json:"userID"`
	SeleUser    string `json:"seleuser"`
}

// table user{
// 	user_id int[PK] มี คือ ID
// ----	password varchar(20) //มี
// ----	fullname varchar(30) //มี
// ----	email varchar(100) // มี
// ----	phone varchar(10) // มี
// 	RoleID uint //user
// 	RentalUSER if เขียน if ให้ gen RoleID
// 	Role Role
//   }
