package main

type order struct {
	ID    string `json:"id"`
	SKU   string `json:"sku"`
	Price int    `json:"price"`
}
