package models

import (
	"reflect"
	"time"
)

var Entities = make(map[string]any, 100)

const (
	ADMIN     = "ADMIN"
	ORGANIZER = "ORGANIZER"
)

// Used to register entities
func init() {
	Entities[reflect.ValueOf(Organizer{}).String()] = Organizer{}
	Entities[reflect.ValueOf(Event{}).String()] = Event{}
	Entities[reflect.ValueOf(User{}).String()] = User{}
}

type Organizer struct {
	ID uint64 `id:"ID" gorm:"primaryKey"`
	*User
	Name      string
	CreatedAt time.Time
}

type User struct {
	ID       uint64 `id:"ID" gorm:"primaryKey"`
	Username string
	Password string
	Role     string
}

type Event struct {
	ID   uint64 `id:"ID" gorm:"primaryKey"`
	Name string
	*Organizer
	Date time.Time
}
