package payment

// PaymentGateway is a port for charging a membership purchase.
// The default implementation always succeeds. In tests, provide a stub that fails when you need to.
type PaymentGateway interface {
	// Charge attempts to charge amountCents for a purchase.
	// Returns PaymentFailedError when the charge is declined or the gateway errors.
	Charge(amountCents int, memberName string) (PaymentResult, error)
}

// PaymentResult is the outcome of a successful charge.
type PaymentResult struct {
	TransactionID string
}

// PaymentFailedError is returned when a charge is declined or the gateway errors.
type PaymentFailedError struct {
	Message string
}

func (e *PaymentFailedError) Error() string {
	return e.Message
}
