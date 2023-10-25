package controller

type PurchaseRequest struct {
	EventId        string `json:"event_id"`
	UserId         string `json:"user_id"`
	Quantity       int    `json:"quantity"`
	Payment_status string `json:"payment_status"`
}
