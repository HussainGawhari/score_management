package main

import (
	"fmt"
	"spotbuzz-backend/pkg/dbhelper"
	"spotbuzz-backend/router"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {

	// DB connection
	db, err := dbhelper.CreateConnection()
	if err != nil {
		return
	}
	tab, err := dbhelper.CreateTable("Players")
	fmt.Println(tab)
	// database connection close
	defer db.Close()
	// create instance of gin framwork
	r := gin.New()
	// register routes to handle router
	router.Router(r)
	// server running on port 8000
	r.Run(":8000")

}
