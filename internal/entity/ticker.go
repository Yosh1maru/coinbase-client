package entity

type Ticker struct {
	Type     string `json:"type"`
	Datetime string `json:"time"`
	Symbol   string `json:"product_id"`
	BestBid  string `json:"best_bid"`
	BestAsk  string `json:"best_ask"`
}
