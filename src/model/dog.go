package model

// Dog structure for holding dog information
type Dog struct {
	Name  string `json:"name"`
	Breed string `json:"breed"`
	Age   int    `json:"age"`
}
