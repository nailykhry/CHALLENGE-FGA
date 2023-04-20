package controllers

import (
	"PROJECT-2/database"
	"PROJECT-2/models"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

func CreateBook(ctx *gin.Context) {
	db := database.GetDB()
	var newBook models.Book

	if err := ctx.ShouldBindJSON(&newBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	book := models.Book{
		Name_book: newBook.Name_book,
		Author:    newBook.Author,
	}

	err := db.Create(&book).Error

	if err != nil {
		fmt.Println("Error creating book data:", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "New book created successfully",
		"data":    book,
	})
}

func GetBooks(ctx *gin.Context) {
	db := database.GetDB()
	book := []models.Book{}
	err := db.Find(&book).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("Book data not found")
			ctx.AbortWithError(http.StatusNotFound, err)
			return
		}
		print("Error finding book:", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
	}
	fmt.Printf("Book data: %+v\n", book)
	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Get books succcessfully showed",
		"data":    book,
	})
}

func GetBookByID(ctx *gin.Context) {
	id := ctx.Param("bookID")
	db := database.GetDB()
	book := models.Book{}
	err := db.First(&book, "id = ?", id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("Book data not found")
			ctx.AbortWithError(http.StatusNotFound, err)
			return
		}
		print("Error finding book:", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
	}
	fmt.Printf("Book data: %+v\n", book)
	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Get book showed successfully",
		"data":    book,
	})
}

func UpdateBook(ctx *gin.Context) {
	db := database.GetDB()
	book := models.Book{}
	id := ctx.Param("bookID")

	var newUpdate models.Book

	if err := ctx.ShouldBindJSON(&newUpdate); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err := db.Model(&book).Where("id = ?", id).Updates(models.Book{
		Name_book: newUpdate.Name_book,
		Author:    newUpdate.Author,
	}).Error

	if err != nil {
		fmt.Println("Error updating data:", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	fmt.Printf("update data's book: %+v \n", book)
	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Book has successfully updated",
		"data":    book,
	})
}

func DeleteBook(ctx *gin.Context) {
	db := database.GetDB()
	id := ctx.Param("bookID")
	book := models.Book{}
	err := db.Where("id = ?", id).Delete(&book).Error
	if err != nil {
		fmt.Println("Error deleting book:", err.Error())
		ctx.AbortWithError(http.StatusNotFound, err)
		return
	}
	fmt.Printf("book with id %v has been successfully deleted", id)
	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Book has successfully deleted",
	})
}
