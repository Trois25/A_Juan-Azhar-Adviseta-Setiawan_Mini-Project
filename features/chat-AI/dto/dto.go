package dto

type RequestData struct {
	Genre  string `json:"genre"`
	Description   string `json:"description"`
	Year string `json:"year"`
}

type Response struct {
	Status         string
	Recommendation string
}
