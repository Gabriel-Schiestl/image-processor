package models

type Data struct {
	Prediction string `json:"prediction"`
	ImgBase64  string `json:"image"`
}

type Message struct {
	Pattern string `json:"pattern"`
	Data    Data   `json:"data"`
}