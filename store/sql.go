package store

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type sqlPlayerStore struct {
	db *sql.DB
}

func (s *sqlPlayerStore) GetPlayerScore(name string) int {
	query := fmt.Sprintf("SELECT score FROM players WHERE name='%s';", name)
	score := 0
	err := s.db.QueryRow(query).Scan(&score)
	if err == sql.ErrNoRows {
		return 0
	}
	if err != nil {
		fmt.Println(err.Error())
		return -1
	}

	return score
}

func (s *sqlPlayerStore) GetPlayers() []string {
	query := "SELECT name FROM players;"
	rows, err := s.db.Query(query)
	if err != nil {
		return []string{}
	}
	defer rows.Close()

	names := make([]string, 0)
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			continue
		}
		names = append(names, name)
	}

	if err := rows.Err(); err != nil {
		log.Println(err.Error())
	}

	return names
}

func (s *sqlPlayerStore) RecordWin(name string) {
	score := s.GetPlayerScore(name)
	var query string

	switch {
	case score > 0:
		query = fmt.Sprintf("UPDATE players SET score = %d WHERE name = '%s';", score+1, name)
		s.db.Exec(query)
	case score == 0:
		query = fmt.Sprintf("INSERT INTO players VALUES ( '%s', %d )", name, 1)
		s.db.Exec(query)
	}
}

func NewSqlPlayerStore(host string, port int32, user string, password string, dbname string) PlayerStore {
	store := &sqlPlayerStore{}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	store.db = db

	return store
}
