package models

import (
	"github.com/djfemz/rave/rave-app/security/otp"
	"reflect"
	"time"
)

var Entities = make(map[string]any, 100)

const (
	ADMIN       = "ADMIN"
	ORGANIZER   = "ORGANIZER"
	EVENT_STAFF = "EVENT_STAFF"
)

const (
	NOT_STARTED = "NOT_STARTED"
	ONGOING     = "ON_GOING"
	ENDED       = "ENDED"
)

// Used to register entities
func init() {
	Entities[reflect.ValueOf(Event{}).String()] = Event{}
	Entities[reflect.ValueOf(Organizer{}).String()] = Organizer{}
	Entities[reflect.ValueOf(EventStaff{}).String()] = EventStaff{}
}

type Organizer struct {
	ID uint64 `id:"ID" gorm:"primaryKey"`
	*User
	Name      string
	CreatedAt time.Time
	Otp       *otp.OneTimePassword `gorm:"embedded;embeddedPrefix:otp"`
	EventId   uint64
	Events    []*Event
}

type User struct {
	ID       uint64 `id:"ID" gorm:"primaryKey" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type Event struct {
	ID                 uint64 `id:"EventId" gorm:"primaryKey" json:"id"`
	Name               string `json:"name"`
	Location           string `json:"location"`
	Date               string `json:"date"`
	Time               string
	ContactInformation string
	Description        string
	OrganizerID        uint64
	Status             string `json:"status"`
	EventStaffID       uint64
	EventStaff         []*EventStaff
}

type EventStaff struct {
	ID      uint64 `id:"ID" gorm:"primaryKey" json:"id"`
	*User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	EventID uint64
}
