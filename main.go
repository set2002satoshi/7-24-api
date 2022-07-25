package main

import (


	"github.com/set2002satoshi/7-24-api/db"
	"github.com/set2002satoshi/7-24-api/router"
)

func main() {
	db.DBInit()
	db.DbConnect()
	router.SetRouter()


}