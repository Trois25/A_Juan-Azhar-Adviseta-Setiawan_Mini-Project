package controller

type EventRequest struct {
	Poster_image    string  `json:"poster_image" form:"poster_image"`
	Title           string  `json:"title" form:"title"`
	Body            string  `json:"body" form:"body"`
	Ticket_quantity int     `json:"ticket_quantity" form:"ticket_quantity"`
	Price           float64 `json:"price" form:"price"`
	Place           string  `json:"place" form:"place"`
	Date            string  `json:"date" form:"date"`
}
