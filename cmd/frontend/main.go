package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
}

type StubPlayerStore struct {
	scores map[string]int
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

type SqlPlayerStore struct {
	db *sql.DB
}

func (s *SqlPlayerStore) GetPlayerScore(name string) int {
	query := fmt.Sprintf("SELECT score FROM players WHERE name='%s';", name)
	score := 0
	err := s.db.QueryRow(query).Scan(&score)
	if err != nil {
		return -1
	}

	return score
}

func NewSqlPlayerStore(host string, port int32, user string, password string, dbname string) PlayerStore {
	store := &SqlPlayerStore{}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	store.db = db

	return store
}

type PlayerServer struct {
	store PlayerStore
}

func (p *PlayerServer) showScore(c *gin.Context) {
	name := c.Param("name")
	score := p.store.GetPlayerScore(name)

	c.String(http.StatusOK, "%v", score)
}

func main() {
	store := NewSqlPlayerStore("172.16.1.100", 5432, "appuser", "VMware1!", "app")
	server := &PlayerServer{store}

	router := gin.Default()

	router.GET("/players/:name", server.showScore)
	router.Run(":8082")
}
