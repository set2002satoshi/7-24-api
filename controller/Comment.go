package controller

import (
	

	"github.com/set2002satoshi/7-24-api/db"
	"github.com/set2002satoshi/7-24-api/model"
	"github.com/gin-gonic/gin"
)



func CreateComment(c *gin.Context) {
	Db := db.DbConnect()
	var PostTable model.Comment
	if err := c.BindJSON(&PostTable); err != nil {
		response := map[string]interface{}{
			"message": "データの受け取りに失敗",
		}
		c.JSON(404, response)
		return 
	}
	if result := Db.Create(&PostTable); result.Error != nil {
		response := map[string]interface{}{
			"message": "404",
		}
		c.JSON(404, response)
		return
	}
	response := map[string]interface{}{
		"post": &PostTable,
		"message": "ok",
	}
	c.JSON(201, response)
}

