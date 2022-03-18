package sqlstore_test

import (
	"testing"

	"github.com/VIWET/GoVocab/app/internal/domain"
	"github.com/VIWET/GoVocab/app/internal/errors"
	"github.com/VIWET/GoVocab/app/internal/repository/sqlstore"
	"github.com/stretchr/testify/assert"
)

func TestListRepository_Create(t *testing.T) {
	tests := []struct {
		valid       bool
		description string
		title       string
		user_id     int
	}{
		{
			valid:       true,
			description: "valid example",
			title:       "test 1",
			user_id:     1,
		},
	}

	for _, test := range tests {
		db, teardown := sqlstore.TestSQLDB(t, config)

		r := sqlstore.NewListRepository(db)

		l := &domain.List{
			Title:  test.title,
			UserID: test.user_id,
		}

		if test.valid {
			err := r.Create(l)
			assert.NoError(t, err)
			assert.NotEqual(t, 0, l.ID)
		} else {
			err := r.Create(l)
			assert.Error(t, err)
		}

		teardown("lists")
	}
}

func TestListRepository_GetLists(t *testing.T) {
	tests := []struct {
		valid       bool
		description string
		userId      int
		count       int
		lists       []*domain.List
		err         error
	}{
		{
			valid:       true,
			description: "two lists belong to user",
			userId:      1,
			count:       2,
			lists: []*domain.List{
				{
					Title:  "first list",
					UserID: 1,
				},
				{
					Title:  "second list",
					UserID: 1,
				},
			},
			err: nil,
		},
		{
			valid:       true,
			description: "only one list belong to user",
			userId:      1,
			count:       1,
			lists: []*domain.List{
				{
					Title:  "first list",
					UserID: 1,
				},
				{
					Title:  "second list",
					UserID: 2,
				},
			},
			err: nil,
		},
		{
			valid:       false,
			description: "no records",
			userId:      1,
			count:       0,
			lists:       []*domain.List{},
			err:         errors.ErrRecordNotFound,
		},
	}

	for _, test := range tests {
		db, teardown := sqlstore.TestSQLDB(t, config)

		r := sqlstore.NewListRepository(db)

		for _, l := range test.lists {
			r.Create(l)
		}

		if test.valid {
			lists, err := r.GetLists(test.userId)
			assert.NoError(t, err)
			assert.Equal(t, test.count, len(lists))
		} else {
			lists, err := r.GetLists(test.userId)
			assert.ErrorIs(t, err, test.err)
			assert.Equal(t, test.count, len(lists))
		}

		teardown("lists")
	}
}

func TestListRepository_GetList(t *testing.T) {
	tests := []struct {
		valid       bool
		description string
		lists       []*domain.List
		err         error
	}{
		{
			valid:       true,
			description: "test without errors",
			lists: []*domain.List{
				{
					Title:  "first list",
					UserID: 1,
				},
				{
					Title:  "second list",
					UserID: 1,
				},
			},
			err: nil,
		},
		{
			valid:       false,
			description: "check id bigger than id of created list",
			lists: []*domain.List{
				{
					Title:  "first list",
					UserID: 1,
				},
			},
			err: errors.ErrRecordNotFound,
		},
	}

	for _, test := range tests {
		db, teardown := sqlstore.TestSQLDB(t, config)

		r := sqlstore.NewListRepository(db)

		for _, l := range test.lists {
			r.Create(l)
		}

		if test.valid {
			for _, tl := range test.lists {
				list, err := r.GetList(tl.ID)
				assert.NoError(t, err)
				assert.Equal(t, tl.ID, list.ID)
			}
		} else {
			tl := test.lists[0]
			list, err := r.GetList(tl.ID + 1)
			assert.ErrorIs(t, err, test.err)
			assert.Nil(t, list)
		}

		teardown("lists")
	}
}

// TODO
func TestListRepository_AddWord(t *testing.T) {}

func TestListRepository_Update(t *testing.T) {
	tests := []struct {
		valid       bool
		description string
		oldTitle    string
		title       string
		list        *domain.List
		err         error
	}{
		{
			valid:       true,
			description: "test without errors",
			oldTitle:    "first list",
			title:       "new first list",
			list: &domain.List{
				Title:  "first list",
				UserID: 1,
			},
			err: nil,
		},
		{
			valid:       false,
			description: "update list with different id",
			oldTitle:    "first list",
			title:       "new first list",
			list: &domain.List{
				Title:  "first list",
				UserID: 1,
			},
			err: errors.ErrRecordNotFound,
		},
	}

	for _, test := range tests {
		db, teardown := sqlstore.TestSQLDB(t, config)

		r := sqlstore.NewListRepository(db)

		r.Create(test.list)

		if test.valid {
			test.list.Title = test.title
			err := r.Update(test.list)
			assert.NoError(t, err)
			tl, _ := r.GetList(test.list.ID)
			assert.Equal(t, test.title, tl.Title)
		} else {
			test.list.Title = test.title
			test.list.ID++
			err := r.Update(test.list)
			assert.ErrorIs(t, err, test.err)
		}

		teardown("lists")
	}
}

func TestListRepository_Delete(t *testing.T) {

}
