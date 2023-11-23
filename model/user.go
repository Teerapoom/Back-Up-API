package model

import (
	"html"
	"strings"

	"github.com/teerapoom/API_Dormitory_v.2/database"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User model
type User struct {
	gorm.Model
	ID          uint   `gorm:"primary_key"`
	RoleID      uint   `gorm:"not null;" json:"role_id"` //สิทธิผู้ใช้ DEFAULT:3
	UserID      string `json:"userID"`                   // A2342
	SeleUser    string `json:"seleuser"`                 // sele ข้อมูล
	UserName    string `gorm:"size:255;not null;unique" json:"username"`
	Email       string `gorm:"size:255;not null;unique" json:"email"`
	Password    string `gorm:"size:255;not null" json:"password"` // temp json:"-"
	Numberphone string `json:"numberphone"`
	Role        Role   `gorm:"constraint:OfnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
}

// Save user details
func (user *User) Save() (*User, error) {
	err := database.Db.Create(&user).Error
	if err != nil {
		return &User{}, err
	}
	return user, nil
}

// Generate encrypted password
func (user *User) BeforeSave(*gorm.DB) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost) //เข้ารหัส
	if err != nil {
		return err
	}
	user.Password = string(passwordHash)                          // ใส่ค่าเข้าไป
	user.Email = html.EscapeString(strings.TrimSpace(user.Email)) //อาจเป็น Fllname ปกกันการโจมตีเเบบ XSS
	return nil
}

// Get all users
func GetUsers(User *[]User) (err error) {
	err = database.Db.Find(User).Error
	if err != nil {
		return err
	}
	return nil
}

// Get user by username
func GetUserByUsername(username string) (User, error) {
	var user User
	err := database.Db.Where("user_id =?", username).Find(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}

// Validate user password
func (user *User) ValidateUserPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}

// Get user by id
func GetUserById(id uint) (User, error) {
	var user User
	err := database.Db.Where("id=?", id).Find(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}

// Get user by id
func GetUser(User *User, id int) (err error) {
	err = database.Db.Where("id = ?", id).First(User).Error
	if err != nil {
		return err
	}
	return nil
}

// Update user
func UpdateUser(User *User) (err error) {
	// err = database.Db.Omit("password").Updates(User).Error
	err = database.Db.Updates(User).Error //.Omit ที่นี้คือยกเว้น อัพเดททุกตัว ยกเว้น password
	if err != nil {
		return err
	}
	return nil
}
