package order

// Price is the price with it's currency and textual representation
type Price struct {
	CurrencyCode string  `json:"currencyCode"`
	Value        float64 `json:"value"`
	Text         string  `json:"text"`
}
