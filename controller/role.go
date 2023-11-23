package controller

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/teerapoom/API_Dormitory_v.2/model"
	"gorm.io/gorm"
)

// create Role
func CreateRole(c *gin.Context) {
	var Role model.Role
	c.BindJSON(&Role)
	err := model.CreateRole(&Role)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, Role)
}

// get Roles
func GetRoles(c *gin.Context) {
	var Role []model.Role
	err := model.GetRoles(&Role)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, Role)
}

// get Role by id
func GetRole(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id")) // .strconv.Atoi แปลง srt -> int
	var Role model.Role
	err := model.GetRole(&Role, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) { //ค่าของข้อผิดพลาดที่ระบุหรือไม่
			c.AbortWithStatus(http.StatusNotFound) //404
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err}) //500
		return
	}
	c.JSON(http.StatusOK, Role)
}

// update Role
func UpdateRole(c *gin.Context) {
	var Role model.Role
	id, _ := strconv.Atoi(c.Param("id")) // ID type int
	err := model.GetRole(&Role, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.BindJSON(&Role)
	err = model.UpdateRole(&Role)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, Role)
}
