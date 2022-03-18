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

func (r *wordRepository) Create(lid int, w *domain.Word) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	err = tx.QueryRow(
		"INSERT INTO words (text) VALUES ($1) RETURNING id",
		w.Text).Scan(&w.ID)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.QueryRow(
		"INSERT INTO words_lists_relation (list_id, word_id) VALUES ($1, $2) RETURNING word_id",
		lid,
		w.ID).Scan(&w.ID)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (r *wordRepository) GetWords(lid int) ([]*domain.Word, error) {
	return nil, nil
}

func (r *wordRepository) GetWord(id int) (*domain.Word, error) {
	return nil, nil
}

func (r *wordRepository) GetSynonyms(id int) ([]*domain.Word, error) {
	return nil, nil
}

func (r *wordRepository) Update(w *domain.Word) error {
	return nil
}

func (r *wordRepository) Delete(id int) error {
	return nil
}
