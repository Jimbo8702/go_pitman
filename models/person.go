package main

type Person struct {
	Email string `json:"email"`
	Address map[string]interface{} `json:"address"`
	Name string `json:"name"`
	Age float64 `json:"age"`
}

