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

func (r *wordRepository) Create(dto *domain.WordCreateDTO) error {
	return nil
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
