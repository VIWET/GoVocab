package teststore_test

// import (
// 	"testing"

// 	"github.com/VIWET/GoVocab/app/internal/domain"
// 	"github.com/VIWET/GoVocab/app/internal/repository/teststore"
// 	"github.com/stretchr/testify/assert"
// )

// func TestWordRepository_Create(t *testing.T) {
// 	dto := &domain.WordCreateDTO{
// 		Text: "Car",
// 		Meanings: []domain.MeaningCreateDTO{
// 			{
// 				TypeOfSpeech: "N (C)",
// 				Description:  "Transport",
// 				Translation:  "Автомобиль",
// 				UseCases: []domain.UseCaseCreateDTO{
// 					{
// 						Sample: "A blue car",
// 					},
// 					{
// 						Sample: "I usually go to the work by car",
// 					},
// 				},
// 			},
// 		},
// 	}

// 	s := teststore.NewStore()

// 	r := teststore.NewWordRepository(s)

// 	_, err := r.Create(1, dto)
// 	assert.NoError(t, err)
// 	words, meanings, usecases := s.GetCount()
// 	assert.Equal(t, 1, words)
// 	assert.Equal(t, 1, meanings)
// 	assert.Equal(t, 2, usecases)
// }

// func TestWordRepository_GetSingleWord(t *testing.T) {
// 	s := teststore.NewStore()
// 	r := teststore.NewWordRepository(s)

// 	out, err := r.GetSingleWord(1)
// 	assert.Error(t, err)
// 	assert.Nil(t, out)

// 	dto := &domain.WordCreateDTO{
// 		Text: "Car",
// 		Meanings: []domain.MeaningCreateDTO{
// 			{
// 				TypeOfSpeech: "N (C)",
// 				Description:  "Transport",
// 				Translation:  "Автомобиль",
// 				UseCases: []domain.UseCaseCreateDTO{
// 					{
// 						Sample: "A blue car",
// 					},
// 					{
// 						Sample: "I usually go to the work by car",
// 					},
// 				},
// 			},
// 		},
// 	}

// 	_, err = r.Create(1, dto)
// 	assert.NoError(t, err)

// 	out, err = r.GetSingleWord(1)
// 	assert.NoError(t, err)
// 	assert.NotNil(t, out)
// }
