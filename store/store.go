package store

type PlayerStore interface {
	GetPlayerScore(name string) int
	GetPlayers() []string
	RecordWin(name string)
}
