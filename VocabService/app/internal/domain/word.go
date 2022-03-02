package domain

type Word struct {
	ID       int       `json:"id"`
	Text     string    `json:"text"`
	Meanings []Meaning `json:"meaings"`
}
