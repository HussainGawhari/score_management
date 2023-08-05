package router

import (
	"spotbuzz-backend/controller"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {

	//  I have handled the routers in this package with basic router handling
	// middleware or authentication middleware must be implemented here to check but for now I made it simple
	// V1 group routes
	v1 := r.Group("/v1")
	//Spotbuzz routes

	v1.POST("/player", controller.CreatePlayerScore)
	v1.PUT("/player/:id", controller.UpdatePlayerScore)
	v1.DELETE("player/:id", controller.DeletePlayerScore)
	v1.GET("/players", controller.GetAllPlayerScores)
	v1.GET("/", controller.GetAllPlayerScores)
	v1.GET("/players/rank/:val", controller.GetPlayerByRankValue)
	v1.GET("/players/randome", controller.GetRandomePlayer)

}
