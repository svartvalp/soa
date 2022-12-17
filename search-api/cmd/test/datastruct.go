package main

type Data struct {
	A   string    `json:"a,omitempty"`
	B   float64   `json:"b,omitempty"`
	Arr []float64 `json:"arr,omitempty"`
}
