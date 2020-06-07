package handler

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
)

type Bookings struct {
	uid      int    `json:"uid"`
	bid      int    `json:"bid"`
	startLoc string `json:"startLoc"`
	endLoc   string `json:"endLoc"`
}

func myBookings(c *gin.Context) {

	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/travel")

	if err != nil {
		panic(err.Error())
	}

	uid := c.Param("uid")
	results, err := db.Query("SELECT uid, bid, startLoc,endLoc FROM bookings WHERE uid=?", uid)

	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var book Bookings

		err = results.Scan(&book.uid, &book.bid, &book.endLoc, &book.endLoc)

		if err != nil {
			panic(err.Error())
		}

		fmt.Print(book.uid)
		fmt.Print(book.bid)
		fmt.Print(book.startLoc)
		fmt.Print(book.endLoc)

		c.JSON(200, gin.H{
			"uid":      book.uid,
			"bid":      book.bid,
			"startLoc": book.startLoc,
			"endLoc":   book.endLoc,
		})
	}

	defer results.Close()
	defer db.Close()

}
