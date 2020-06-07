package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func homepage(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "hello world",
	})

	fmt.Println("Endpoint Hit: homePage")

}
