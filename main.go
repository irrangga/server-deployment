package main

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Account struct {
	Username  string `json:"username"`
	Followers int    `json:"followers"`
}

func main() {
	var dataJSON = `{
		"sammy": {
				"username": "SammyShark",
				"followers": 987
		},
		"jesse": {
				"username": "JesseOctopus",
				"followers": 432
		},
		"drew": {
				"username": "DrewSquid",
				"followers": 321
		},
		"jamie": {
				"username": "JamieMantisShrimp",
				"followers": 654
		}
	}`

	dataMap := map[string]Account{}
	json.Unmarshal([]byte(dataJSON), &dataMap)

	r := gin.Default()

	r.GET("/:userID", func(c *gin.Context) {
		name := c.Param("userID")
		c.JSON(http.StatusOK, dataMap[name])
	})

	r.GET("/follower/:username", func(c *gin.Context) {
		name := c.Param("username")
		result := 0
		for _, v := range dataMap {
			if v.Username == name {
				result = v.Followers
			}
		}
		c.JSON(http.StatusOK, result)
	})

	r.Run(":8080")
}
