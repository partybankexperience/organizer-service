package dtos

const (
	MAIL_SENDING_SUCCESS_MESSAGE = "mail sent successfully"
)

type RaveResponse[T any] struct {
	Data T `json:"data"`
}

type LoginResponse struct {
}

type CreateOrganizerResponse struct {
}
