package main

type order struct {
	ID    string `json:"id" csv:"id"`
	SKU   string `json:"sku" csv:"sku"`
	Price int    `json:"price" csv:"price"`
}

type file struct {
	ID       string `json:"id" db:"id"`
	File     []byte `json:"file" db:"file"`
	FileName string `json:"file_name" db:"file_name"`
}
