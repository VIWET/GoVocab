package domain

type UseCase struct {
	ID        int    `json:"id"`
	MeaningID int    `json:"meaningid"`
	Sample    string `json:"sample"`
}

type UseCaseCreateDTO struct {
	Sample string `json:"sample"`
}
