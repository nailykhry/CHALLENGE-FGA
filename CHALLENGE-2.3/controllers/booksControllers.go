package controllers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Book struct {
	BookID int    `json:"book_id"`
	Title  string `json:"title"`
	Price  int64  `json:"price"`
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "BOOKCOLLECTIONS"
)

var (
	db  *sql.DB
	err error
)

func DBConn() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err = sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected to database")
	return db
}

func CreateBook(ctx *gin.Context) {
	db := DBConn()
	var newBook = Book{}
	if err := ctx.ShouldBindJSON(&newBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	sqlStatement := `
        INSERT INTO books (title, price)
        VALUES ($1, $2)
        RETURNING bookid, title, price
    `
	err := db.QueryRow(sqlStatement, newBook.Title, newBook.Price).
		Scan(&newBook.BookID, &newBook.Title, &newBook.Price)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"book": newBook})
	db.Close()
}

func GetBooks(ctx *gin.Context) {
	db := DBConn()
	var results = []Book{}
	sqlStatement := `SELECT * from books`
	rows, err := db.Query(sqlStatement)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var book = Book{}
		err = rows.Scan(&book.BookID, &book.Title, &book.Price)
		if err != nil {
			panic(err)
		}
		results = append(results, book)
	}
	fmt.Println("Books datas:", results)
	ctx.JSON(http.StatusCreated, gin.H{
		"books": results,
	})
	db.Close()
}

func GetBooksByID(ctx *gin.Context) {
	db := DBConn()
	var results = []Book{}
	bookID := ctx.Param("bookID")
	sqlStatement := `SELECT * from books WHERE bookid = $1`
	rows, err := db.Query(sqlStatement, bookID)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var book = Book{}
		err = rows.Scan(&book.BookID, &book.Title, &book.Price)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error_status":  "Data Not Found",
				"error_message": fmt.Sprintf("book with id %v not found", bookID),
			})
			return
		}

		results = append(results, book)
	}
	fmt.Println("Book datas:", results)
	ctx.JSON(http.StatusOK, gin.H{
		"book": results,
	})
	db.Close()
}

func UpdateBook(ctx *gin.Context) {
	db := DBConn()
	bookID := ctx.Param("bookID")
	var UpdateBook Book

	if err := ctx.ShouldBindJSON(&UpdateBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	sqlStatement := `
		UPDATE books
		SET title = $2, price = $3
		WHERE bookid = $1;
	`
	res, err := db.Exec(sqlStatement, bookID, UpdateBook.Title, UpdateBook.Price)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("Book with id %v not found", bookID),
		})
		panic(err)
	}
	count, err := res.RowsAffected()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Error updating data",
			"error_message": fmt.Sprintf("Error updating data %v", bookID),
		})
		panic(err)
	}
	fmt.Println("Updated data:", count)
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Book with id %v has been successfully updated", bookID),
	})
	db.Close()
}

func DeleteBook(ctx *gin.Context) {
	db := DBConn()
	bookID := ctx.Param("bookID")
	sqlStatement := `DELETE from books WHERE bookid = $1;`
	res, err := db.Exec(sqlStatement, bookID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("book with id %v not found", bookID),
		})
		panic(err)
	}
	count, err := res.RowsAffected()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Error deleting data",
			"error_message": fmt.Sprintf("Error deleting data with id %v", bookID),
		})
		panic(err)
	}
	fmt.Println("Deleted data amount:", count)
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Book with id %v has been successfully deleted", bookID),
	})
	db.Close()
}
