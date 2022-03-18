package sqlstore_test

import (
	"testing"

	"github.com/VIWET/GoVocab/app/internal/domain"
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
