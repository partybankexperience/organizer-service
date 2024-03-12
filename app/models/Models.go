package models

import (
	"reflect"
	"time"
)

var Entities = make(map[string]any, 100)

func init() {
	Entities[reflect.ValueOf(Organizer{}).String()] = Organizer{}
	Entities[reflect.ValueOf(Event{}).String()] = Event{}
}

type Organizer struct {
	ID        uint64 `id:"ID" gorm:"primaryKey"`
	Name      string
	CreatedAt time.Time
}

type Event struct {
	ID   uint64 `id:"ID" gorm:"primaryKey"`
	Name string
	Date time.Time
}
