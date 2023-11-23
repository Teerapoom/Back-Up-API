package controller

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/teerapoom/API_Dormitory_v.2/model"
	"gorm.io/gorm"
)

// create StatusRoom
func CreateStatusRoom(c *gin.Context) {
	var StatusRoom model.StatusRoom
	c.BindJSON(&StatusRoom)
	err := model.CreateStatusRoom(&StatusRoom)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, StatusRoom)
}

// get StatusRoom
func GetStatusRooms(c *gin.Context) {
	var StatusRoom []model.StatusRoom
	err := model.GetStatusRooms(&StatusRoom)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, StatusRoom)
}

// get StatusRoom by id
func GetStatusRoom(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id")) // .strconv.Atoi แปลง srt -> int
	var StatusRoom model.StatusRoom
	err := model.GetStatusRoom(&StatusRoom, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) { //ค่าของข้อผิดพลาดที่ระบุหรือไม่
			c.AbortWithStatus(http.StatusNotFound) //404
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err}) //500
		return
	}
	c.JSON(http.StatusOK, StatusRoom)
}

// update Role
func UpdateStatusRoom(c *gin.Context) {
	var StatusRoom model.StatusRoom
	id, _ := strconv.Atoi(c.Param("id")) // ID type int
	err := model.GetStatusRoom(&StatusRoom, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.BindJSON(&StatusRoom)
	err = model.UpdateStatusRoom(&StatusRoom)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, StatusRoom)
}
