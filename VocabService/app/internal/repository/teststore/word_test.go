package teststore_test

import (
	"fmt"
	"testing"

	"github.com/VIWET/GoVocab/app/internal/domain"
	"github.com/VIWET/GoVocab/app/internal/repository/teststore"
	"github.com/stretchr/testify/assert"
)

func TestWordRepository_Create(t *testing.T) {
	dto := &domain.WordCreateDTO{
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

	r := teststore.NewWordRepository()

	err := r.Create(dto)
	assert.NoError(t, err)
	words, meanings, usecases := r.DB.GetCount()
	assert.Equal(t, 1, words)
	assert.Equal(t, 1, meanings)
	assert.Equal(t, 2, usecases)
}

func TestWordRepository_GetSingleWord(t *testing.T) {
	r := teststore.NewWordRepository()

	out, err := r.GetSingleWord(1)
	assert.Error(t, err)
	assert.Nil(t, out)

	dto := &domain.WordCreateDTO{
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

	err = r.Create(dto)
	assert.NoError(t, err)

	out, err = r.GetSingleWord(1)
	assert.NoError(t, err)
	assert.NotNil(t, out)
	fmt.Println(*out)
}
