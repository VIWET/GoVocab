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

func (r *wordRepository) Create(lid int, dto *domain.WordCreateDTO) (*domain.Word, error) {
	word := &domain.Word{
		Text: dto.Text,
	}

	err := r.db.QueryRow(
		"INSERT INTO words (text) VALUES ($1) RETURNING id",
		dto.Text).Scan(&word.ID)
	if err != nil {
		return nil, err
	}

	err = r.db.QueryRow(
		"INSERT INTO words_lists_relation (list_id, word_id) VALUES ($1, $2) RETURNING word_id",
		lid, word.ID).Scan(&word.ID)
	if err != nil {
		return nil, err
	}

	return word, nil
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
