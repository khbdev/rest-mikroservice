package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var Users []User


func AddUser(u User) {
	Users = append(Users, u)
}

func GetUsers() []User {
	return Users
}


func GetUsersHandler(c *gin.Context) {
	users := GetUsers()
	c.JSON(http.StatusOK, users)
}

func CreateUserHandler(c *gin.Context) {
	var newUser User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	AddUser(newUser)
	c.JSON(http.StatusOK, newUser)
}


func main() {
	r := gin.Default()


	r.GET("/users", GetUsersHandler)
	r.POST("/users", CreateUserHandler)


	r.Run(":8083")
}
