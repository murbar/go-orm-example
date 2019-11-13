package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type User struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email"`
}

// GetUsers fetches all users
func GetUsers(c *gin.Context) {
	var users []User
	if err := db.Find(&users).Error; err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		fmt.Println(err)
	} else {
		c.JSON(http.StatusOK, users)
}
}

// GetUser gets a user by ID
func GetUser(c *gin.Context) {
	id := c.Param("id")
	var user User
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
	} else {
		c.JSON(http.StatusOK, user)
	}
}

// AddUser creates a new user
func AddUser(c *gin.Context) {
	var user User
	if c.BindJSON(&user) == nil {
		db.Create(&user)
		c.JSON(http.StatusCreated, user)
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to create new user"})
	}

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	var user User
	db.Where("name = ?", name).Find(&user)
	db.Delete(&user)

	fmt.Fprintf(w, "User deleted")
}

func UpdateUser(c *gin.Context) {
	var user User
	id := c.Param("id")

	db.Where("id = ?", id).Find(&user)

	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "User not found"})
	}

	c.BindJSON(&user)
	db.Save(&user)
	c.JSON(http.StatusOK, user)

}
