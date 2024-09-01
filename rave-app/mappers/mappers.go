package mappers

import (
	response "github.com/djfemz/rave/rave-app/dtos/response"
	"github.com/djfemz/rave/rave-app/models"
	"github.com/djfemz/rave/rave-app/services"
	"log"
)

func MapSeriesCollectionToSeriesResponseCollection(series []*models.Series) []*response.SeriesResponse {
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
	log.Println("events: ", events)
	responses := make([]*response.EventResponse, 0)
	for _, event := range events {
		eventResponse := &response.EventResponse{
			ID:                 event.ID,
			Name:               event.Name,
			Status:             event.Status,
			Date:               event.EventDate,
			Time:               event.StartTime,
			Description:        event.Description,
			Location:           event.Location,
			ContactInformation: event.ContactInformation,
			MapEmbeddedUrl:     event.MapEmbeddedUrl,
			MapUrl:             event.MapUrl,
			ImageUrl:           event.ImageUrl,
		}
		if event.Location != nil {
			eventResponse.Location = event.Location
		}
		responses = append(responses, eventResponse)
	}
	return responses
}

func MapEventToEventResponse(event *models.Event, seriesService services.SeriesService) *response.EventResponse {
	series, err := seriesService.GetById(event.SeriesID)
	if err != nil {
		return nil
	}
	eventResponse := &response.EventResponse{
		ID:      event.ID,
		Message: "event created successfully",
		Name:    event.Name,

		Date:               event.EventDate,
		Time:               event.StartTime,
		ContactInformation: event.ContactInformation,
		Description:        event.Description,
		Status:             event.Status,
		SeriesID:           series.ID,
		Venue:              event.Venue,
		MapUrl:             event.MapUrl,
		MapEmbeddedUrl:     event.MapEmbeddedUrl,
		AttendeeTerm:       event.AttendeeTerm,
		EventTheme:         event.EventTheme,
		ImageUrl:           event.ImageUrl,
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
