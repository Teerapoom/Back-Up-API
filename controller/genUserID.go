package controller

import (
	"fmt"
	"sync"

	"github.com/teerapoom/API_Dormitory_v.2/database"
	"github.com/teerapoom/API_Dormitory_v.2/model"
	"gorm.io/gorm"
)

// Check if the given user ID exists in the database
func isUserIDExists(db *gorm.DB, userID string) bool {
	var count int64
	database.Db.Model(&model.User{}).Where("user_id = ?", userID).Count(&count)
	return count > 0
}

// Gen UserID L A T
// เจ้าหน้าที่
func GenerateIDUser_forAdmin(db *gorm.DB) string {
	var id string
	for {
		// Generate the ID
		id = "A" + randomNumbers_forAdmin()

		// Check if it exists in the database
		if !isUserIDExists(db, id) {
			break // Unique ID found
		}
		// If duplicate, loop will continue and generate a new ID
	}
	return id
}

var (
	currentID_forAdmin int = 0
	lock_forAdmin      sync.Mutex
)

func randomNumbers_forAdmin() string {
	lock_forAdmin.Lock()
	defer lock_forAdmin.Unlock()

	currentID_forAdmin++
	if currentID_forAdmin > 9999 {
		currentID_forAdmin = 1 // รีเซ็ตกลับไปที่ 1 เมื่อถึง 9999
	}
	return fmt.Sprintf("%04d", currentID_forAdmin)
}

// เจ้าหน้าที่
// ----------------- ผู้เช่า -----------------
func GenerateIDUser_forTenant(db *gorm.DB) string {
	var id string
	for {
		// Generate the ID
		id = "T" + randomNumbers_forTenant()

		// Check if it exists in the database
		if !isUserIDExists(db, id) {
			break // Unique ID found
		}
		// If duplicate, loop will continue and generate a new ID
	}
	return id
}

var (
	currentID_forTenant int = 0
	lock_forTenant      sync.Mutex
)

func randomNumbers_forTenant() string {
	lock_forTenant.Lock()
	defer lock_forTenant.Unlock()

	currentID_forTenant++
	if currentID_forTenant > 9999 {
		currentID_forTenant = 1 // รีเซ็ตกลับไปที่ 1 เมื่อถึง 9999
	}
	return fmt.Sprintf("%04d", currentID_forTenant)
}

// ----------------- ผู้เช่า -----------------

// ----------------- เจ้าของ -----------------
// L เช่าของหอ landlord
func GenerateIDUser_forlandlord(db *gorm.DB) string {
	var id string
	for {
		// Generate the ID
		id = "L" + randomNumbers_forlandlord()

		// Check if it exists in the database
		if !isUserIDExists(db, id) {
			break // Unique ID found
		}
		// If duplicate, loop will continue and generate a new ID
	}
	return id
}

var (
	currentID_forlandlord int = 0
	lock_forlandlord      sync.Mutex
)

func randomNumbers_forlandlord() string {
	lock_forlandlord.Lock()
	defer lock_forlandlord.Unlock()

	currentID_forlandlord++
	if currentID_forlandlord > 9999 {
		currentID_forlandlord = 1 // รีเซ็ตกลับไปที่ 1 เมื่อถึง 9999
	}
	return fmt.Sprintf("%04d", currentID_forlandlord)
}

// ----------------- เจ้าของ -----------------
