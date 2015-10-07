package jsonapi

type ErrorResponse struct {
	Errors []Error `json:"errors"`
}

type Error struct {
	Status string `json:"status"`
	Code   string `json:"code"`
	Detail string `json:"detail"`
}
