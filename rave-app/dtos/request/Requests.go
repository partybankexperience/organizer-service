package dtos

import "github.com/djfemz/rave/rave-app/models"

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
	Type                         string                             `json:"ticket_type"`
	Name                         string                             `json:"name"`
	Stock                        uint64                             `json:"stock"`
	NumberAvailable              uint64                             `json:"number_in_stock"`
	Price                        float64                            `json:"price"`
	PurchaseLimit                uint64                             `json:"purchase_limit"`
	DiscountType                 string                             `json:"discount_type"`
	Percentage                   float64                            `json:"percentage"`
	DiscountPrice                float64                            `json:"discount_price"`
	DiscountCode                 string                             `json:"discount_code"`
	AvailableDiscountedTickets   uint64                             `json:"available_discounted_tickets"`
	IsTransferPaymentFeesToGuest bool                               `json:"is_transfer_payment_fees_to_guest"`
	AdditionalInformationFields  models.AdditionalInformationFields `json:"additional_information_fields"`
	EventId                      uint64                             `json:"event_id" validate:"required"`
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

type NewTicketMessage struct {
	Type                         string                             `json:"ticket_type"`
	Name                         string                             `json:"name"`
	Stock                        uint64                             `json:"stock"`
	NumberAvailable              uint64                             `json:"number_in_stock"`
	Price                        float64                            `json:"price"`
	PurchaseLimit                uint64                             `json:"purchase_limit"`
	DiscountType                 string                             `json:"discount_type"`
	Percentage                   float64                            `json:"percentage"`
	DiscountPrice                float64                            `json:"discount_price"`
	DiscountCode                 string                             `json:"discount_code"`
	AvailableDiscountedTickets   uint64                             `json:"available_discounted_tickets"`
	IsTransferPaymentFeesToGuest bool                               `json:"is_transfer_payment_fees_to_guest"`
	AdditionalInformationFields  models.AdditionalInformationFields `json:"additional_information_fields,omitempty"`
	Message                      string                             `json:"message,omitempty"`
	EventName                    string                             `json:"event_name"`
	Organizer                    string                             `json:"organizer"`
	Location                     *models.Location                   `json:"location"`
	Date                         string                             `json:"date"`
	Time                         string                             `json:"time"`
	ContactInformation           string                             `json:"contact_information"`
	Description                  string                             `json:"description"`
	Status                       string                             `json:"status"`
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
