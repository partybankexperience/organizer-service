package dtos

import (
	"database/sql/driver"
	"errors"
	"strings"
)

const (
	LIMITED   = "LIMITED"
	UNLIMITED = "UNLIMITED"
)

type AuthRequest struct {
	CreateUserRequest
}

type CreateDiscountRequest struct {
	TicketId uint64  `json:"ticket_id" validate:"required"`
	Name     string  `json:"name"`
	Code     string  `json:"code"`
	Count    uint64  `json:"count"`
	Value    string  `json:"value"`
	Price    float64 `json:"price"`
}

type CreateEventRequest struct {
	Name               string `json:"name" validate:"required"`
	City               string `json:"city"`
	State              string `json:"state" validate:"required"`
	Country            string `json:"country" validate:"required"`
	Date               string `json:"date" validate:"required"`
	Time               string `json:"time" validate:"required"`
	ContactInformation string `json:"contact_information"`
	Description        string `json:"description"`
	SeriesId           uint64 `json:"series_id" validate:"required"`
	OrganizerId        uint64 `json:"organizer_id" validate:"required"`
	EventTheme         string `json:"event_theme"`
	MapUrl             string `json:"map_url"`
	MapEmbeddedUrl     string `json:"map_embedded_url"`
	AttendeeTerm       string `json:"attendee_term"`
	Venue              string `json:"venue" validate:"required"`
	ImageUrl           string `json:"image_url"`
}

type UpdateEventRequest struct {
	Name               string `json:"name"`
	Location           string `json:"location"`
	Date               string `json:"date"`
	Time               string `json:"time"`
	ContactInformation string `json:"contact_information"`
	Description        string `json:"description"`
	OrganizerId        uint64 `json:"organizer_id" validate:"required"`
	EventTheme         string `json:"event_theme"`
	MapUrl             string `json:"map_url"`
	MapEmbeddedUrl     string `json:"map_embedded_url"`
	AttendeeTerm       string `json:"attendee_term"`
	Venue              string `json:"venue" validate:"required"`
}

type CreateTicketRequest struct {
	Type          string  `json:"ticket_type"`
	Name          string  `json:"name"`
	Capacity      uint64  `json:"capacity"`
	Stock         string  `json:"stock"`
	Price         float64 `json:"price"`
	PurchaseLimit uint64  `json:"purchase_limit"`
	//DiscountType                 string                             `json:"discount_type"`
	//Percentage float64 `json:"percentage"`
	//DiscountPrice                float64                            `json:"discount_price"`
	//DiscountCode                 string                             `json:"discount_code"`
	//AvailableDiscountedTickets   uint64                             `json:"available_discounted_tickets"`
	IsTransferPaymentFeesToGuest bool `json:"is_transfer_payment_fees_to_guest"` //TODO: Default: false
	//AdditionalInformationFields  models.AdditionalInformationFields `json:"additional_information_fields"`
	EventId         uint64      `json:"event_id" validate:"required"`
	Colour          string      `json:"colour"`
	SaleEndDate     string      `json:"ticket_sale_end_date"`
	SalesEndTime    string      `json:"ticket_sales_end_time"`
	TicketPerks     TicketPerks `json:"ticket_perks"`
	PriceChangeDate string      `json:"price_change_date"`
	PriceChangeTime string      `json:"price_change_time"`
}

type TicketPerks []string

func (o *TicketPerks) Scan(src any) error {
	bytes, ok := src.(string)
	if !ok {
		return errors.New("src value cannot cast to []byte")
	}
	*o = strings.Split(bytes, ",")
	return nil
}
func (o TicketPerks) Value() (driver.Value, error) {
	if len(o) == 0 {
		return nil, nil
	}
	return strings.Join(o, ","), nil
}

type CreateUserRequest struct {
	Username string `json:"username" validate:"required,email"`
}

type AddEventStaffRequest struct {
	StaffEmails []string `json:"staff_emails"`
	EventId     uint64   `json:"event_id" validate:"required"`
}

type CreateEventStaffRequest struct {
	StaffEmails []string `json:"staff_emails"`
	EventId     uint64   `json:"event_id" validate:"required"`
}

type TicketType struct {
	Reference string  `json:"ticketTypeReference"`
	Reserved  uint64  `json:"reservedSeats"`
	MaxSeats  uint64  `json:"maxSeats"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	Colour    string  `json:"color"`
	Category  uint64  `json:"category"`
	Stock     string
}

//Ticket Type [Free, Paid]
//Ticket Name
//Ticket Price
//Ticket Capacity [Limited, Unlimited]
//Capacity/ Available Tickets [if limited]
//Ticket Purchase Limit [How many tickets can be bought at a time?]
//Ticket Perks [What are the benefits attached to ticket type?]

//"reference": "TCKT002",
//"category": "Single Ticket",
//"type": "Paid",
//"name": "Standard",
//"price": "5000",
//"color": "Yellow",
//"stock": "Limited",
//"capacity": 1000,
//"purchaseLimit": 5,
//"perks": "Free food",
//"salesEndDate": "2024-08-28",
//"salesEndTime": "17:00",
//"priceChangeDate": "2024-08-20",
//"priceChangeTime": "12:00"

type NewTicketMessage struct {
	Reference    string        `json:"eventReference"`
	Types        []*TicketType `json:"ticketTypes"`
	Name         string        `json:"eventName"`
	Venue        string        `json:"venue"`
	TimeFrame    string        `json:"timeFrame"`
	AttendeeTerm string        `json:"attendeeTerm"`
	Date         string        `json:"eventDate"`
}

type Sender struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Recipient struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type EmailNotificationRequest struct {
	Sender     Sender      `json:"sender"`
	Recipients []Recipient `json:"to"`
	Subject    string      `json:"subject"`
	Content    string      `json:"htmlContent"`
}

type CreateSeriesRequest struct {
	Name        string `json:"name"`
	OrganizerID uint64 `json:"organizer_id" validate:"required"`
	Description string `json:"description"`
	ImageUrl    string `json:"image_url"`
	SeriesLogo  string `json:"series_logo"`
}

func NewEmailNotificationRequest(recipient, content string) *EmailNotificationRequest {
	return &EmailNotificationRequest{
		Sender: Sender{
			Email: "noreply@email.com",
			Name:  "rave",
		},
		Recipients: []Recipient{
			{Email: recipient, Name: "Friend"},
		},
		Subject: "rave",
		Content: content,
	}
}
