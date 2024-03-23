package dtos

type AuthRequest struct {
	CreateOrganizerRequest
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
	Sender     Sender      `json:"sender"`
	Recipients []Recipient `json:"to"`
	Subject    string      `json:"subject"`
	Content    string      `json:"htmlContent"`
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
