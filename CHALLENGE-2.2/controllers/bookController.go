package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	BookID    string `json:"book_id"`
	Title     string `json:"title"`
	Cover     string `json:"cover"`
	Author    string `json:"author"`
	Publisher string `json:"publisher"`
	Year      string `json:"year"`
	Genre     string `json:"genre"`
	Price     string `json:"price"`
}

var BookDatas = []Book{}

func GetAllBook(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{
		"books": BookDatas,
	})
}

func CreateBook(ctx *gin.Context) {
	var newBook Book

	if err := ctx.ShouldBindJSON(&newBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	newBook.BookID = fmt.Sprintf("b%d", len(BookDatas)+1)
	BookDatas = append(BookDatas, newBook)

	ctx.JSON(http.StatusCreated, gin.H{
		"book": newBook,
	})

}

func UpdateBook(ctx *gin.Context) {
	bookID := ctx.Param("bookID")
	condition := false
	var UpdateBook Book

	if err := ctx.ShouldBindJSON(&UpdateBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	for i, book := range BookDatas {
		if bookID == book.BookID {
			condition = true
			BookDatas[i] = UpdateBook
			BookDatas[i].BookID = bookID
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("Book with id %v not found", bookID),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Book with id %v has been successfully updated", bookID),
	})

}

func GetBook(ctx *gin.Context) {
	bookID := ctx.Param("bookID")
	condition := false
	var bookData Book

	for i, book := range BookDatas {
		if bookID == book.BookID {
			condition = true
			bookData = BookDatas[i]
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("book with id %v not found", bookID),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"book": bookData,
	})
}

func DeleteBook(ctx *gin.Context) {
	bookID := ctx.Param("bookID")
	condition := false
	var bookIndex int

	for i, book := range BookDatas {
		if bookID == book.BookID {
			condition = true
			bookIndex = i
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("book with id %v not found", bookID),
		})
		return
	}
	copy(BookDatas[bookIndex:], BookDatas[bookIndex+1:])
	BookDatas[len(BookDatas)-1] = Book{}
	BookDatas = BookDatas[:len(BookDatas)-1]

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Book with id %v has been successfully deleted", bookID),
	})
}
