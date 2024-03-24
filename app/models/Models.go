package models

import (
	"github.com/djfemz/rave/app/security/otp"
	"reflect"
	"time"
)

var Entities = make(map[string]any, 100)

const (
	ADMIN       = "ADMIN"
	ORGANIZER   = "ORGANIZER"
	EVENt_STAFF = "EVENT_STAFF"
)

const (
	NOT_STARTED = "NOT_STARTED"
	ONGOING     = "ON_GOING"
	ENDED       = "ENDED"
)

// Used to register entities
func init() {
	Entities[reflect.ValueOf(Organizer{}).String()] = Organizer{}
	Entities[reflect.ValueOf(Event{}).String()] = Event{}
	Entities[reflect.ValueOf(EventStaff{}).String()] = EventStaff{}
}

type Organizer struct {
	ID uint64 `id:"ID" gorm:"primaryKey"`
	*User
	Name      string
	CreatedAt time.Time
	Otp       *otp.OneTimePassword `gorm:"embedded;embeddedPrefix:otp"`
}

type User struct {
	ID       uint64 `id:"ID" gorm:"primaryKey" json:"id"`
	Username string `json:"username"`
	Password string `json:"password" "omitempty"`
	Role     string `json:"role"`
}

type Event struct {
	ID                 uint64 `id:"ID" gorm:"primaryKey" json:"id"`
	Name               string `json:"name"`
	*Organizer         `json:"_organizer"`
	Location           string `json:"location"`
	Date               string `json:"date"`
	Time               string
	ContactInformation string
	Description        string
	Status             string `json:"status"`
}

type EventStaff struct {
	User  *User
	Event *Event
}
