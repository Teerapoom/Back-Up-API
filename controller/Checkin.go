package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/teerapoom/API_Dormitory_v.2/database"
	"github.com/teerapoom/API_Dormitory_v.2/model"
	"github.com/teerapoom/API_Dormitory_v.2/util"
)

func CreateCheckin(c *gin.Context) {
	var input model.Checkin
	var User model.User
	var user_id = util.CurrentUser(c).ID
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// user_id มาจากตาราง users
	if err := database.Db.Where("user_id = ?", input.UserNameCheckin).First(&User).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Message": "User not found"})
		return
	}
	// fmt.Println(User.ID)
	input.UserNameCheckinID = User.ID
	// Room
	ChangeStatus(c, input.RoomID)
	// Room

	if input.UserID != 0 {
		user_id = input.RoomID
	}

	checkin := model.Checkin{
		UserID:            user_id,
		UserNameCheckin:   input.UserNameCheckin, //ชื่อห้อง
		UserNameCheckinID: input.UserNameCheckinID,
		RoomID:            input.RoomID,
		Deposit:           input.Deposit,
		RentRate:          input.RentRate,
		ContractDate:      input.ContractDate,
		Fullname:          input.Fullname,
		DirthDate:         input.DirthDate,
		IssuedBy:          input.IssuedBy,
		CardNumber:        input.CardNumber,
		IssuedDate:        input.IssuedDate,
		CardCopyIMG:       input.CardCopyIMG,
		Phone1:            input.Phone1,
		Addr1:             input.Addr1,
		Place1:            input.Place1,
		Renter2:           input.Renter2,
		Birth_Date2:       input.Birth_Date2,
		IssuedBy2:         input.IssuedBy2,
		Card_number2:      input.Card_number2,
		IssuedDate2:       input.IssuedDate2,
		CardCopyIMG2:      input.CardCopyIMG2,
		Phone2:            input.Phone2,
		Addr2:             input.Addr2,
		Place2:            input.Place2,
	}

	_, err := checkin.Save()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Checkin successfuly"})
}

// get all Checkin
func GetCheckinS(c *gin.Context) {
	var checkins []model.Checkin
	err := model.GetCheckinS(&checkins)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err}) // ส่งคืน response พร้อมกับสถานะ HTTP
		return
	}
	c.JSON(http.StatusOK, checkins)
}

// เพิ่ม Fucn เมื่อ Checkin แล้วเปลี่ยน สถานะไม่ว่าง
func ChangeStatus(c *gin.Context, roomID uint) {
	var room model.Room
	if err := database.Db.First(&room, roomID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Room not found"})
		return
	}
	room.SeleStatus = "ไม่ว่าง"
	room.StatusID = 2
	if err := database.Db.Save(&room).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating room status"})
		return
	}
}
