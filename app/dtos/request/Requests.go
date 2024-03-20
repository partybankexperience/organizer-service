package dtos

type LoginRequest struct {
	Username string
	Password string
}

type CreateOrganizerRequest struct {
	Username string
	Password string
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
	Sender    Sender      `json:"sender"`
	Recipient []Recipient `json:"to"`
	Subject   string      `json:"subject"`
	Content   string      `json:"htmlContent"`
}
