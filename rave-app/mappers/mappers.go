package mappers

import (
	response "github.com/djfemz/rave/rave-app/dtos/response"
	"github.com/djfemz/rave/rave-app/models"
	"github.com/djfemz/rave/rave-app/utils"
	"log"
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
			Events:      MapEventsToEventResponses(userSeries.Events),
		}
		seriesResponses = append(seriesResponses, seriesResponse)
	}
	return seriesResponses
}

func MapEventsToEventResponses(events []*models.Event) []*response.EventResponse {
	responses := make([]*response.EventResponse, 0)
	for _, event := range events {
		ticketResponses := GetTicketsFrom(event)
		log.Println("tickets: ", ticketResponses)
		eventResponse := MapEventToEventResponse(event)

		eventResponse.Tickets = ticketResponses
		responses = append(responses, eventResponse)
	}
	return responses
}

func GetTicketsFrom(event *models.Event) []*response.TicketResponse {
	log.Println("tickkets: ", event.Tickets)
	ticketResponses := make([]*response.TicketResponse, 0)
	for _, ticket := range event.Tickets {
		ticketResponse := &response.TicketResponse{
			Type:                         ticket.Type,
			Name:                         ticket.Name,
			Capacity:                     ticket.Capacity,
			NumberAvailable:              ticket.NumberAvailable,
			Price:                        ticket.Price,
			PurchaseLimit:                ticket.PurchaseLimit,
			DiscountType:                 ticket.DiscountType,
			Percentage:                   ticket.Percentage,
			DiscountAmount:               ticket.DiscountAmount,
			DiscountCode:                 ticket.DiscountCode,
			AvailableDiscountedTickets:   ticket.AvailableDiscountedTickets,
			IsTransferPaymentFeesToGuest: ticket.IsTransferPaymentFeesToGuest,
			AdditionalInformationFields:  ticket.AdditionalInformationFields,
			SaleEndDate:                  ticket.SaleEndDate,
			SalesEndTime:                 ticket.SalesEndTime,
			Stock:                        ticket.Stock,
		}
		isTicketSaleEnded := utils.IsTicketSaleEndedFor(ticket)
		ticketResponse.IsTicketSaleEnded = isTicketSaleEnded
		ticketResponses = append(ticketResponses, ticketResponse)
	}
	return ticketResponses
}

func MapEventToEventResponse(event *models.Event) *response.EventResponse {
	tickets := GetTicketsFrom(event)
	eventResponse := &response.EventResponse{
		ID:                 event.ID,
		Message:            "event created successfully",
		Name:               event.Name,
		Date:               event.EventDate,
		Time:               event.StartTime,
		ContactInformation: event.ContactInformation,
		Description:        event.Description,
		Status:             event.Status,
		SeriesID:           event.SeriesID,
		Venue:              event.Venue,
		MapUrl:             event.MapUrl,
		MapEmbeddedUrl:     event.MapEmbeddedUrl,
		AttendeeTerm:       event.AttendeeTerm,
		EventTheme:         event.EventTheme,
		ImageUrl:           event.ImageUrl,
		Reference:          event.Reference,
		CreatedBy:          event.CreatedBy,
		PublicationState:   event.PublicationState,
		Tickets:            tickets,
	}

	if event.Location != nil {
		eventResponse.Location = &models.Location{
			State:   event.Location.State,
			Country: event.Location.Country,
			City:    event.Location.City,
		}
	}
	return eventResponse
}
