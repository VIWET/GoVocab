package sqlstore

import (
	"database/sql"

	"github.com/VIWET/GoVocab/app/internal/domain"
	"github.com/VIWET/GoVocab/app/internal/repository"
)

type wordRepository struct {
	db *sql.DB
}

func NewWordRepository(db *sql.DB) repository.WordRepository {
	return &wordRepository{
		db: db,
	}
}

func (r *wordRepository) Create(lid int, dto *domain.WordCreateDTO) (*domain.WordOutputDTO, error) {
	word := &domain.WordOutputDTO{
		Text: dto.Text,
	}

	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}

	err = tx.QueryRow(
		"INSERT INTO words (text) "+
			"VALUES ($1) "+
			"RETURNING id",
		dto.Text).Scan(&word.ID)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.QueryRow(
		"INSERT INTO words_lists_relation (list_id, word_id) "+
			"VALUES ($1, $2) "+
			"RETURNING word_id",
		lid,
		word.ID).Scan(&word.ID)
	if err != nil {
		return nil, err
	}

	for _, meaning := range dto.Meanings {
		m := domain.MeaningOutputDTO{
			WordID:       word.ID,
			TypeOfSpeech: meaning.TypeOfSpeech,
			Description:  meaning.Description,
			Translation:  meaning.Translation,
		}
		err := tx.QueryRow("INSERT INTO meanings (word_id, type_of_speech, description, translation) "+
			"VALUES ($1, $2, $3, $4) "+
			"RETURNING id",
			m.WordID,
			m.TypeOfSpeech,
			m.Description,
			m.Translation).Scan(&m.ID)
		if err != nil {
			return nil, err
		}
		for _, useCase := range meaning.UseCases {
			uc := domain.UseCase{
				MeaningID: m.ID,
				Sample:    useCase.Sample,
			}
			err := tx.QueryRow("INSERT INTO use_cases (meaning_id, sample) "+
				"VALUES ($1, $2) "+
				"RETURNING id",
				uc.MeaningID,
				uc.Sample).Scan(&uc.ID)
			if err != nil {
				return nil, err
			}
			m.UseCases = append(m.UseCases, uc)
		}
		word.Meanings = append(word.Meanings, m)
	}

	return word, tx.Commit()
}

func (r *wordRepository) GetSingleWord(id int) (*domain.WordOutputDTO, error) {
	return nil, nil
}

func (r *wordRepository) GetRandomWords(n int) ([]*domain.WordOutputDTO, error) {
	return nil, nil
}

func (r *wordRepository) Update(w *domain.Word) error {
	return nil
}

func (r *wordRepository) Delete(id int) error {
	return nil
}
