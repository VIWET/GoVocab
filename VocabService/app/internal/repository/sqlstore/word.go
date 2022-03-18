package sqlstore

import (
	"database/sql"

	"github.com/VIWET/GoVocab/app/internal/domain"
	"github.com/VIWET/GoVocab/app/internal/errors"
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
	rows, err := r.db.Query("SELECT w.id, w.text FROM words AS w INNER JOIN words_lists_relation AS wl ON w.id = wl.word_id WHERE wl.list_id = $1", lid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var words []*domain.Word

	for rows.Next() {
		word := &domain.Word{}
		err := rows.Scan(&word.ID, &word.Text)
		if err != nil {
			return words, nil
		}

		words = append(words, word)
	}

	if len(words) == 0 {
		return nil, errors.ErrRecordNotFound
	}

	return words, nil
}

func (r *wordRepository) GetWord(id int) (*domain.Word, error) {
	w := &domain.Word{
		ID: id,
	}
	err := r.db.QueryRow(
		"SELECT text FROM words WHERE id = $1",
		id).Scan(&w.Text)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.ErrRecordNotFound
		}
		return nil, err
	}

	return w, nil
}

func (r *wordRepository) AddSynonym(wid int, sid int) error {
	rows, err := r.db.Query(
		"INSERT INTO synonyms (word_id, synonym_id) VALUES ($1, $2), ($2, $1)",
		wid,
		sid)
	if err != nil {
		return err
	}

	defer rows.Close()

	err = rows.Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *wordRepository) GetSynonyms(id int) ([]*domain.Word, error) {
	rows, err := r.db.Query(
		"SELECT s.id, s.text FROM words AS s INNER JOIN synonyms AS sw ON sw.synonym_id = s.id WHERE sw.word_id = $1",
		id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var synonyms []*domain.Word

	for rows.Next() {
		s := &domain.Word{}
		err := rows.Scan(&s.ID, &s.Text)
		if err != nil {
			return synonyms, err
		}

		synonyms = append(synonyms, s)
	}

	return synonyms, nil
}

func (r *wordRepository) Update(w *domain.Word) error {
	err := r.db.QueryRow(
		"UPDATE words SET text = $1 WHERE id = $2 RETURNING id",
		w.Text,
		w.ID).Scan(&w.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.ErrRecordNotFound
		}
		return err
	}

	return nil
}

func (r *wordRepository) Delete(id int) error {
	return nil
}
