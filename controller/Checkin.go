package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/teerapoom/API_Dormitory_v.2/model"
	"github.com/teerapoom/API_Dormitory_v.2/util"
)

func CreateCheckin(c *gin.Context) {
	var input model.Checkin
	var user_id = util.CurrentUser(c).ID
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.UserID != 0 {
		user_id = input.RoomID
	}

	checkin := model.Checkin{
		UserID:       user_id,
		RoomID:       input.RoomID,
		Deposit:      input.Deposit,
		RentRate:     input.RentRate,
		ContractDate: input.ContractDate,
		Fullname:     input.Fullname,
		DirthDate:    input.DirthDate,
		IssuedBy:     input.IssuedBy,
		CardNumber:   input.CardNumber,
		IssuedDate:   input.IssuedDate,
		CardCopyIMG:  input.CardCopyIMG,
		Phone1:       input.Phone1,
		Addr1:        input.Addr1,
		Place1:       input.Place1,
	}
	savedcheckin, err := checkin.Save()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Booking done successfuly", "Checkin": savedcheckin})
}

// get all Checkin
func GetCheckinS(c *gin.Context) {
	var checkin []model.Checkin
	err := model.GetCheckinS(&checkin)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err}) // ส่งคืน response พร้อมกับสถานะ HTTP
		return
	}
	c.JSON(http.StatusOK, checkin)
}
