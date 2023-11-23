package model

import (
	"github.com/teerapoom/API_Dormitory_v.2/database"
	"gorm.io/gorm"
)

type StatusRoom struct {
	gorm.Model
	ID         uint   `gorm:"primary_key"`
	StatusName string `gorm:"size:50;not null;unique" json:"statusname"`
}

// Create a StatusRoom
func CreateStatusRoom(StatusRoom *StatusRoom) (err error) {
	err = database.Db.Create(StatusRoom).Error
	if err != nil {
		return err
	}
	return nil
}

// Get all StatusRoom
func GetStatusRooms(StatusRoom *[]StatusRoom) (err error) {
	err = database.Db.Find(StatusRoom).Error
	if err != nil {
		return err
	}
	return nil
}

// Get GetStatusRoom by id
func GetStatusRoom(StatusRoom *StatusRoom, id int) (err error) {
	err = database.Db.Where("id = ?", id).First(StatusRoom).Error
	if err != nil {
		return err
	}
	return err
}

func UpdateStatusRoom(StatusRoom *StatusRoom) (err error) {
	database.Db.Save(StatusRoom)
	return nil
}
