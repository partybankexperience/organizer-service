package test

import (
	"github.com/djfemz/organizer-service/partybank-app/integrations"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var paymentService integrations.PaymentService

func initPaymentService() {
	paymentService = integrations.NewPaymentService(os.Getenv("PAYMENT_SERVICE_CLIENT_ID"),
		os.Getenv("PAYMENT_SERVICE_CLIENT_SECRET"))
}
func TestPaymentServiceDeleteEvent(t *testing.T) {
	initPaymentService()
	eventReference := "evt-ODFhNDJhZDAtMmUwNS00YjFiLTk5YTctYjBmYThiMGFmZGFj"
	err := paymentService.DeleteEventBy(eventReference)
	assert.Nil(t, err)
}
