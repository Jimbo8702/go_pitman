package main

type Person struct {
	Name string `json:"name"`
	Age float64 `json:"age"`
	Email string `json:"email"`
	Address map[string]interface{} `json:"address"`
}

