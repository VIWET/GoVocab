package sqlstore_test

import (
	"testing"

	"github.com/VIWET/GoVocab/app/internal/domain"
	"github.com/VIWET/GoVocab/app/internal/repository/sqlstore"
	"github.com/stretchr/testify/assert"
)

func TestWordRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestSQLDB(t, config)
	defer teardown("lists", "words_lists_relation", "use_cases", "meanings", "words")

	rw := sqlstore.NewWordRepository(db)

	rl := sqlstore.NewListRepository(db)

	dtol := &domain.ListCreateDTO{
		UserID: 1,
		Title:  "Test List",
	}

	l, err := rl.Create(dtol)
	assert.NoError(t, err)
	assert.NotNil(t, l)

	dtow := &domain.WordCreateDTO{
		Text: "Car",
		Meanings: []domain.MeaningCreateDTO{
			{
				TypeOfSpeech: "N (C)",
				Description:  "Transport",
				Translation:  "Автомобиль",
				UseCases: []domain.UseCaseCreateDTO{
					{
						Sample: "A blue car",
					},
					{
						Sample: "I usually go to the work by car",
					},
				},
			},
		},
	}

	dtow2 := &domain.WordCreateDTO{
		Text: "Car 2",
		Meanings: []domain.MeaningCreateDTO{
			{
				TypeOfSpeech: "N (C)",
				Description:  "Transport 2",
				Translation:  "Автомобиль",
				UseCases: []domain.UseCaseCreateDTO{
					{
						Sample: "A blue car 2",
					},
					{
						Sample: "I usually go to the work by car 2",
					},
				},
			},
		},
	}

	word, err := rw.Create(l.ID, dtow)
	assert.NoError(t, err)
	assert.NotNil(t, word)
	word2, err := rw.Create(l.ID, dtow2)
	assert.NoError(t, err)
	assert.NotNil(t, word2)
}
