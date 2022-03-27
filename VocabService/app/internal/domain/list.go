package domain

type List struct {
	ID     int    `json:"id"`
	UserID int    `json:"uid"`
	Title  string `json:"title"`
}

type ListCreateDTO struct {
	UserID int             `json:"uid"`
	Title  string          `json:"title"`
	Words  []WordCreateDTO `json:"words"`
}

type ListOutputDTO struct {
	ID     int    `json:"id"`
	UserID int    `json:"uid"`
	Title  string `json:"title"`
	Words  []Word `json:"words"`
}
