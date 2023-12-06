package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/teerapoom/API_Dormitory_v.2/controller"
	"github.com/teerapoom/API_Dormitory_v.2/database"
	"github.com/teerapoom/API_Dormitory_v.2/model"
	"github.com/teerapoom/API_Dormitory_v.2/util"
)

func main() {
	// load environment file
	loadEnv()

	// load database configuration and connection
	loadDatabase()

	// start the server
	serveApplication()
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	log.Println(".env file loaded successfully")
}

func loadDatabase() {
	database.InitDb()
	database.Db.AutoMigrate(&model.Role{})
	database.Db.AutoMigrate(&model.User{})
	database.Db.AutoMigrate(&model.StatusRoom{})
	database.Db.AutoMigrate(&model.Room{})
	database.Db.AutoMigrate(&model.Checkin{})
	seedData()
}

// 3 -> เจ้าของหอพัก เจ้าหน้าที่(Admin) ผู้เช่า
func seedData() {
	var roles = []model.Role{{Name: "เจ้าหน้าที่"},
		{Name: "ผู้เช่า"},
		{Name: "เจ้าของหอพัก"}}
	var status_room = []model.StatusRoom{
		{StatusName: "ว่าง"},
		{StatusName: "ไม่ว่าง"},
		{StatusName: "ปรับปรุง"}}
	var user = []model.User{
		{UserName: os.Getenv("ADMIN_FULLNAME"),
			Email:       os.Getenv("ADMIN_EMAIL"),
			Password:    os.Getenv("ADMIN_PASSWORD"),
			UserID:      os.Getenv("ADMIN_USERID"),
			Numberphone: os.Getenv("ADMIN_PHONE"),
			SeleUser:    os.Getenv("ADMIN_SELEUSER"),
			RoleID:      1},
		{UserName: os.Getenv("CEO_FULLNAME"),
			Email:       os.Getenv("CEO_EMAIL"),
			Password:    os.Getenv("CEO_PASSWORD"),
			UserID:      os.Getenv("CEO_USERID"),
			Numberphone: os.Getenv("CEO_PHONE"),
			SeleUser:    os.Getenv("CEO_SELEUSER"),
			RoleID:      3}}
	database.Db.Save(&roles) // & เข้าถึง การเข้าถึงที่อยู่ของตัวแปรในหน่วยความจำ UPDATE
	database.Db.Save(&user)
	database.Db.Save(&status_room)
}

func serveApplication() {
	router := gin.Default()

	//ทุกคนสามารถใช้งานได้
	authRoutes := router.Group("/auth/user")
	{
		// USER route
		authRoutes.POST("/login", controller.Login)
	}

	adminRoutes := router.Group("/admin")
	{
		adminRoutes.Use(util.JWTAuth()) // part นี้ต้องเป็น Admin เท่านั้น !!! ส่ง token เพื่อยืนยัน
		// User
		adminRoutes.POST("/register", controller.Register)
		adminRoutes.GET("/users", controller.GetUsers)
		adminRoutes.GET("/user/:id", controller.GetUser)
		adminRoutes.PUT("/update/user/:id", controller.UpdateUser)
		adminRoutes.DELETE("/user/remove/:id", controller.RemoveUser)
		// Role
		adminRoutes.POST("/user/role", controller.CreateRole)
		adminRoutes.GET("/user/role", controller.GetRoles)
		adminRoutes.GET("/user/role/:id", controller.GetRole)
		adminRoutes.PUT("/update/user/role/:id", controller.UpdateRole)
		//Room
		adminRoutes.POST("/add/room", controller.CreateRoom)
		adminRoutes.GET("/view/room/:id", controller.GetRoom)       //  GET ตาม ID
		adminRoutes.GET("/view/roow/:name", controller.GetRoomName) // GET ตาม Name
		adminRoutes.GET("/view/all/room", controller.GetRooms)
		adminRoutes.PUT("/update/room/:id", controller.UpdateRoom)
		//Checkin
		adminRoutes.POST("/user/checkin", controller.CreateCheckin)
		adminRoutes.GET("/user/checkin/view", controller.GetCheckinS)
	}

	router.Run(":8080")
	fmt.Println("Server running on port 8080")
}
