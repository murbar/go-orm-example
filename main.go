package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var err error

const port = ":8080"

func connectDb() {
	db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect")
	}
}

func root(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello Go"})
}

func handleRequests() {
	router := gin.Default()
	router.GET("/", root)
	router.GET("/users", GetUsers)
	router.POST("/users", AddUser)
	router.GET("/user/:id", GetUser)
	router.DELETE("/user/:id", DeleteUser)
	router.PUT("/user/:id", UpdateUser)
	router.Run(port)
}

func main() {
	connectDb()
	defer db.Close()

	db.AutoMigrate(&User{})

	handleRequests()
}
