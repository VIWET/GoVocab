package domain

type Word struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}

type WordCreateDTO struct {
	Text     string             `json:"text"`
	Meanings []MeaningCreateDTO `json:"meaings"`
}

type WordOutputDTO struct {
	ID       int                `json:"id"`
	Text     string             `json:"text"`
	Meanings []MeaningOutputDTO `json:"meaings"`
}
