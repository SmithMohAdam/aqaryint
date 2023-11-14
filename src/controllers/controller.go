package controllers

import (
	"aqaryint/src/models"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostMethod(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"Data": string("From limited post method"),
	})
}

// CreateTable
func AddUser(c *gin.Context) {

	user := &models.User{}
	j, _ := json.Marshal(user)

	c.JSON(http.StatusOK, gin.H{
		"Data": string(j),
	})
}
func CreateTable(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"Data": string("Table been created successfuly"),
	})
}
