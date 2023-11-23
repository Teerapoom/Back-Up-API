package util

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/teerapoom/API_Dormitory_v.2/model"
)

// retrieve JWT key from .env file
var privateKey = []byte(os.Getenv("JWT_PRIVATE_KEY"))

// generate JWT token
func GenerateJWT(user model.User) (string, error) {
	tokenTTL, _ := strconv.Atoi(os.Getenv("TOKEN_TTL")) //แปลง str -> int เพราะเวลา TOKEN_TTL เป็น str
	//log.Println(time.Now())
	//log.Println(time.Now().Add(time.Second * time.Duration(tokenTTL)))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{ //สร้าง token (sing token รูปแบบ HMAC , กำหนดข้อมูลใน token)
		"id":   user.ID,
		"role": user.RoleID,
		"iat":  time.Now().Unix(),                                            //เวลาที่ token ถูกสร้าง
		"eat":  time.Now().Add(time.Second * time.Duration(tokenTTL)).Unix(), // เวลาหมดอายุของ token
	})
	return token.SignedString(privateKey) //token จะถูกเซ็นด้วย privateKey
}

// validate JWT token ตรวจสอบ
func ValidateJWT(context *gin.Context) error {
	token, err := getToken(context)
	if err != nil {
		return err
	}
	_, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return nil
	}
	return errors.New("invalid token provided")
}

// validate Admin role
func ValidateAdminRoleJWT(context *gin.Context) error {
	token, err := getToken(context) //ตรวจ
	if err != nil {
		return err
	}
	claims, ok := token.Claims.(jwt.MapClaims) //แปลงข้อมูลที่เก็บไว้ในโทเค็น MapClaims
	userRole := uint(claims["role"].(float64)) //ดึงค่าของ role จากโทเค็น, แปลงค่านั้นจาก float64
	if ok && token.Valid && userRole == 1 {    //token.Valid อายุ
		return nil //"ไม่มีข้อผิดพลาด"
	}
	return errors.New("invalid admin token provided")
}

// validate CEO role เจ้าของหอพัก
func ValidateCeoRoleJWT(context *gin.Context) error {
	token, err := getToken(context) //ตรวจ
	if err != nil {
		return err
	}
	claims, ok := token.Claims.(jwt.MapClaims) //แปลงข้อมูลที่เก็บไว้ในโทเค็น MapClaims
	userRole := uint(claims["role"].(float64)) //ดึงค่าของ role จากโทเค็น, แปลงค่านั้นจาก float64
	if ok && token.Valid && userRole == 3 {    //token.Valid อายุ
		return nil //"ไม่มีข้อผิดพลาด"
	}
	fmt.Println("pass CEO")
	return errors.New("invalid Landlord(เจ้าของหอพัก) token provided")
}

// validate Customer role
func ValidateCustomerRoleJWT(context *gin.Context) error {
	token, err := getToken(context)
	if err != nil {
		return err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	userRole := uint(claims["role"].(float64))
	if ok && token.Valid && userRole == 2 || userRole == 1 {
		return nil
	}
	return errors.New("invalid customer or admin token provided")
}

// fetch user details from the token ดึงรายละเอียดผู้ใช้จากโทเค็น
// การเชื่อมโยง Room กับ UserID ช่วยให้ระบบสามารถติดตามว่าห้องนั้นถูกสร้างโดยผู้ใช้ใด.
func CurrentUser(context *gin.Context) model.User {
	err := ValidateJWT(context)
	if err != nil {
		return model.User{}
	}
	token, _ := getToken(context) //ตรวจความถูกต้องของ Token context นี้มีข้อมูลที่เกี่ยวข้องกับคำขอ HTTP ปัจจุบัน
	claims, _ := token.Claims.(jwt.MapClaims)
	userId := uint(claims["id"].(float64)) //ฟังก์ชันแยกส่วนข้อมูลในโทเค็น (claims) และระบุ ID ของผู้ใช้จากโทเค็นนั้น

	user, err := model.GetUserById(userId)
	if err != nil {
		return model.User{}
	}
	return user
}

// check token validity
func getToken(context *gin.Context) (*jwt.Token, error) {
	tokenString := getTokenFromRequest(context)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return privateKey, nil
	})
	return token, err
}

// extract token from request Authorization header แยกหัว
func getTokenFromRequest(context *gin.Context) string {
	bearerToken := context.Request.Header.Get("Authorization") // Header  ที่ใช้ในการเข้ารหัส-ถอดรหัส
	splitToken := strings.Split(bearerToken, " ")              // เเบงออกเป็น 2 ส่วน ยืนยันตัวตน , โทเค็นยืนยันตัวตนของผู้ใช้จริง
	if len(splitToken) == 2 {                                  // ดูว่ามี 2 ส่วนรึป่าว
		return splitToken[1]
	}
	return ""
}
