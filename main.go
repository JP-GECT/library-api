package main

type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "in search of lost time", Author: "marcel proust", Quantity: 2},
	{ID: "2", Title: "in search of lost time", Author: "marcel proust", Quantity: 5},
	{ID: "3", Title: "in search of lost time", Author: "marcel proust", Quantity: 6},
}
