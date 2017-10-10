package db

type License struct {
	Type  string `json:"type"`
	Text  string `json:"text"`
	Promo string `json:"promo"`
}