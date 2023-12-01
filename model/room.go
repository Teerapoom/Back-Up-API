package model

import (
	"github.com/teerapoom/API_Dormitory_v.2/database"
	"gorm.io/gorm"
)

type Room struct {
	gorm.Model
	ID                 uint       `gorm:"primary_key"`
	UserID             uint       `gorm:"not null" json:"user_id"`
	Name               string     `gorm:"not null" json:"name"`
	Description        string     `json:"description"`
	SeleStatus         string     `json:"selestatus"`
	Rent               int        `json:"rent"`
	Bed_combo_mattress *bool      `json:"bed_combo_mattress"` //เตียงพร้อมฟูก
	Table              *bool      `json:"table"`              // โต๊ะ
	Wardrobe           *bool      `json:"wardrobe"`           //ตู้เสื้อผ้า
	TVShelf            *bool      `json:"tv_shelf"`           // ชั้นวางทีวี
	ShoeRack           *bool      `json:"shoe_rack"`          //ชั้นวางรองเท้า
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
func GetRoomByName(name string, room *Room) error {
	err := database.Db.Where("name = ?", name).Preload("StatusRoom").First(room).Error
	return err
}

// Update user
func UpdateRoom(Room *Room) (err error) {
	// err = database.Db.Omit("password").Updates(User).Error
	err = database.Db.Updates(Room).Preload("StatusRoom").First(&Room).Error
	if err != nil {
		return err
	}
	return nil
}

// Get user by id
func GetRoom(Room *Room, id int) (err error) {
	err = database.Db.Where("id = ?", id).Preload("StatusRoom").First(Room).Error
	if err != nil {
		return err
	}
	return nil
}
