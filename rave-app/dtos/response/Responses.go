package dtos

import "github.com/djfemz/rave/rave-app/models"

const (
	MAIL_SENDING_SUCCESS_MESSAGE = "mail sent successfully"
	USER_CREATED_SUCCESSFULLY    = "user created successfully"
)

type RaveResponse[T any] struct {
	Data T `json:"data"`
}

type CreateDiscountResponse struct {
	Id uint64 `json:"id"`
}

type LoginResponse struct {
	Message  string `json:"message,omitempty"`
	Username string `json:"username,omitempty"`
}

type CreateOrganizerResponse struct {
	Message  string `json:"message,omitempty"`
	Username string `json:"username,omitempty"`
}

type CreateCalendarResponse struct {
	Message string
	ID      uint64
	Name    string
	Events  []*EventResponse
}

type CalendarResponse struct {
	ID     uint64
	Name   string
	Events []*EventResponse
}

type EventResponse struct {
	ID                 uint64 `json:"id"`
	Message            string `json:"message,omitempty"`
	Name               string `json:"name"`
	Calendar           string `json:"calendar"`
	Location           string `json:"location"`
	Date               string `json:"date"`
	Time               string `json:"time"`
	ContactInformation string `json:"contact_information"`
	Description        string `json:"description"`
	Status             string `json:"status"`
}

type TicketResponse struct {
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
}
