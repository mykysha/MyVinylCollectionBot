package entities

import "time"

type Info struct {
	StartTime time.Time
}

type User struct {
	ChatID   string
	UserName string
}

type Location struct {
	Owner User
	Name  string
}

type Artist struct {
	Name string
}

type Album struct {
	Artist      Artist
	Name        string
	Genre       string
	ReleaseYear int
	ReissueYear int
	Label       string
	Coloured    bool
	CoverID     string
}

type Collection struct {
	Albums    []Album
	Locations map[Album]Location
}

type Wishlist struct {
	Owner  User
	Albums []Album
	Stores map[Album]string
}
