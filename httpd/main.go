package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	fmt.Println("Hello")

	r := gin.Default()

	r.GET("/", handler.homepage())
	/*r.GET("/nearBy/:location", handler.nearBy())
	r.GET("/myBookings/:uid", handler.myBookings())
	r.POST("/bookCab", handler.bookCab())*/

	r.Run()

}
