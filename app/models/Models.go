package models

import (
	"github.com/djfemz/rave/app/security/otp"
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
}

type Organizer struct {
	ID uint64 `id:"ID" gorm:"primaryKey"`
	*User
	Name      string
	CreatedAt time.Time
	Otp       *otp.Otp `gorm:"embedded;embeddedPrefix:otp"`
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
