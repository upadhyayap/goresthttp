package domain

type User struct {
	Id        uint   `json:"id"` // mapping fields in json modal
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	email     string `json:"email"` // if field name starts with small case then it will not be exported in the json
}
