package models

import "time"

type Organizer struct {
	name      string
	createdAt *time.Time
}

type Event struct {
	name string
	date *time.Time
}
