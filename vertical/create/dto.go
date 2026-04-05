package create

type Request struct {
	CustomerID string `json:"customer_id"`
	ItemID     string `json:"item_id"`
	Quantity   int    `json:"quantity"`
}

type Response struct {
	OrderID string `json:"order_id"`
	Status  string `json:"status"`
}
