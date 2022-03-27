package sqlstore

import (
	"database/sql"

	"github.com/VIWET/GoVocab/app/internal/domain"
	"github.com/VIWET/GoVocab/app/internal/repository"
)

type meaningRepository struct {
	db *sql.DB
}

func NewMeaningRepository(db *sql.DB) repository.MeaningRepository {
	return &meaningRepository{
		db: db,
	}
}

func (r *meaningRepository) Create(wid int, m *domain.Meaning) error {
	return nil
}

func (r *meaningRepository) GetMeanings(wid int) ([]*domain.Meaning, error) {
	return nil, nil
}

func (r *meaningRepository) GetMeaning(id int) (*domain.Meaning, error) {
	return nil, nil
}

func (r *meaningRepository) Update(m *domain.Meaning) error {
	return nil
}

func (r *meaningRepository) Delete(id int) error {
	return nil
}
