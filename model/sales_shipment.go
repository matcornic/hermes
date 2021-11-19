package model

import "time"

type SalesShipment struct {
	ID                  int64     `json:"id"`
	OrderID             int64     `json:"order_id"`
	CustomerID          int64     `json:"customer_id"`
	SalesMerchantID		int64	  `json:"sales_merchant_id"`
	SalesMerchantCode   string	  `json:"sales_merchant_code"`
	AwbNumber           string    `json:"awb_number"`
	ShippingMethodCode  string    `json:"shipping_method_code"`
	ShippingMethodTitle string    `json:"shipping_method_title"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
}
