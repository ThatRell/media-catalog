package main

import "time"

type Media struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Rating    int       `json:"rating"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"createdat"`
	UpdatedAt time.Time `json:"updatedat"`
}
