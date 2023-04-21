package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"CHALLENGE-3.2/database"
	"CHALLENGE-3.2/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateProduct(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)

	Product := models.Product{}
	userID := uint(userData["id"].(float64))
	c.ShouldBindJSON(&Product)

	Product.UserID = userID

	err := db.Debug().Create(&Product).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
			"err":     err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, Product)
}

func GetAllProduct(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	isAdmin := userData["isAdmin"].(bool)
	userID := uint(userData["id"].(float64))
	Product := []models.Product{}

	c.ShouldBindJSON(&Product)

	var err error
	if isAdmin {
		err = db.Find(&Product).Error
	} else {
		err = db.Find(&Product, "user_id = ?", userID).Error
	}

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("Product data not found")
			c.AbortWithError(http.StatusNotFound, err)
			return
		}
		print("Error finding product:", err)
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Get all product succcessfully showed",
		"data":    Product,
	})
}

func GetBookById(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)

	userID := uint(userData["id"].(float64))
	Product := models.Product{}
	productId, _ := strconv.Atoi(c.Param("productId"))

	c.ShouldBindJSON(&Product)

	Product.UserID = userID
	Product.ID = uint(productId)

	err := db.Find(&Product, "id = ?", productId).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("Product data not found")
			c.AbortWithError(http.StatusNotFound, err)
			return
		}
		print("Error finding product:", err)
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Get product succcessfully showed",
		"data":    Product,
	})
}

func UpdateProduct(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	Product := models.Product{}
	productId, _ := strconv.Atoi(c.Param("productId"))
	userID := uint(userData["id"].(float64))

	c.ShouldBindJSON(&Product)

	Product.UserID = userID
	Product.ID = uint(productId)

	err := db.Model(&Product).Where("id = ?", productId).Updates(models.Product{Title: Product.Title, Description: Product.Description}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, Product)
}

func DeleteProduct(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	Product := models.Product{}
	productId, _ := strconv.Atoi(c.Param("productId"))
	userID := uint(userData["id"].(float64))

	c.ShouldBindJSON(&Product)

	Product.UserID = userID
	Product.ID = uint(productId)

	err := db.Where("id = ?", productId).Delete(&Product).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Product deleted successfully",
	})
}
