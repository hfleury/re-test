package models

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type CalculateResponse struct {
	Packs map[int]int `json:"packs"`
}
