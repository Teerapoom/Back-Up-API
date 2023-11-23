package model

import (
	"github.com/teerapoom/API_Dormitory_v.2/database"
	"gorm.io/gorm"
)

type Room struct {
	gorm.Model
	ID                 uint       `gorm:"primary_key"`
	UserID             uint       `gorm:"not null" json:"user_id"`
	Name               string     `gorm:"not null;unique" json:"name"`
	Description        string     `json:"description"`
	SeleStatus         string     `json:"selestatus"`
	Rent               int        `json:"rent"`
	Bed_combo_mattress bool       `json:"bed_combo_mattress"` //เตียงพร้อมฟูก
	Table              bool       `json:"table"`              // โต๊ะ
	Wardrobe           bool       `json:"wardrobe"`           //ตู้เสื้อผ้า
	TVShelf            bool       `json:"tv_shelf"`           // ชั้นวางทีวี
	ShoeRack           bool       `json:"shoe_rack"`          //ชั้นวางรองเท้า
	StatusID           uint       `gorm:"not null;DEFAULT:2;" json:"statusID"`
	User               User       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	StatusRoom         StatusRoom `gorm:"foreignKey:StatusID;references:ID"`
}

// Save Room details
func (room *Room) Save() (*Room, error) {
	if err := database.Db.Create(&room).Error; err != nil {
		return nil, err
	}
	return room, nil
}

// Get all Rooms
func GetRooms(Room *[]Room) (err error) {
	// err = database.Db.Find(Room).Error
	err = database.Db.Preload("StatusRoom").Find(&Room).Error
	if err != nil {
		return err
	}
	return nil
}

// Get Room by name name คือ เลขห้อง
func GetRoomByName(name string) (Room, error) {
	var room Room
	err := database.Db.Where("Name = ? ", name).Find(&room).Error
	if err != nil {
		return Room{}, err
	}
	return room, nil
}

// Get user by id
func GetRoomById(id uint) (Room, error) {
	var room Room
	err := database.Db.Where("id=?", id).Find(&room).Error
	if err != nil {
		return Room{}, err
	}
	return room, nil
}

// Update user
func UpdateRoom(Room *Room) (err error) {
	// err = database.Db.Omit("password").Updates(User).Error
	err = database.Db.Updates(Room).Error //.Omit ที่นี้คือยกเว้น อัพเดททุกตัว ยกเว้น password
	if err != nil {
		return err
	}
	return nil
}
