package domain

type UseCase struct {
	UUID        int    `json:"id"`
	MeaningUUID int    `json:"meaningid"`
	Sample      string `json:"sample"`
}
