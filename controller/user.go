package controller

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/teerapoom/API_Dormitory_v.2/model"
	"github.com/teerapoom/API_Dormitory_v.2/util"

	"gorm.io/gorm"
)

// Register user
func Register(context *gin.Context) {
	var db *gorm.DB
	var input model.Register
	roleID := 2
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.SeleUser == "เจ้าของหอพัก" {
		input.UserID = GenerateIDUser_forlandlord(db)
		roleID = 3
	} else if input.SeleUser == "เจ้าหน้าที่" {
		input.UserID = GenerateIDUser_forAdmin(db)
		roleID = 1
	} else if input.SeleUser == "ผู้เช่า" {
		input.UserID = GenerateIDUser_forTenant(db)
		roleID = 2
	}

	//ทางซ้ายคือเก็บใน Database ขวาเป็นการส่ง register
	user := model.User{
		UserID:      input.UserID,   // สุ่ม L A T
		SeleUser:    input.SeleUser, // sele ด็อปดาว
		UserName:    input.FullName,
		Email:       input.Email,
		Password:    input.Password,
		Numberphone: input.Numberphone,
		RoleID:      uint(roleID), //สมัครแล้วเป็น ลูกค้าเลย
	}

	savedUser, err := user.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"user": savedUser})

}

// User Login
func Login(context *gin.Context) {
	var input model.Login

	if err := context.ShouldBindJSON(&input); err != nil {
		var errorMessage string
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			validationError := validationErrors[0]
			if validationError.Tag() == "required" {
				errorMessage = fmt.Sprintf("%s not provided", validationError.Field())
			}
		}
		context.JSON(http.StatusBadRequest, gin.H{"error": errorMessage})
		return
	}

	user, err := model.GetUserByUsername(input.UserName)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = user.ValidateUserPassword(input.Password)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	jwt, err := util.GenerateJWT(user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"token": jwt, "username": input.UserName, "message": "Successfully logged in"})

}

// get all users
func GetUsers(context *gin.Context) {
	var user []model.User
	err := model.GetUsers(&user)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	context.JSON(http.StatusOK, user)
}

// get user by id
func GetUser(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))
	var user model.User
	err := model.GetUser(&user, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			context.AbortWithStatus(http.StatusNotFound)
			return
		}

		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	context.JSON(http.StatusOK, user)
}

// update user
func UpdateUser(c *gin.Context) {
	//var input model.Update
	var db *gorm.DB
	var User model.User
	id, _ := strconv.Atoi(c.Param("id"))

	err := model.GetUser(&User, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.BindJSON(&User)

	if User.SeleUser == "เจ้าของหอพัก" {
		User.RoleID = 3
		if !strings.HasPrefix(User.UserID, "L") {
			User.UserID = GenerateIDUser_forlandlord(db)
		}
	} else if User.SeleUser == "เจ้าหน้าที่" {
		User.RoleID = 1
		if !strings.HasPrefix(User.UserID, "A") {
			User.UserID = GenerateIDUser_forAdmin(db)
		}
	} else if User.SeleUser == "ผู้เช่า" {
		User.RoleID = 2
		if !strings.HasPrefix(User.UserID, "T") {
			User.UserID = GenerateIDUser_forTenant(db)
		}

		// User.UserID = GenerateIDUser_forTenant(db)
	}

	err = model.UpdateUser(&User)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, User)
}

func RemoveUser(c *gin.Context) {
	var User model.User
	id, _ := strconv.Atoi(c.Param("id"))

	err := model.GetUser(&User, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	err = model.Remove(&User)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "Successfully Remove"})
}
