package teststore

import (
	"github.com/VIWET/GoVocab/app/internal/domain"
	"github.com/VIWET/GoVocab/app/internal/errors"
	"github.com/VIWET/GoVocab/app/internal/repository"
)

type listRepository struct {
	db *store
}

func NewListRepository() repository.ListRepository {
	return &listRepository{
		db: NewStore(),
	}
}

func (r *listRepository) Create(l *domain.List) error {
	l.ID = len(r.db.lists) + 1

	r.db.lists[l.ID] = *l

	return nil
}

func (r *listRepository) GetLists(uid int) ([]*domain.List, error) {
	var lists []*domain.List

	for i := 0; i < len(r.db.lists); i++ {
		l, ok := r.db.lists[i+1]
		if !ok {
			continue
		}
		if l.UserID == uid {
			lists = append(lists, &l)
		}
	}

	if len(lists) == 0 {
		return nil, errors.ErrRecordNotFound
	}

	return lists, nil
}

func (r *listRepository) GetList(id int) (*domain.List, error) {
	l, ok := r.db.lists[id]
	if !ok {
		return nil, errors.ErrRecordNotFound
	}

	return &l, nil
}

// TODO
func (r *listRepository) AddWord(id int, wid int) error {
	return nil
}

func (r *listRepository) Update(l *domain.List) error {
	_, ok := r.db.lists[l.ID]
	if !ok {
		return errors.ErrRecordNotFound
	}

	r.db.lists[l.ID] = *l

	return nil
}

func (r *listRepository) Delete(id int) error {
	_, ok := r.db.lists[id]
	if !ok {
		return errors.ErrRecordNotFound
	}

	delete(r.db.lists, id)

	return nil
}
