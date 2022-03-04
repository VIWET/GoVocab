package sqlstore_test

import (
	"fmt"
	"testing"

	"github.com/VIWET/GoVocab/app/internal/domain"
	"github.com/VIWET/GoVocab/app/internal/errors"
	"github.com/VIWET/GoVocab/app/internal/repository/sqlstore"
	"github.com/stretchr/testify/assert"
)

func TestListRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestSQLDB(t, config)
	defer teardown("lists")

	r := sqlstore.NewListRepository(db)

	dto := &domain.ListCreateDTO{
		UserID: 1,
		Title:  "Test List",
	}

	l, err := r.Create(dto)
	assert.NoError(t, err)
	assert.NotNil(t, l)
}

func TestListRepository_GetAllLists(t *testing.T) {
	db, teardown := sqlstore.TestSQLDB(t, config)
	defer teardown("lists")

	r := sqlstore.NewListRepository(db)

	uid := 1

	lists, err := r.GetAllList(uid)
	assert.ErrorIs(t, errors.ErrRecordNotFound, err)
	assert.Nil(t, lists)

	for i := 0; i < 5; i++ {
		dto := &domain.ListCreateDTO{
			UserID: uid,
			Title:  fmt.Sprintf("Test List %d", i),
		}

		l, err := r.Create(dto)
		assert.NoError(t, err)
		assert.NotNil(t, l)
	}

	lists, err = r.GetAllList(uid)
	assert.NoError(t, err)
	assert.NotNil(t, lists)
	assert.Equal(t, 5, len(lists))
}

func TestListRepository_GetList(t *testing.T) {
	db, teardown := sqlstore.TestSQLDB(t, config)
	defer teardown("lists, words, words_lists_relation")

	rl := sqlstore.NewListRepository(db)

	dto := &domain.ListCreateDTO{
		UserID: 1,
		Title:  "Test List",
	}

	l, err := rl.Create(dto)
	assert.NoError(t, err)
	assert.NotNil(t, l)

	rw := sqlstore.NewWordRepository(db)

	for i := 0; i < 10; i++ {
		wdto := &domain.WordCreateDTO{
			Text: fmt.Sprintf("Word %d", i+1),
		}
		w, err := rw.Create(l.ID, wdto)
		assert.NoError(t, err)
		assert.NotNil(t, w)
	}

	list, err := rl.GetList(l.UserID, l.ID)

	assert.NoError(t, err)
	assert.NotNil(t, list)
	assert.NotEqual(t, 0, len(list.Words))
}
