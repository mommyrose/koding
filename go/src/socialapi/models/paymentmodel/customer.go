package paymentmodel

import (
	"time"
)

type Customer struct {
	Id int64 `json:"id,string"`

	Username string `json:"username"`

	// Id of customer in 3rd payment provider like Stripe.
	ProviderCustomerId string `json:"providerPlanId"`

	// Name of provider. Enum:
	//    'stripe', 'paypal'
	Provider string `json:"provider"`

	// Timestamps.
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" `
}

func NewCustomer(username, providerId, provider string) *Customer {
	return &Customer{
		Username:           username,
		ProviderCustomerId: providerId,
		Provider:           provider,
	}
}

func (c *Customer) Create() error {
	return nil
}

// func (c *Customer) FindOrCreate() error {}
// func (c *Customer) Find() error         {}
