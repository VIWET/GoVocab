package teststore

import "github.com/VIWET/GoVocab/app/internal/domain"

type store struct {
	lists    map[int]domain.List
	words    map[int]domain.Word
	meanings map[int]domain.Meaning
	usecases map[int]domain.UseCase
}

func NewStore() *store {
	return &store{
		lists:    make(map[int]domain.List),
		words:    make(map[int]domain.Word),
		meanings: make(map[int]domain.Meaning),
		usecases: make(map[int]domain.UseCase),
	}
}

func (s *store) GetCount() (int, int, int) {
	return len(s.words), len(s.meanings), len(s.usecases)
}
