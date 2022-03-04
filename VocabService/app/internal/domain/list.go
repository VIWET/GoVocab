package domain

type List struct {
	ID     int `json:"id"`
	UserID int `json:"uid"`
}

type ListCreateDTO struct {
	UserID int `json:"uid"`
}
