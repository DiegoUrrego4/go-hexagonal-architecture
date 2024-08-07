package domain

type UserRequest struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Age      uint8  `json:"age"`
}

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Age      uint8  `json:"age"`
}
