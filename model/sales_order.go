package model

import "time"

type SalesOrder struct {
	ID                             int64                 `json:"id"`
	IsOrdered                      int                   `json:"is_ordered"`
	IsInvoiced                     int                   `json:"is_invoiced"`
	IsShipment                     int                   `json:"is_shipment"`
	IsCompleted                    int                   `json:"is_completed"`
	IsCanceled                     int                   `json:"is_canceled"`
	IsRefunded                     int                   `json:"is_refunded"`
	HoldBeforeStatus               string                `json:"hold_before_status"`
	Source                         string                `json:"source"`
	IsVirtual                      int                   `json:"is_virtual"`
	CustomerID                     int64                 `json:"customer_id"`
	CustomerGroupID                int                   `json:"customer_group_id"`
	QuoteID                        int64                 `json:"quote_id"`
	OrderNumber                    string                `json:"order_number"`
	Status                         string                `json:"status"`
	StatusLabel                    string                `json:"status_label"`
	CouponCode                     string                `json:"coupon_code"`
	ShippingMethod                 string                `json:"shipping_method"`
	ShippingDescription            string                `json:"shipping_description"`
	CustomerEmail                  string                `json:"customer_email"`
	CustomerFirstname              string                `json:"customer_firstname"`
	CustomerLastname               string                `json:"customer_lastname"`
	CustomerDob                    time.Time             `json:"customer_dob"`
	CustomerGender                 int                   `json:"customer_gender"`
	CustomerIsGuest                int                   `json:"customer_is_guest"`
	CustomerIp                     string                `json:"customer_ip"`
	TotalDiscount                  int64                 `json:"total_discount"`
	TotalDiscountFormated          string                `json:"total_discount_formated"`
	DiscountGiftcardAmount         int64                 `json:"discount_giftcard_amount"`
	DiscountGiftcardAmountFormated string                `json:"discount_giftcard_amount_formated"`
	DiscountCrmAmount              int64                 `json:"discount_crm_amount"`
	DiscountCrmAmountFormated      string                `json:"discount_crm_amount_formated"`
	DiscountPointAmount            int64                 `json:"discount_point_amount" gorm:"column:discount_point_amount"`
	DiscountPointRedeemtionid      string                `json:"discount_point_redeemtionid"`
	DiscountInternalAmount         int64                 `json:"discount_internal_amount"`
	DiscountAmount                 int64                 `json:"discount_amount"`
	DiscountAmountFormated         string                `json:"discount_amount_formated"`
	O2oBuCode                      string                `json:"o2o_bu_code"`
	O2oAddress                     string                `json:"o2o_address"`
	ISO2O                          int                   `json:"is_o2o" gorm:"column:is_o2o"`
	O2oNik                         string                `json:"o2o_nik"`
	O2oPickupDate                  time.Time             `json:"o2o_pickup_date"`
	O2oPickupCode                  string                `json:"o2o_pickup_code"`
	O2oPickupStore                 string                `json:"o2o_pickup_store"`
	O2oBookingNumber               string                `json:"o2o_booking_number"`
	O2oPayatstore                  string                `json:"o2o_payatstore"`
	O2oIdCard                      string                `json:"o2o_id_card"`
	PosNumber                      string                `json:"pos_number"`
	EraExpressOriginStore          string                `json:"era_express_origin_store"`
	VoucherCode                    string                `json:"voucher_code"`
	VoucherDescription             string                `json:"voucher_description"`
	VoucherAmount                  int64                 `json:"voucher_amount"`
	VoucherAmountFormated          string                `json:"voucher_amount_formated"`
	CashbackPaymentAmount          int64                 `json:"cashback_payment_amount"`
	CashbackPaymentAmountFormated  string                `json:"cashback_payment_amount_formated"`
	CashbackPaymentDetail          string                `json:"cashback_payment_detail"`
	CashbackPaymentCode            string                `json:"cashback_payment_code"`
	MultiWarehouseOrigin           string                `json:"multi_warehouse_origin"`
	MultiWarehouseCode             string                `json:"multi_warehouse_code"`
	TradeinAmount                  int64                 `json:"tradein_amount"`
	TradeinAmountFormated          string                `json:"tradein_amount_formated"`
	GiftCardAmount                 int64                 `json:"gift_card_amount"`
	GiftCardAmountFormated         string                `json:"gift_card_amount_formated"`
	ShippingAmountTotal            int64                 `json:"shipping_amount_total"`
	ShippingAmountTotalFormated    string                `json:"shipping_amount_total_formated"`
	Subtotal                       int64                 `json:"subtotal"`
	SubtotalFormated               string                `json:"subtotal_formated"`
	GrandTotal                     int64                 `json:"grand_total"`
	GrandTotalFormated             string                `json:"grand_total_formated"`
	CreatedAt                      time.Time             `json:"created_at"`
	UpdatedAt                      time.Time             `json:"updated_at"`
	Items                          *[]SalesOrderItem     `json:"items" gorm:"foreignKey:OrderID;"`
	ItemGroups                     []SalesOrderItemGroup `json:"item_groups" gorm:"foreignKey:OrderID;"`
	Address                        *[]SalesOrderAddress  `json:"address" gorm:"foreignKey:OrderID"`
	Merchant                       *[]SalesMerchant      `json:"merchant" gorm:"foreignKey:OrderID"`
	Shipment                       *[]SalesShipment      `json:"shipment" gorm:"foreignKey:OrderID"`
	Payment                        *[]SalesPayment       `json:"payment" gorm:"foreignKey:OrderID;references:OrderNumber"`
	OrderExpired                   time.Time             `json:"order_expired"`
	DiscountCrmOtpcode             string                `json:"discount_crm_otpcode"`
	DiscountCrmVouchercode         string                `json:"discount_crm_vouchercode"`
}

type SalesOrderItemGroup struct {
	ID                             int64            `json:"id"`
	OrderID                        int64            `json:"order_id"`
	MerchantCode                   string           `json:"merchant_code"`
	MerchantShippingTitle          string           `json:"merchant_shipping_title"`
	MerchantShippingCode           string           `json:"merchant_shipping_code"`
	MerchantShippingAmount         int64            `json:"merchant_shipping_amount"`
	MerchantShippingAmountFormated string           `json:"merchant_shipping_amount_formated"`
	Items                          []SalesOrderItem `json:"items"  gorm:"foreignKey:OrderID"`
	CreatedAt                      time.Time        `json:"created_at"`
	UpdatedAt                      time.Time        `json:"updated_at"`
}
