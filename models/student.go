package student

type Student struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}

var Students []Student
