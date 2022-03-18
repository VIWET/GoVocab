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

func TestWordRepository_GetWord(t *testing.T) {
	tests := []struct {
		valid       bool
		description string
		list        *domain.List
		word        *domain.Word
	}{
		{
			valid:       true,
			description: "valid example",
			list: &domain.List{
				Title:  "TEST",
				UserID: 1,
			},
			word: &domain.Word{
				Text: "TEST 1",
			},
		},
		{
			valid:       false,
			description: "word doesn't exists",
			list: &domain.List{
				Title:  "TEST",
				UserID: 1,
			},
			word: &domain.Word{
				Text: "TEST 1",
			},
		},
	}

	for _, test := range tests {
		db, teardown := sqlstore.TestSQLDB(t, config)

		rw := sqlstore.NewWordRepository(db)
		rl := sqlstore.NewListRepository(db)

		rl.Create(test.list)
		rw.Create(test.list.ID, test.word)

		if test.valid {
			word, err := rw.GetWord(test.word.ID)
			assert.NoError(t, err)
			assert.Equal(t, test.word.Text, word.Text)
		} else {
			word, err := rw.GetWord(test.word.ID + 1)
			assert.ErrorIs(t, err, errors.ErrRecordNotFound)
			assert.Nil(t, word)
		}

		teardown("lists", "words", "words_lists_relation")
	}
}

func TestWordRepository_AddSynonym(t *testing.T) {
	tests := []struct {
		valid       bool
		description string
		list        *domain.List
		word        *domain.Word
		synonym     *domain.Word
	}{
		{
			valid:       true,
			description: "valid example",
			list: &domain.List{
				Title:  "TEST",
				UserID: 1,
			},
			word: &domain.Word{
				Text: "TEST 1",
			},
			synonym: &domain.Word{
				Text: "TEST 2",
			},
		},
		{
			valid:       false,
			description: "word doesn't exists",
			list: &domain.List{
				Title:  "TEST",
				UserID: 1,
			},
			word: &domain.Word{
				Text: "TEST 1",
			},
			synonym: &domain.Word{
				Text: "TEST 2",
			},
		},
	}

	for _, test := range tests {
		db, teardown := sqlstore.TestSQLDB(t, config)

		rw := sqlstore.NewWordRepository(db)
		rl := sqlstore.NewListRepository(db)

		rl.Create(test.list)
		rw.Create(test.list.ID, test.word)
		rw.Create(test.list.ID, test.synonym)

		if test.valid {
			err := rw.AddSynonym(test.word.ID, test.synonym.ID)
			assert.NoError(t, err)
		} else {
			err := rw.AddSynonym(test.word.ID, test.synonym.ID+1)
			assert.Error(t, err)
		}

		teardown("lists", "words", "words_lists_relation")
	}
}

func TestWordRepository_GetSynonyms(t *testing.T) {
	tests := []struct {
		valid       bool
		description string
		list        *domain.List
		word        *domain.Word
		synonym     *domain.Word
	}{
		{
			valid:       true,
			description: "valid example",
			list: &domain.List{
				Title:  "TEST",
				UserID: 1,
			},
			word: &domain.Word{
				Text: "TEST 1",
			},
			synonym: &domain.Word{
				Text: "TEST 2",
			},
		},
		{
			valid:       true,
			description: "valid example",
			list: &domain.List{
				Title:  "TEST",
				UserID: 1,
			},
			word: &domain.Word{
				Text: "TEST 1",
			},
			synonym: &domain.Word{
				Text: "TEST 2",
			},
		},
		{
			valid:       false,
			description: "word doesn't exists",
			list: &domain.List{
				Title:  "TEST",
				UserID: 1,
			},
			word: &domain.Word{
				Text: "TEST 1",
			},
			synonym: &domain.Word{
				Text: "TEST 2",
			},
		},
	}

	for _, test := range tests {
		db, teardown := sqlstore.TestSQLDB(t, config)

		rw := sqlstore.NewWordRepository(db)
		rl := sqlstore.NewListRepository(db)

		rl.Create(test.list)
		rw.Create(test.list.ID, test.word)
		rw.Create(test.list.ID, test.synonym)
		rw.AddSynonym(test.word.ID, test.synonym.ID)

		if test.valid {
			synonyms, err := rw.GetSynonyms(test.word.ID)
			assert.NoError(t, err)
			assert.NotNil(t, synonyms)
			assert.Equal(t, test.synonym.ID, synonyms[0].ID)
			synonyms, err = rw.GetSynonyms(test.synonym.ID)
			assert.Equal(t, test.word.ID, synonyms[0].ID)
			assert.NoError(t, err)
			assert.NotNil(t, synonyms)
			w := &domain.Word{
				Text: "Word without synonyms",
			}
			rw.Create(test.list.ID, w)
			synonyms, err = rw.GetSynonyms(w.ID)
			assert.NoError(t, err)
			assert.Empty(t, synonyms)
		} else {
			synonyms, err := rw.GetSynonyms(test.word.ID + 2)
			assert.NoError(t, err)
			assert.Nil(t, synonyms)

		}

		teardown("lists", "words", "words_lists_relation", "synonyms")
	}
}

