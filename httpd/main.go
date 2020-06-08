package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type Bookings struct {
	uid      int    `json:"uid"`
	bid      int    `json:"bid"`
	startLoc string `json:"startLoc"`
	endLoc   string `json:"endLoc"`
}

type Cab struct {
	vid      int    `json:"vid"`
	vname    string `json:"vname"`
	location int    `json:"location"`
}

func main() {

	fmt.Println("Hello")

	r := gin.Default()

	r.GET("/", handler.homepage())
	r.GET("/nearBy/:location", handler.nearBy())
	r.GET("/myBookings/:uid", handler.myBookings())
	r.POST("/bookCab", handler.bookCab())

	r.Run()

}
