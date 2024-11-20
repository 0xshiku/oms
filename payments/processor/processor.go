package processor

import pb "common/api"

type PaymentProcessor interface {
	CreatePaymentLink(*pb.Order) (string, error)
}
