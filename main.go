package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type account struct {
	Username  string `json:"username"`
	Followers int    `json:"followers"`
}

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	port := os.Getenv("PORT")

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

	dataMap := map[string]account{}
	json.Unmarshal([]byte(dataJSON), &dataMap)

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Welcome to Homework Server-Deployment")
	})

	r.GET("/:userID", func(c *gin.Context) {
		name := c.Param("userID")
		c.JSON(http.StatusOK, dataMap[name])
	})

	r.GET("/followers/:username", func(c *gin.Context) {
		name := c.Param("username")
		result := 0
		for _, v := range dataMap {
			if v.Username == name {
				result = v.Followers
			}
		}
		dataResult := fmt.Sprintf("followers: %d", result)
		c.JSON(http.StatusOK, dataResult)
	})

	r.Run(":" + port)
}
