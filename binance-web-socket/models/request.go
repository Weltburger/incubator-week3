package models

type Request struct {
	Method string    `json:"method"`
	Params [1]string `json:"params"`
	ID     int       `json:"id"`
}
