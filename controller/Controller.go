package controller

// The controller package in Go web server plays a crucial
// role in controlling the incoming HTTP requests and defining the
//  appropriate responses. It acts as an intermediary between the
//  web server's routing system and the business logic of your application.

import (
	"fmt"
	"net/http"
	"spotbuzz-backend/models"
	"spotbuzz-backend/pkg/dbhelper"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreatePlayerScore(c *gin.Context) {
	fmt.Println(" creating player")
	var player models.Player
	if err := c.ShouldBindJSON(&player); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	fmt.Println("Player details", player.Name, player.Country, player.Score)

	if len(player.Name) > 15 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Name should have a maximum of 15 characters"})
		return
	}
	err := dbhelper.InsertScores(player)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSONP(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    player,
	})

}

func GetAllPlayerScores(c *gin.Context) {
	// var response models.Response
	fmt.Println("this is a controller")
	players, err := dbhelper.GetPlayers()
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSONP(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    players,
	})

}
func UpdatePlayerScore(c *gin.Context) {

	playerID := c.Param("id")
	// Ensure that the player ID is a valid integer
	// id, err := strconv.Atoi(playerID)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid player ID"})
	// 	return
	// }
	var player models.Player

	if err := c.ShouldBindJSON(&player); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	if len(player.Name) > 15 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Name should have a maximum of 15 characters"})
		return
	}
	// Update the player's name and score in the database
	res, err := dbhelper.UpdatePlayer(player, playerID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error updating player"})
		return
	}

	// Return a success response
	c.JSON(http.StatusOK, gin.H{
		"message": "Player updated successfully",
		"status":  true,
		"data":    res,
	})

}

func DeletePlayerScore(c *gin.Context) {
	playerID := c.Param("id")

	// Ensure that the player ID is a valid integer
	// id, err := strconv.Atoi(playerID)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid player ID"})
	// 	return
	// }

	// Perform the deletion operation in the database based on the player ID
	err := dbhelper.DeletePlayer(playerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete player"})
		return
	}

	// Return a success response
	c.JSON(http.StatusOK, gin.H{"message": "Player deleted successfully"})

}

func GetPlayerByRankValue(c *gin.Context) {

	rankVal := c.Param("val")
	// Ensure that the player ID is a valid integer
	rank, err := strconv.Atoi(rankVal)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid player Rank"})
		return
	}

	// Fetch the player data based on the rank from the database
	player, err := dbhelper.GetByRankFromDatabase(rank)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch player"})
		return
	}

	// Return the player data as a JSON response
	c.JSON(http.StatusOK, player)
}

func GetRandomePlayer(c *gin.Context) {

	randPlayer, err := dbhelper.GetRandomePlayerfromDatabase()
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSONP(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    randPlayer,
	})

}
