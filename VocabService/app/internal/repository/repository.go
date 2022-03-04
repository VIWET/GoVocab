package repository

import "github.com/VIWET/GoVocab/app/internal/domain"

type ListRepository interface {
	Create(dto *domain.ListCreateDTO) (*domain.List, error)
	GetAllList(uid int) ([]*domain.List, error)
	GetList(uid int, id int) (*domain.ListOutputDTO, error)
	AddWord(id int, wid int) error
	Update(l *domain.List) error
	Delete(id int) error
}

type WordRepository interface {
	Create(lid int, dto *domain.WordCreateDTO) (*domain.WordOutputDTO, error)
	GetSingleWord(id int) (*domain.WordOutputDTO, error)
	GetRandomWords(n int) ([]*domain.WordOutputDTO, error)
	Update(w *domain.Word) error
	Delete(id int) error
}
