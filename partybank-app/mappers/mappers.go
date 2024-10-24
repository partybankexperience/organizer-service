package mappers

import (
	"log"
	"time"

	dtos "github.com/djfemz/organizer-service/partybank-app/dtos/request"
	response "github.com/djfemz/organizer-service/partybank-app/dtos/response"
	"github.com/djfemz/organizer-service/partybank-app/models"
)

func MapSeriesCollectionToSeriesResponseCollection(series []*models.Series, organizer *models.Organizer) []*response.SeriesResponse {
	seriesResponses := make([]*response.SeriesResponse, 0)
	for _, userSeries := range series {
		seriesResponse := &response.SeriesResponse{
			ID:          userSeries.ID,
			OrganizerID: userSeries.OrganizerID,
			Name:        userSeries.Name,
			Description: userSeries.Description,
			ImageUrl:    userSeries.ImageUrl,
			Events:      MapEventsToEventResponses(userSeries.Events, series[0]),
		}
		seriesResponses = append(seriesResponses, seriesResponse)
	}
	return seriesResponses
}

func MapEventsToEventResponses(events []*models.Event, series *models.Series) []*response.EventResponse {
	responses := make([]*response.EventResponse, 0)
	for _, event := range events {
		ticketResponses := GetTicketsFrom(event)
		log.Println("tickets: ", ticketResponses)
		eventResponse := MapEventToEventResponse(event)
		eventResponse.SeriesName = series.Name
		eventResponse.Tickets = ticketResponses
		eventResponse.SeriesLogo = series.Logo
		responses = append(responses, eventResponse)
	}
	return responses
}

func GetTicketsFrom(event *models.Event) []*response.TicketResponse {
	log.Println("tickkets: ", event.Tickets)
	ticketResponses := make([]*response.TicketResponse, 0)
	for _, ticket := range event.Tickets {
		ticketResponse := MapTicketToTicketResponse(ticket)
		//isTicketSaleEnded := IsTicketSaleEndedFor(ticket)
		//ticketResponse.IsTicketSaleEnded = isTicketSaleEnded
		ticketResponses = append(ticketResponses, ticketResponse)
	}
	return ticketResponses
}

func MapEventToEventResponse(event *models.Event) *response.EventResponse {
	tickets := GetTicketsFrom(event)
	eventTime := buildEventTimeForEventResponse(event)
	eventResponse := &response.EventResponse{
		ID:                 event.ID,
		Message:            "event created successfully",
		Name:               event.Name,
		Date:               event.EventDate,
		Time:               eventTime,
		ContactInformation: event.ContactInformation,
		Description:        event.Description,
		Status:             event.Status,
		SeriesID:           event.SeriesID,
		Venue:              event.Venue,
		AttendeeTerm:       event.AttendeeTerm,
		EventTheme:         event.EventTheme,
		ImageUrl:           event.ImageUrl,
		Reference:          event.Reference,
		CreatedBy:          event.CreatedBy,
		PublicationState:   event.PublicationState,
		Tickets:            tickets,
	}

	if event.Location != nil {
		eventResponse.Location = event.Location
	}
	return eventResponse
}

func buildEventTimeForEventResponse(event *models.Event) string {
	var eventTime string
	if event.StartTime == "" && event.EndTime != "" {
		eventTime = event.EndTime
	} else if event.StartTime != "" && event.EndTime == "" {
		eventTime = event.StartTime
	} else {
		eventTime = event.StartTime + " - " + event.EndTime
	}
	return eventTime
}

func MapTicketToTicketResponse(ticket *models.Ticket) *response.TicketResponse {
	ticketResponse := &response.TicketResponse{
		Id:                           ticket.ID,
		Type:                         ticket.Type,
		Name:                         ticket.Name,
		Capacity:                     ticket.Capacity,
		Stock:                        ticket.Stock,
		Reference:                    ticket.Reference,
		NumberAvailable:              ticket.NumberAvailable,
		Price:                        ticket.Price,
		PurchaseLimit:                ticket.PurchaseLimit,
		DiscountType:                 ticket.DiscountType,
		AvailableDiscountedTickets:   ticket.AvailableDiscountedTickets,
		IsTransferPaymentFeesToGuest: ticket.IsTransferPaymentFeesToGuest,
		AdditionalInformationFields:  ticket.AdditionalInformationFields,
		Colour:                       ticket.Colour,
		TicketPerks:                  ticket.TicketPerks,
		IsTicketSaleEnded:            ticket.IsSoldOutTicket,
	}
	if ticket.ActivePeriod != nil {
		ticketResponse.SaleEndDate = ticket.ActivePeriod.EndDate
		ticketResponse.SalesEndTime = ticket.ActivePeriod.EndTime
		ticketResponse.SalesStartTime = ticket.ActivePeriod.StartTime
		ticketResponse.SalesStartDate = ticket.ActivePeriod.StartDate
	}
	return ticketResponse
}

func MapCreateAttendeeRequestToAttendee(createAttendeeRequest *dtos.CreateAttendeeRequest) *models.Attendee {

	return &models.Attendee{
		FirstName: createAttendeeRequest.FullName,
		User: &models.User{
			Username: createAttendeeRequest.Username,
			Role:     models.ATTENDEE,
		},
	}
}

func MapAttendeeToAttendeeResponse(attendee *models.Attendee) *response.AttendeeResponse {
	return &response.AttendeeResponse{
		Username:    attendee.Username,
		Message:     "User registered successfully",
		FirstName:   attendee.FirstName,
		LastName:    attendee.LastName,
		PhoneNumber: attendee.PhoneNumber,
	}
}

func MapEditTicketRequestToTicket(editTicketRequest *dtos.EditTicketRequest, ticket *models.Ticket) *models.Ticket {
	ticket.Colour = editTicketRequest.Colour
	ticket.Name = editTicketRequest.Name
	ticket.Stock = editTicketRequest.Stock
	ticket.Price = editTicketRequest.Price
	ticket.TicketPerks = editTicketRequest.TicketPerks
	ticket.IsTransferPaymentFeesToGuest = editTicketRequest.IsTransferPaymentFeesToGuest
	if ticket.ActivePeriod != nil {
		ticket.ActivePeriod.StartTime = editTicketRequest.SalesStartTime
		ticket.ActivePeriod.EndTime = editTicketRequest.SalesEndTime
		ticket.ActivePeriod.StartDate = editTicketRequest.SalesStartDate
		ticket.ActivePeriod.EndDate = editTicketRequest.SaleEndDate
	}
	ticket.PurchaseLimit = editTicketRequest.PurchaseLimit
	return ticket
}

func IsTicketSaleEndedFor(ticket *models.Ticket) bool {
	if ticket.ActivePeriod == nil {
		return false
	}
	ticketEndTime := ticket.ActivePeriod.EndDate + " " + ticket.ActivePeriod.EndTime
	endTime, err := time.Parse("2006-01-02 15:04:05", ticketEndTime)
	if err != nil {
		log.Println("err: ", err)
		return false
	}
	log.Println("true: ", endTime)
	return time.Now().After(endTime)
}
