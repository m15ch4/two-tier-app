package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"micze.io/app/store"
)

func main() {
	store := store.NewSqlPlayerStore("localhost", 5432, "postgres", "VMware1!", "app")
	//store := store.NewStubPlayerStore()
	server := &PlayerServer{store}

	router := gin.Default()

	router.GET("/players/:name", server.showScore)
	router.GET("/players", server.showPlayers)
	router.POST("players/:name", server.processWin)
	router.Run(":8082")
}
