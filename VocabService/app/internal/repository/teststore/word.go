package teststore

import (
	"github.com/VIWET/GoVocab/app/internal/domain"
	"github.com/VIWET/GoVocab/app/internal/errors"
)

type wordRepository struct {
	DB *store
}

func NewWordRepository() *wordRepository {
	return &wordRepository{
		DB: NewStore(),
	}
}

func (r *wordRepository) Create(dto *domain.WordCreateDTO) error {
	w := domain.Word{
		ID:   len(r.DB.words) + 1,
		Text: dto.Text,
	}

	r.DB.words[w.ID] = w

	for _, meaning := range dto.Meanings {
		m := domain.Meaning{
			ID:           len(r.DB.meanings) + 1,
			WordID:       w.ID,
			TypeOfSpeech: meaning.TypeOfSpeech,
			Description:  meaning.Description,
			Translation:  meaning.Translation,
		}

		for _, useCase := range meaning.UseCases {
			uc := domain.UseCase{
				ID:        len(r.DB.usecases) + 1,
				MeaningID: m.ID,
				Sample:    useCase.Sample,
			}

			r.DB.usecases[uc.ID] = uc
		}

		r.DB.meanings[m.ID] = m
	}

	return nil
}

func (r *wordRepository) GetSingleWord(id int) (*domain.WordOutputDTO, error) {
	word, ok := r.DB.words[id]
	if !ok {
		return nil, errors.ErrRecordNotFound
	}

	var meanings []domain.MeaningOutputDTO

	for _, m := range r.DB.meanings {
		if m.WordID == id {

			var useCases []domain.UseCase

			for _, uc := range r.DB.usecases {
				if uc.MeaningID == m.ID {
					useCases = append(useCases, uc)
				}
			}

			meanings = append(meanings, domain.MeaningOutputDTO{
				ID:           m.ID,
				WordID:       m.WordID,
				TypeOfSpeech: m.TypeOfSpeech,
				Description:  m.Description,
				Translation:  m.Translation,
				UseCases:     useCases,
			})
		}
	}

	out := &domain.WordOutputDTO{
		ID:       word.ID,
		Text:     word.Text,
		Meanings: meanings,
	}

	return out, nil
}

func (r *wordRepository) GetRandomWords(n int) ([]domain.Word, error) {
	return nil, nil
}

func (r *wordRepository) Update(w *domain.Word) error {
	return nil

}

func (r *wordRepository) Delete(id int) error {
	return nil
}
