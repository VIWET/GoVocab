package sqlstore_test

import (
	"testing"

	"github.com/VIWET/GoVocab/app/internal/domain"
	"github.com/VIWET/GoVocab/app/internal/errors"
	"github.com/VIWET/GoVocab/app/internal/repository/sqlstore"
	"github.com/stretchr/testify/assert"
)

func TestWordRepository_Create(t *testing.T) {
	tests := []struct {
		valid       bool
		description string
		title       string
		user_id     int
		word        domain.Word
	}{
		{
			valid:       true,
			description: "valid example",
			title:       "test 1",
			user_id:     1,
			word: domain.Word{
				Text: "TEST",
			},
		},
		{
			valid:       false,
			description: "list doesn't exists",
			title:       "test 2",
			user_id:     1,
			word: domain.Word{
				Text: "TEST",
			},
		},
	}

	for _, test := range tests {
		db, teardown := sqlstore.TestSQLDB(t, config)

		rw := sqlstore.NewWordRepository(db)
		rl := sqlstore.NewListRepository(db)

		l := &domain.List{
			Title:  test.title,
			UserID: test.user_id,
		}
		rl.Create(l)

		if test.valid {
			err := rw.Create(l.ID, &test.word)
			assert.NoError(t, err)
			assert.NotEqual(t, 0, test.word.ID)
		} else {
			err := rw.Create(l.ID+1, &test.word)
			assert.Error(t, err)
		}

		teardown("lists", "words", "words_lists_relation")
	}
}

func TestWordRepository_GetWords(t *testing.T) {
	tests := []struct {
		valid       bool
		description string
		list        *domain.List
		words       []*domain.Word
	}{
		{
			valid:       true,
			description: "valid example",
			list: &domain.List{
				Title:  "TEST",
				UserID: 1,
			},
			words: []*domain.Word{
				{
					Text: "TEST 1",
				},
				{
					Text: "TEST 2",
				},
			},
		},
		{
			valid:       false,
			description: "list doesn't exists",
			list: &domain.List{
				Title:  "TEST",
				UserID: 1,
			},
			words: []*domain.Word{
				{
					Text: "TEST 1",
				},
			},
		},
	}

	for _, test := range tests {
		db, teardown := sqlstore.TestSQLDB(t, config)

		rw := sqlstore.NewWordRepository(db)
		rl := sqlstore.NewListRepository(db)

		rl.Create(test.list)
		for _, w := range test.words {
			rw.Create(test.list.ID, w)
		}

		if test.valid {
			words, err := rw.GetWords(test.list.ID)
			assert.NoError(t, err)
			assert.Equal(t, len(test.words), len(words))
		} else {
			words, err := rw.GetWords(test.list.ID + 1)
			assert.ErrorIs(t, err, errors.ErrRecordNotFound)
			assert.Nil(t, words)
		}

		teardown("lists", "words", "words_lists_relation")
	}
}
