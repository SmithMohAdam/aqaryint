package models

type Student struct {
	Id    int    `Json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	City  string `json:"city"`
	Class rune   `json:"class"`
	//Fit   bool   `json:fit`
}
