package sqlstore

import (
	"database/sql"

	"github.com/VIWET/GoVocab/app/internal/domain"
	"github.com/VIWET/GoVocab/app/internal/repository"
)

type useCaseRepository struct {
	db *sql.DB
}

func NewUseCaseRepository(db *sql.DB) repository.UseCaseRepository {
	return &useCaseRepository{
		db: db,
	}
}

func (r *useCaseRepository) Create(mid int, dto *domain.UseCase) error {
	return nil
}

func (r *useCaseRepository) GetUseCases(wid int) ([]*domain.UseCase, error) {
	return nil, nil
}

func (r *useCaseRepository) GetUseCase(id int) (*domain.UseCase, error) {
	return nil, nil
}

func (r *useCaseRepository) Update(m *domain.UseCase) error {
	return nil
}

func (r *useCaseRepository) Delete(id int) error {
	return nil
}
