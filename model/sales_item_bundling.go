package model

import "time"

type SalesItemBundling struct {
	ID                 int64     `json:"id"`
	OrderItemID        int64     `json:"order_item_id"`
	Sku                string    `json:"sku"`
	Name               string    `json:"name"`
	Price              int64     `json:"price"`
	Fixed              string    `json:"fixed"`
	Qty                int64     `json:"qty"`
	DiscountAmount     int64     `json:"discount_amount"`
	DiscountAmountJson string    `json:"discount_amount_json"`
	SiteCode           string    `json:"site_code"`
	SiteCodeStock      string    `json:"site_code_stock"`
	MaterialGroup      string    `json:"material_group"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
	PromotionID        string    `json:"promotion_id"`
}
