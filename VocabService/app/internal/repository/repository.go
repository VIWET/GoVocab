package repository

import "github.com/VIWET/GoVocab/app/internal/domain"

type ListRepository interface {
	Create(l *domain.List) error
	GetLists(uid int) ([]*domain.List, error)
	GetList(id int) (*domain.List, error)
	AddWord(id int, wid int) error
	Update(l *domain.List) error
	Delete(id int) error
}

type WordRepository interface {
	Create(lid int, w *domain.Word) error
	GetWords(lid int) ([]*domain.Word, error)
	GetWord(id int) (*domain.Word, error)
	AddSynonym(wid int, sid int) error
	GetSynonyms(id int) ([]*domain.Word, error)
	Update(w *domain.Word) error
	Delete(id int) error
}

type MeaningRepository interface {
	Create(wid int, m *domain.Meaning) error
	GetMeanings(wid int) ([]*domain.Meaning, error)
	GetMeaning(id int) (*domain.Meaning, error)
	Update(m *domain.Meaning) error
	Delete(id int) error
}

type UseCaseRepository interface {
	Create(mid int, uc *domain.UseCase) error
	GetUseCases(wid int) ([]*domain.UseCase, error)
	GetUseCase(id int) (*domain.UseCase, error)
	Update(m *domain.UseCase) error
	Delete(id int) error
}
