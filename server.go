package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"micze.io/app/store"
)

type PlayerScore struct {
	Score int    `json:"score"`
	Name  string `json:"name"`
}

type PlayerServer struct {
	store store.PlayerStore
}

func (p *PlayerServer) showScore(c *gin.Context) {
	name := c.Param("name")
	score := p.store.GetPlayerScore(name)
	scoreJson := PlayerScore{Name: name, Score: score}

	c.JSON(http.StatusOK, scoreJson)
}

func (p *PlayerServer) showPlayers(c *gin.Context) {
	names := p.store.GetPlayers()

	type player struct {
		Name string `json:"name,omitempty"`
	}

	players := make([]player, 0)

	for _, name := range names {
		players = append(players, player{Name: name})
	}

	c.JSON(http.StatusOK, players)
}

func (p *PlayerServer) processWin(c *gin.Context) {
	name := c.Param("name")
	p.store.RecordWin(name)

	c.String(http.StatusAccepted, "")
}
