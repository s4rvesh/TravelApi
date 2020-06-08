package handler

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)



func nearBy(c *gin.Context) {

	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/travel")

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Connected")

	key := c.Param("location")
	ki, _ := strconv.Atoi(key)
	kb, kt := ki-3, ki+3

	results, err := db.Query("SELECT vname, location, vid FROM cabs WHERE location BETWEEN ? AND ?", kb, kt)

	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var cab Cab

		err = results.Scan(&cab.vname, &cab.location, &cab.vid)

		if err != nil {
			panic(err.Error())
		}

		c.JSON(200, gin.H{
			"vname":            cab.vname,
			"location":         cab.location,
			"vid":              cab.vid,
			"current_Location": key,
		})

	}

	defer results.Close()
	defer db.Close()

}
