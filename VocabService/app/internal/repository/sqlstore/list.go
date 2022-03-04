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

func (r *listRepository) Create(dto *domain.ListCreateDTO) (*domain.List, error) {
	list := &domain.List{
		UserID: dto.UserID,
		Title:  dto.Title,
	}

	err := r.db.QueryRow(
		"INSERT INTO lists (user_id, title) "+
			"VALUES ($1, $2) "+
			"RETURNING id",
		dto.UserID,
		dto.Title).Scan(&list.ID)
	if err != nil {
		return nil, err
	}

	return list, nil
}

func (r *listRepository) GetAllList(uid int) ([]*domain.List, error) {
	rows, err := r.db.Query(
		"SELECT id, user_id, title "+
			"FROM lists "+
			"WHERE user_id = $1", uid)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.ErrRecordNotFound
		}
		return nil, err
	}

	defer rows.Close()

	var lists []*domain.List

	for rows.Next() {
		var list domain.List
		if err := rows.Scan(&list.ID, &list.UserID, &list.Title); err != nil {
			return lists, err
		}

		lists = append(lists, &list)
	}

	if err = rows.Err(); err != nil {
		return lists, err
	}

	if len(lists) == 0 {
		return lists, errors.ErrRecordNotFound
	}

	return lists, nil
}

func (r *listRepository) GetList(uid int, id int) (*domain.ListOutputDTO, error) {
	rows, err := r.db.Query(
		"SELECT l.title, w.id, w.text "+
			"FROM lists AS l "+
			"INNER JOIN words_lists_relation AS wl ON l.id = wl.list_id "+
			"INNER JOIN words AS w ON w.id = wl.word_id "+
			"WHERE l.user_id = $1 AND l.id = $2", uid, id)
	if err != nil {
		return nil, err
	}

	list := &domain.ListOutputDTO{
		ID:     id,
		UserID: uid,
	}

	for rows.Next() {
		var word domain.Word
		if err := rows.Scan(&list.Title, &word.ID, &word.Text); err != nil {
			return list, err
		}

		list.Words = append(list.Words, word)
	}

	defer rows.Close()

	return list, nil
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
