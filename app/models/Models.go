package models

import (
	"reflect"
	"time"
)

var Entities = make(map[string]any, 100)

const (
	ADMIN     = "ADMIN"
	ORGANIZER = "ORGANIZER"
	EVENT     = "EVENT"
)

// Used to register entities
func init() {
	Entities[reflect.ValueOf(Organizer{}).String()] = Organizer{}
	Entities[reflect.ValueOf(Event{}).String()] = Event{}
}

type Organizer struct {
	ID        uint64 `id:"ID" gorm:"primaryKey"`
	Username  string
	Password  string
	Name      string
	Role      string
	CreatedAt time.Time
}

type Event struct {
	ID   uint64 `id:"ID" gorm:"primaryKey"`
	Name string
	Organizer
	Date time.Time
}
