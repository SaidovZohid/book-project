package main

type ResponseError struct {
	Message string `json:"message"`
}

type ResponOK struct {
	Message string `json:"message"`
}

type CreateOrUpdateRequest struct {
	Title      string    `json:"title"`
	AuthorName string    `json:"author_name"`
	Price      float64   `json:"price"`
	Amount     int       `json:"amount"`
}