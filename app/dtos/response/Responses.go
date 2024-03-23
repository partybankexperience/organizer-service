package dtos

const (
	MAIL_SENDING_SUCCESS_MESSAGE = "mail sent successfully"
	USER_CREATED_SUCCESSFULLY    = "user created successfully"
)

type RaveResponse[T any] struct {
	Data T `json:"data"`
}

type LoginResponse struct {
	Message  string `json:"message,omitempty"`
	Username string `json:"username,omitempty"`
}

type CreateOrganizerResponse struct {
	Message  string `json:"message,omitempty"`
	Username string `json:"username,omitempty"`
}
