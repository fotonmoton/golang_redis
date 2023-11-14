package warehouse

type Product struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Qty    int    `json:"qty"`
	Active bool   `json:"active"`
}
