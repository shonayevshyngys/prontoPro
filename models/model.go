package models

type User struct {
	id int64
	username string
}

type Provider struct {
	id int64
	description string
	rating float32
}