package repository

import "github.com/VIWET/GoVocab/app/internal/domain"

type WordRepository interface {
	Create(dto *domain.WordCreateDTO) error
	GetSingleWord(id int) (*domain.WordOutputDTO, error)
	GetRandomWords(n int) ([]*domain.WordOutputDTO, error)
	Update(w *domain.Word) error
	Delete(id int) error
}

type MeaningRepository interface {
	Create(wid int, m *domain.MeaningCreateDTO) error
	Update(m *domain.Meaning) error
	Delete(id int) error
}

type UseCaseRepository interface {
	Create(mid int, m *domain.UseCaseCreateDTO) error
	Update(uc *domain.UseCase) error
	Delete(id int) error
}
