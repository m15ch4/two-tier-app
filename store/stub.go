package store

type stubPlayerStore struct {
	scores map[string]int
}

func (s *stubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func (s *stubPlayerStore) GetPlayers() []string {
	keys := make([]string, 0, len(s.scores))
	for k := range s.scores {
		keys = append(keys, k)
	}
	return keys
}

func (s *stubPlayerStore) RecordWin(name string) {
	s.scores[name] += 1
}

func NewStubPlayerStore() PlayerStore {
	return &stubPlayerStore{
		scores: map[string]int{
			"Pepper": 10,
			"Floyd":  20,
			"John":   1,
		},
	}
}
