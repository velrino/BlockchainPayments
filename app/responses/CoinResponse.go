package responses

type CoinResponse struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Symbol string `json:"symbol"`
	PriceUsd string `json:"price_usd"`
	PriceBtc string `json:"price_btc"`
	PriceBrl string `json:"price_brl"`
	TotalSupply string `json:"total_supply"`
	MaxSupply string `json:"max_supply"`
	Rank string `json:"rank"`
	LastUpdated string `json:"last_updated"`
	MarketCapBrl string `json:"market_cap_brl"`
	Calculate float64 
}