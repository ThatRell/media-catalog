package main

import "time"

type Media struct {
	id        int
	title     string
	status    string
	rating    string
	createdAt time.Time
	updatedAt time.Time
}

/*
type Game struct {
	title string
	genre string
}

type Book struct {
	title string
	genre string
}

type Anime struct {
	title     string
	genre     string
	completed bool
}

type Manga struct {
}

type Manwha struct {
}

type Movie struct {
}

type Show struct {
}
*/
