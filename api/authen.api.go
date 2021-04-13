package api

import (
	"main/db"
	"main/model"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func SetupAuthenAPI(router *gin.Engine) {
	authenAPI := router.Group("api/v2")
	{
		authenAPI.POST("/login", login)
		authenAPI.POST("/register", register)
	}
}

func login(c *gin.Context) {
	var user model.User
	if c.ShouldBind(&user) == nil {
		var quriedUser model.User
		err := db.GetDB().First(&quriedUser, "username = ?", user.Username).Error
		if err != nil {
			c.JSON(200, gin.H{"result": "NG", "error": err})
		} else if !checkPasswordHash(user.Password, quriedUser.Password) {
			c.JSON(200, gin.H{"result": "NG", "error": "Invalid password"})
		} else {
			c.JSON(200, gin.H{"result": "OK", "data": user})
		}
	} else {
		c.JSON(401, gin.H{"result": "You fucked up"})
	}
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func register(c *gin.Context) {
	var user model.User
	if c.ShouldBind(&user) == nil {
		user.Password, _ = hashPassword(user.Password)
		user.CreateAt = time.Now()
		err := db.GetDB().Create(&user).Error
		if err != nil {
			c.JSON(200, gin.H{"result": "NG", "error": err})
		} else {
			c.JSON(200, gin.H{"result": "OK", "data": user})
		}
	} else {
		c.JSON(401, gin.H{"result": "You fucked up"})
	}
}
