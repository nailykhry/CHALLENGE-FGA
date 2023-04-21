package controllers

import (
	"net/http"

	"CHALLENGE-3.2/database"
	"CHALLENGE-3.2/helpers"
	"CHALLENGE-3.2/models"

	"github.com/gin-gonic/gin"
)

func UserRegister(c *gin.Context) {
	var User models.User
	if err := c.ShouldBindJSON(&User); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	db := database.GetDB()

	if User.Email == "" || User.FullName == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "Email and Full Name are required",
		})
		return
	}

	err := db.Debug().Create(&User).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":        User.ID,
		"email":     User.Email,
		"full_name": User.FullName,
	})
}

func UserLogin(c *gin.Context) {
	db := database.GetDB()
	_ = db
	User := models.User{}

	if err := c.ShouldBindJSON(&User); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	password := User.Password
	err := db.Debug().Where("email = ?", User.Email).Take(&User).Error
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "invalid email/password",
		})
		return
	}

	comparePass := helpers.ComparePass([]byte(User.Password), []byte(password))

	if !comparePass {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "invalid email/password",
		})
		return
	}

	token := helpers.GenerateToken(User.ID, User.Email, User.IsAdmin)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})

}
