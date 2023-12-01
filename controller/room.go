package controller

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/teerapoom/API_Dormitory_v.2/database"
	"github.com/teerapoom/API_Dormitory_v.2/model"
	"github.com/teerapoom/API_Dormitory_v.2/util"
	"gorm.io/gorm"
)

// add Room
func CreateRoom(c *gin.Context) {
	var input model.Room
	StatusID := 2
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.SeleStatus == "ว่าง" {
		StatusID = 1
	} else if input.SeleStatus == "ไม่ว่าง" {
		StatusID = 2
	} else if input.SeleStatus == "ปรับปรุง" {
		StatusID = 3
	}

	database.Db.Where("Name = ?", input.Name).First(&input)
	if input.ID > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"Message": "Room Exist"})
		return
	}

	room := model.Room{
		Name:               input.Name,        // ชื่อห้องหรือเลขห้อง
		Description:        input.Description, // คำอธิบาย
		Rent:               input.Rent,        //ค่าเช่า
		Bed_combo_mattress: input.Bed_combo_mattress,
		Table:              input.Table,
		Wardrobe:           input.Wardrobe,
		TVShelf:            input.TVShelf,
		ShoeRack:           input.ShoeRack,
		SeleStatus:         input.SeleStatus,
		StatusID:           uint(StatusID),
		UserID:             util.CurrentUser(c).ID, //การเชื่อมโยง Room กับ UserID ช่วยให้ระบบสามารถติดตามว่าห้องนั้นถูกสร้างโดยผู้ใช้ใด.
	}
	savedRoom, err := room.Save()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var viewdata model.Room //ระบุตัวแปรที่จะเก็บผลลัพธ์,ระบุเงื่อนไขการค้นหา
	if err := database.Db.Preload("StatusRoom").First(&viewdata, savedRoom.ID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Room added successfuly", "room": viewdata})

}

// get Rooms
func GetRooms(c *gin.Context) {
	var Room []model.Room
	err := model.GetRooms(&Room)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, Room)
}

// Get Room BY ID
func GetRoom(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var room model.Room
	err := model.GetRoom(&room, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	c.JSON(http.StatusOK, room)
}

func UpdateRoom(c *gin.Context) {
	var room model.Room
	id, _ := strconv.Atoi(c.Param("id"))
	err := model.GetRoom(&room, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.BindJSON(&room)
	if room.SeleStatus == "ว่าง" {
		room.StatusID = 1
	} else if room.SeleStatus == "ไม่ว่าง" {
		room.StatusID = 2
	} else if room.SeleStatus == "ปรับปรุง" {
		room.StatusID = 3
	}
	err = model.UpdateRoom(&room)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, room)
}

func GetRoomName(c *gin.Context) {
	name := c.Param("name")
	var room model.Room
	err := model.GetRoomByName(name, &room)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status Code": "400", "message": "Room not found"})
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	c.JSON(http.StatusOK, room)
}
