package model

import "time"

type SalesPayment struct {
	ID                       int64     `json:"id"`
	OrderID                  string    `json:"order_id"`
	StatusPayment            string    `json:"status_payment"`
	InvoiceNumber            string    `json:"invoice_number"`
	PaymentMethodCode        string    `json:"payment_method_code"`
	PaymentMethodDescription string    `json:"payemnt_method_description"`
	PaymentIssuer            string    `json:"payment_issuer"`
	Tenor                    int64     `json:"tenor"`
	PaymentGateway           string    `json:"payment_gateway"`
	VaNumber                 string    `json:"va_number"`
	AmountPaid               int64     `json:"amount_paid"`
	AmountRefunded           int64     `json:"amount_refunded"`
	PaymentID                int64     `json:"payment_id"`
	CreatedAt                time.Time `json:"created_at"`
	UpdatedAt                time.Time `json:"updated_at"`
	ExpiredAt                time.Time `json:"expired_at"`
}
