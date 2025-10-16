package controller

import (
	"crud-golang/src/models"
	"github.com/gin-gonic/gin"
)

func FindUserByID(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if err := models.DB.First(&user, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "User not found!"})
		return
	}
	c.JSON(200, user)
}

func FindUserByEmail(c *gin.Context) {
	email := c.Param("email")
	var user models.User
	if err := models.DB.Where("email = ?", email).First(&user).Error; err != nil {
		c.JSON(404, gin.H{"error": "User not found!"})
		return
	}
	c.JSON(200, user)
}