package model

type Response struct {
	IsSuccess  bool
	StatusCode int
	Message    string
	Data       *[]Course
}

type Req struct {
	Query string `json:"query,omitempty"`
	Page  int    `json:"page,omitempty"`
}
type Course struct {
	Title    string `json:"title"`
	Provider string `json:"provider"`
	Level    string `json:"level"`
	Rating   string `json:"rating"`
}
type Pagination struct {
	Limit int    `json:"limit"`
	Page  int    `json:"page"`
	Sort  string `json:"sort"`
}
