package model

import (
	"time"

	"github.com/teerapoom/API_Dormitory_v.2/database"
	"gorm.io/gorm"
)

type Checkin struct {
	gorm.Model
	ID                uint      `gorm:"primary_key"`
	UserID            uint      `gorm:"not null" json:"user_id"`
	RoomID            uint      `gorm:"not null" json:"room_id"`
	UserNameCheckin   string    `json:"user_name_checkin"`
	UserNameCheckinID uint      `json:"user_name_checkinid"`
	Deposit           int       `json:"deposit"`
	RentRate          int       `json:"rentrate"`
	ContractDate      time.Time `json:"contractdate"`
	Fullname          string    `json:"fullname"`
	DirthDate         time.Time `json:"dirth_date"`
	IssuedBy          string    `json:"issued_by"`
	CardNumber        string    `json:"card_number"`
	IssuedDate        time.Time `json:"issued_date"`
	CardCopyIMG       string    `json:"card_copyimg"`
	Phone1            string    `json:"phone1"`
	Addr1             string    `json:"addr1"`
	Place1            string    `json:"place1"`
	//เพิ่ม
	Renter2      string    `json:"renter2"`
	Birth_Date2  time.Time `json:"birth_date2"`
	IssuedBy2    string    `json:"issued_by2"`
	Card_number2 string    `json:"card_number2"`
	IssuedDate2  time.Time `json:"issued_date2"`
	CardCopyIMG2 string    `json:"card_copyimg2"`
	Phone2       string    `json:"phone2"`
	Addr2        string    `json:"addr2"`
	Place2       string    `json:"place2"`
	User         User      `gorm:"foreignKey:UserNameCheckinID;references:ID"`
	Room         Room      `gorm:"foreignKey:RoomID;references:ID"`
}

// `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
// add a Checkin
func (checkin *Checkin) Save() (*Checkin, error) {
	err := database.Db.Create(&checkin).Error
	if err != nil {
		return &Checkin{}, err
	}
	return checkin, nil
}

// get all Checki
func GetCheckinS(Checkin *[]Checkin) (err error) {
	// err = database.Db.Find(Checkin).Error
	err = database.Db.Preload("User").Preload("Room").Preload("Room.StatusRoom").Find(&Checkin).Error
	if err != nil {
		return err
	}
	return nil
}

// get user Checkin  ใช้ UserID ที่มาจากต้นฉบับ
func GetUserCheckin(Checkin *Checkin, uid uint) (err error) {
	err = database.Db.Where("user_id = ?", uid).First(Checkin).Error
	if err != nil {
		return err
	}
	return nil
}
