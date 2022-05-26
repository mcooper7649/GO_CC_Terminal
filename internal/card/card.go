package card

import "github.com/stripe/stripe-go/v72"

type Card struct {
	Secret   string
	Key      string
	Currency string
}

type Transaction struct {
	TransactionStatusId int
	Amount              int
	Currency            string
	LastFour            int
	BankReturnCode      string
}

func (c *Card) Charge(currency string, amount int)(*stripe.PaymentIntent, string, error){
	return c.createPaymentIntent(currency, amount)
}

func (c *Card) createPaymentIntent(currency string, amount int) (*stripe.PaymentIntent, string, error) {
	stripe.Key = c.Secret

	// create a payment intent
	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(int64(amount)),
		Currency: stripe.String(currency),
	}

	//params.AddMetadata("key", "value")

	pi, err := paymentIntent.New(params)
	if err != nil {
		msg := ""
		if stripeErr, ok := err.(*stripe.Error); ok {
			msg = cardErrorMessage(stripeErr.Code)
		}
		return nil, msg, err
	}
	return pi, "", nil
}


func cardErrorMessage(code stripe.ErrorCode) string {
	var msg = ""
	switch code {
	case: stripe.ErrorCodeCardDeclined:
		msg = "Your card was declined"
	case: stripe.ErrorCodeCardExpiredCard:
		msg: "Your card was expired"
	default: 
		msg = "Your card was declined"
	}

	return msg;
}