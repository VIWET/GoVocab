package domain

type Meaning struct {
	ID           int    `json:"id"`
	WordID       int    `json:"wordid"`
	TypeOfSpeech string `json:"type"`
	Description  string `json:"description"`
	Translation  string `json:"translation"`
}

type MeaningCreateDTO struct {
	TypeOfSpeech string             `json:"type"`
	Description  string             `json:"description"`
	Translation  string             `json:"translation"`
	UseCases     []UseCaseCreateDTO `json:"useCases"`
}

type MeaningOutputDTO struct {
	ID           int       `json:"id"`
	WordID       int       `json:"wordid"`
	TypeOfSpeech string    `json:"type"`
	Description  string    `json:"description"`
	Translation  string    `json:"translation"`
	UseCases     []UseCase `json:"useCases"`
}
