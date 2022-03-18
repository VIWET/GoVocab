package teststore

// import (
// 	"github.com/VIWET/GoVocab/app/internal/domain"
// 	"github.com/VIWET/GoVocab/app/internal/errors"
// 	"github.com/VIWET/GoVocab/app/internal/repository"
// )

// type wordRepository struct {
// 	db *store
// }

// func NewWordRepository(db *store) repository.WordRepository {
// 	return &wordRepository{
// 		db: db,
// 	}
// }

// func (r *wordRepository) Create(lid int, dto *domain.WordCreateDTO) (*domain.WordOutputDTO, error) {
// 	w := domain.Word{
// 		ID:   len(r.db.words) + 1,
// 		Text: dto.Text,
// 	}

// 	r.db.words[w.ID] = w

// 	for _, meaning := range dto.Meanings {
// 		m := domain.Meaning{
// 			ID:           len(r.db.meanings) + 1,
// 			WordID:       w.ID,
// 			TypeOfSpeech: meaning.TypeOfSpeech,
// 			Description:  meaning.Description,
// 			Translation:  meaning.Translation,
// 		}

// 		for _, useCase := range meaning.UseCases {
// 			uc := domain.UseCase{
// 				ID:        len(r.db.usecases) + 1,
// 				MeaningID: m.ID,
// 				Sample:    useCase.Sample,
// 			}

// 			r.db.usecases[uc.ID] = uc
// 		}

// 		r.db.meanings[m.ID] = m
// 	}

// 	return nil, nil
// }

// func (r *wordRepository) GetSingleWord(id int) (*domain.WordOutputDTO, error) {
// 	word, ok := r.db.words[id]
// 	if !ok {
// 		return nil, errors.ErrRecordNotFound
// 	}

// 	var meanings []domain.MeaningOutputDTO

// 	for _, m := range r.db.meanings {
// 		if m.WordID == id {

// 			var useCases []domain.UseCase

// 			for _, uc := range r.db.usecases {
// 				if uc.MeaningID == m.ID {
// 					useCases = append(useCases, uc)
// 				}
// 			}

// 			meanings = append(meanings, domain.MeaningOutputDTO{
// 				ID:           m.ID,
// 				WordID:       m.WordID,
// 				TypeOfSpeech: m.TypeOfSpeech,
// 				Description:  m.Description,
// 				Translation:  m.Translation,
// 				UseCases:     useCases,
// 			})
// 		}
// 	}

// 	out := &domain.WordOutputDTO{
// 		ID:       word.ID,
// 		Text:     word.Text,
// 		Meanings: meanings,
// 	}

// 	return out, nil
// }

// func (r *wordRepository) GetRandomWords(n int) ([]*domain.WordOutputDTO, error) {
// 	return nil, nil
// }

// func (r *wordRepository) Update(w *domain.Word) error {
// 	return nil

// }

// func (r *wordRepository) Delete(id int) error {
// 	return nil
// }