func TestWordRepository_Update(t *testing.T) {
	tests := []struct {
		valid       bool
		description string
		list        *domain.List
		word        *domain.Word
		updatedText string
	}{
		{
			valid:       true,
			description: "valid example",
			list: &domain.List{
				Title:  "TEST",
				UserID: 1,
			},
			word: &domain.Word{
				Text: "TEST",
			},
			updatedText: "New test",
		},
		{
			valid:       false,
			description: "word doesn't exists",
			list: &domain.List{
				Title:  "TEST",
				UserID: 1,
			},
			word: &domain.Word{
				Text: "TEST",
			},
			updatedText: "New test",
		},
	}

	for _, test := range tests {
		db, teardown := sqlstore.TestSQLDB(t, config)

		rw := sqlstore.NewWordRepository(db)
		rl := sqlstore.NewListRepository(db)

		rl.Create(test.list)
		rw.Create(test.list.ID, test.word)

		if test.valid {
			test.word.Text = test.updatedText
			err := rw.Update(test.word)
			assert.NoError(t, err)
			w, _ := rw.GetWord(test.word.ID)
			assert.Equal(t, test.updatedText, w.Text)
		} else {
			test.word.Text = test.updatedText
			test.word.ID++
			err := rw.Update(test.word)
			assert.ErrorIs(t, err, errors.ErrRecordNotFound)
		}

		teardown("lists", "words", "words_lists_relation", "synonyms")
	}
}

func TestWordRepository_Delete(t *testing.T) {
	tests := []struct {
		valid       bool
		description string
		list        *domain.List
		word        *domain.Word
	}{
		{
			valid:       true,
			description: "valid example",
			list: &domain.List{
				Title:  "TEST",
				UserID: 1,
			},
			word: &domain.Word{
				Text: "TEST",
			},
		},
		{
			valid:       false,
			description: "word doesn't exists",
			list: &domain.List{
				Title:  "TEST",
				UserID: 1,
			},
			word: &domain.Word{
				Text: "TEST",
			},
		},
	}

	for _, test := range tests {
		db, teardown := sqlstore.TestSQLDB(t, config)

		rw := sqlstore.NewWordRepository(db)
		rl := sqlstore.NewListRepository(db)

		rl.Create(test.list)
		rw.Create(test.list.ID, test.word)

		if test.valid {
			err := rw.Delete(test.word.ID)
			assert.NoError(t, err)
			w, err := rw.GetWord(test.word.ID)
			assert.ErrorIs(t, err, errors.ErrRecordNotFound)
			assert.Nil(t, w)
		} else {
			err := rw.Delete(test.word.ID + 1)
			assert.ErrorIs(t, err, errors.ErrRecordNotFound)
		}

		teardown("lists", "words", "words_lists_relation", "synonyms")
	}
}
