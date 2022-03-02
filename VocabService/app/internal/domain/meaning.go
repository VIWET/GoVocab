package domain

type Meaning struct {
	ID           int       `json:"id"`
	WordID       int       `json:"wordid"`
	TypeOfSpeech string    `json:"type"`
	Description  string    `json:"description"`
	Translation  string    `json:"translation"`
	UseCases     []UseCase `json:"useCases"`
}
