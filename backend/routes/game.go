package routes

import (
	"github.com/gin-gonic/gin"
	"softball_record/models"
	"time"
)

func CreateGameRoutes(c *gin.Context) {
	game := models.GameMeta{}
	err := c.BindJSON(&game)
	if err != nil {
		c.JSON(400, gin.H{"result": "error", "message": err.Error()})
		return
	}
	if game.Date.IsZero() {
		game.Date = time.Now()
	}
	id, err := models.CreateGame(game.HomeTeam, game.AwayTeam, game.Date, game.Location)
	if err != nil {
		c.JSON(500, gin.H{"result": "error", "message": err.Error()})
	} else {
		c.JSON(200, gin.H{"result": "success", "message": "Game created", "id": id})
	}
}

func GetGameRoutes(c *gin.Context) {
	id := c.Param("id")
	game, err := models.GetGame(id)
	if err != nil {
		c.JSON(404, gin.H{"result": "error", "message": err.Error()})
	} else {
		c.JSON(200, gin.H{"result": "success", "game": game})
	}
}

func GetAllGamesRoutes(c *gin.Context) {
	games, err := models.GetAllGames()
	if err != nil {
		c.JSON(500, gin.H{"result": "error", "message": err.Error()})
	} else {
		c.JSON(200, gin.H{"result": "success", "games": games})
	}
}

func AddRecordRoutes(c *gin.Context) {
	id := c.Param("id")
	record := models.Record{}
	err := c.BindJSON(&record)
	if err != nil {
		c.JSON(400, gin.H{"result": "error", "message": err.Error()})
		return
	}
	err = models.AddRecord(id, record)
	if err != nil {
		c.JSON(500, gin.H{"result": "error", "message": err.Error()})
	} else {
		c.JSON(200, gin.H{"result": "success", "message": "Record added"})
	}
}

func UpdateRecordsRoutes(c *gin.Context) {
	id := c.Param("id")
	records := []models.Record{}
	err := c.BindJSON(&records)
	if err != nil {
		c.JSON(400, gin.H{"result": "error", "message": err.Error()})
		return
	}
	err = models.UpdateRecords(id, records)
	if err != nil {
		c.JSON(500, gin.H{"result": "error", "message": err.Error()})
	} else {
		c.JSON(200, gin.H{"result": "success", "message": "Records updated"})
	}
}

func SaveGameRoutes(c *gin.Context) {
    id := c.Param("id")
    err := models.SaveGame(id)
    if err != nil {
        c.JSON(500, gin.H{"result": "error", "message": err.Error()})
    } else {
        c.JSON(200, gin.H{"result": "success", "message": "Game saved"})
    }
}

func InitGameRoutes(router *gin.RouterGroup) {
	router.POST("/game", CreateGameRoutes)
	router.GET("/game/:id", GetGameRoutes)
	router.GET("/games", GetAllGamesRoutes)
	router.POST("/game/:id/record", AddRecordRoutes)
	router.PUT("/game/:id/record", UpdateRecordsRoutes)
    router.POST("/game/:id/save", SaveGameRoutes)
}
