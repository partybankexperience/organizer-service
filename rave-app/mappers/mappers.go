package mappers

import (
	response "github.com/djfemz/rave/rave-app/dtos/response"
	"github.com/djfemz/rave/rave-app/models"
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
		}
		responses = append(responses, eventResponse)
	}
	return responses
}
