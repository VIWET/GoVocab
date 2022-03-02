package repository

import "github.com/VIWET/GoVocab/app/internal/domain"

type WordRepository interface {
	Create(w *domain.Word) error
	GetSingleWord(id int) (domain.Word, error)
	GetRandomWords(n int) ([]domain.Word, error)
	Update(w *domain.Word) error
	Delete(id int) error
}

type MeaningRepository interface {
	Create(wid int, m *domain.Meaning) error
	Update(m *domain.Meaning) error
	Delete(id int) error
}

type UseCaseRepository interface {
	Create(mid int, m *domain.UseCase) error
	Update(uc *domain.UseCase) error
	Delete(id int) error
}
