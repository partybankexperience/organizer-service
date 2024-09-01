package models

import (
	"gorm.io/gorm"
	"reflect"
	"time"

	"github.com/djfemz/rave/rave-app/security/otp"
)

var Entities = make(map[string]any, 100)

const (
	ADMIN       = "ADMIN"
	ORGANIZER   = "ORGANIZER"
	EVENT_STAFF = "EVENT_STAFF"
)

const (
	PAST     = "ON_GOING"
	UPCOMING = "ENDED"
)

const (
	ACTIVE    = "ACTIVE"
	SUSPENDED = "SUSPENDED"
	IN_ACTIVE = "IN_ACTIVE"
)

// Used to register entities
func init() {
	Entities[reflect.ValueOf(Event{}).String()] = Event{}
	Entities[reflect.ValueOf(Organizer{}).String()] = Organizer{}
	Entities[reflect.ValueOf(EventStaff{}).String()] = EventStaff{}
	Entities[reflect.ValueOf(Ticket{}).String()] = Ticket{}
	Entities[reflect.ValueOf(Series{}).String()] = Series{}
}

type Organizer struct {
	ID uint64 `id:"ID" gorm:"primaryKey"`
	*User
	Name      string
	CreatedAt time.Time
	Otp       *otp.OneTimePassword `gorm:"embedded;embeddedPrefix:otp"`
	EventId   uint64
	Series    []*Series
}

type User struct {
	ID       uint64 `id:"ID" gorm:"primaryKey" json:"id"`
	Username string `json:"username" gorm:"unique"`
	Role     string `json:"role"`
}

type AdditionalInformationFields []string

type Ticket struct {
	ID                           uint64 `gorm:"primaryKey"`
	Type                         string
	Name                         string                      `json:"name"`
	Stock                        uint64                      `json:"stock"`
	NumberAvailable              uint64                      `json:"number_available"`
	Price                        float64                     `json:"price"`
	PurchaseLimit                uint64                      `json:"purchase_limit"`
	DiscountType                 string                      `json:"discount_type"`
	Percentage                   float64                     `json:"percentage"`
	DiscountPrice                float64                     `json:"discount_price"`
	DiscountCode                 string                      `json:"discount_code"`
	AvailableDiscountedTickets   uint64                      `json:"available_discounted_tickets"`
	AdditionalInformationFields  AdditionalInformationFields `gorm:"type:VARCHAR(255)" json:"additional_information_fields,omitempty"`
	IsTransferPaymentFeesToGuest bool
	EventId                      uint64
}

type Event struct {
	ID                 uint64    `id:"EventId" gorm:"primaryKey" json:"id"`
	Name               string    `json:"name"`
	Location           *Location `json:"location" gorm:"embedded"`
	EventDate          string    `json:"date"`
	StartTime          string    `json:"event_start"`
	EndTime            string    `json:"event_end"`
	ContactInformation string    `json:"contact_information"`
	ImageUrl           string    `json:"image_url"`
	Description        string    `json:"description"`
	SeriesID           uint64    `json:"series_id"`
	Status             string    `json:"status"`
	EventStaffID       uint64    `json:"event_staff_id"`
	TicketID           uint64    `json:"ticket_id"`
	EventTheme         string    `json:"event_theme"`
	MapUrl             string    `json:"map_url"`
	MapEmbeddedUrl     string    `json:"map_embedded_url"`
	AttendeeTerm       string    `json:"attendee_term"`
	Venue              string    `json:"venue"`
	Tickets            []*Ticket
	EventStaff         []*EventStaff
}

type Location struct {
	State   string `json:"state"`
	Country string `json:"country"`
	City    string `json:"city"`
}

type Series struct {
	ID   uint64 `id:"seriesId" gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
	gorm.Model
	Events      []*Event `json:"events"`
	OrganizerID uint64   `json:"organizer_id"`
	ImageUrl    string   `json:"image_url"`
	Description string   `json:"description"`
}

type EventStaff struct {
	ID      uint64 `id:"ID" gorm:"primaryKey" json:"id"`
	*User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user,omitempty"`
	EventID uint64
}

type Discount struct {
	ID uint64 `id:"ID" gorm:"primaryKey" json:"id"`
	*Ticket
	Name  string
	Code  string
	Count uint64
	Value string
	Price float64
}
