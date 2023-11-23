package model

import (
	"github.com/teerapoom/API_Dormitory_v.2/database"
	"gorm.io/gorm"
)

// Role model
type Role struct {
	gorm.Model
	ID   uint   `gorm:"primary_key"`
	Name string `gorm:"size:50;not null;unique" json:"name"`
}

// Create a role
func CreateRole(Role *Role) (err error) {
	err = database.Db.Create(Role).Error // Create เพิ่มข้อมูลใหม่
	if err != nil {
		return err
	}
	return nil
}

// Get all roles
func GetRoles(Role *[]Role) (err error) {
	err = database.Db.Find(Role).Error
	if err != nil {
		return err
	}
	return nil
}

// Get role by id
func GetRole(Role *Role, id int) (err error) {
	err = database.Db.Where("id = ?", id).First(Role).Error
	if err != nil {
		return err
	}
	return nil
}

// Update role
func UpdateRole(Role *Role) (err error) {
	database.Db.Save(Role) //อัปเดตหรือบันทึกข้อมูล
	return nil
}
