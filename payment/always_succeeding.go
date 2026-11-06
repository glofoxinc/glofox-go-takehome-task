package payment

import "github.com/google/uuid"

// AlwaysSucceedingPaymentGateway always succeeds.
type AlwaysSucceedingPaymentGateway struct{}

// Charge always returns a successful PaymentResult.
func (AlwaysSucceedingPaymentGateway) Charge(amountCents int, memberName string) (PaymentResult, error) {
	return PaymentResult{TransactionID: "txn-" + uuid.New().String()}, nil
}
