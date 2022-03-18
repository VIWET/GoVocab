package sqlstore

import (
	"database/sql"

	"github.com/VIWET/GoVocab/app/internal/domain"
	"github.com/VIWET/GoVocab/app/internal/errors"
	"github.com/VIWET/GoVocab/app/internal/repository"
)

type listRepository struct {
	db *sql.DB
}

func NewListRepository(db *sql.DB) repository.ListRepository {
	return &listRepository{
		db: db,
	}
}

func (r *listRepository) Create(l *domain.List) error {
	err := r.db.QueryRow(
		"INSERT INTO lists (user_id, title) VALUES ($1, $2) RETURNING id",
		l.UserID,
		l.Title).Scan(&l.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *listRepository) GetLists(uid int) ([]*domain.List, error) {
	rows, err := r.db.Query(
		"SELECT id, title FROM lists WHERE user_id = $1",
		uid)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var lists []*domain.List

	for rows.Next() {
		l := &domain.List{
			UserID: uid,
		}

		err := rows.Scan(&l.ID, &l.Title)
		if err != nil {
			return lists, err
		}

		lists = append(lists, l)
	}

	if len(lists) == 0 {
		return nil, errors.ErrRecordNotFound
	}

	return lists, nil
}

func (r *listRepository) GetList(id int) (*domain.List, error) {
	l := &domain.List{
		ID: id,
	}

	err := r.db.QueryRow(
		"SELECT title, user_id FROM lists WHERE id = $1",
		id).Scan(&l.Title, &l.UserID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.ErrRecordNotFound
		}
		return nil, err
	}

	return l, nil
}

func (r *listRepository) AddWord(id int, wid int) error {
	return nil
}

func (r *listRepository) Update(l *domain.List) error {
	return nil
}

func (r *listRepository) Delete(id int) error {
	return nil
}
