package controllers

import (
	"aqaryint/aqaryint/src/models"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostMethod(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"Data": string("From limited post method"),
	})
}

func AddUser(c *gin.Context) {

	user := &models.User{}
	j, _ := json.Marshal(user)

	c.JSON(http.StatusOK, gin.H{
		"Data": string(j),
	})
}
