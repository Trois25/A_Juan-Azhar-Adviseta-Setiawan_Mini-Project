package controller

type EventRequest struct {
	Title           string  `json:"title"`
	Body            string  `json:"body"`
	Ticket_quantity int     `json:"ticket_quantity"`
	Price           float64 `json:"price"`
	Place           string  `json:"place"`
	Date            string  `json:"date"`
}
