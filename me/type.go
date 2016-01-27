package me

import (
	"fmt"

	"github.com/toorop/govh"
	"github.com/toorop/govh/order"
)

// Bill represents the billing.Bill OVH type
type Bill struct {
	ID              string        `json:"billId"`
	PdfURL          string        `json:"pdfUrl"`
	Date            govh.DateTime `json:"date"`
	PriceWithoutTax order.Price   `json:"priceWithoutTax"`
	Tax             order.Price   `json:"tax"`
	PriceWithTax    order.Price   `json:"priceWithTax"`
	Password        string        `json:"password"`
	OrderID         int           `json:"orderID"`
	URL             string        `json:"url"`
}

// String is the Bill struct stringer
func (b Bill) String() string {
	out := "ID: " + b.ID + "\n"
	out += fmt.Sprintf("Oder ID: %d\n", b.OrderID)
	out += fmt.Sprintf("Date: %d\n", b.Date.Unix())
	out += "Price HT: " + b.PriceWithoutTax.Text + "\n"
	out += "Tax: " + b.Tax.Text + "\n"
	out += "Price TTC: " + b.PriceWithTax.Text + "\n"
	out += "URL: " + b.URL + "\n"
	out += "PDF URL: " + b.PdfURL + "\n"
	out += "Password: " + b.Password + "\n"
	return out
}
