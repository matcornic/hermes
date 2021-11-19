package model

import (
	"time"
)

type SalesOrderItem struct {
	ID                             int64                `json:"id"`
	OrderID                        int64                `json:"order_id"`
	ProductID                      int64                `json:"product_id"`
	MerchantID                     int64                `json:"merchant_id"`
	ParentID                       int64                `json:"parent_id"`
	Sku                            string               `json:"sku"`
	UrlKey                         string               `json:"url_key"`
	Name                           string               `json:"name"`
	ProductType                    string               `json:"product_type"`
	Weight                         float64              `json:"weight"`
	Qty                            int                  `json:"qty"`
	Price                          int64                `json:"price"`
	PriceFormated                  string               `json:"price_formated"`
	DiscountInternalAmount         int64                `json:"discount_internal_amount"`
	DiscountInternalAmountFormated string               `json:"discount_internal_amount_formated"`
	DiscountInternalJson           string               `json:"discount_internal_json"`
	OrderItemBundlingID            int64                `json:"order_item_bundling_id"`
	ImageUrl                       string               `json:"image_url"`
	OriginalPrice                  int64                `json:"original_price"`
	OriginalPriceFormated          string               `json:"original_price_formated"`
	SiteCode                       string               `json:"site_code"`
	SiteCodeStock                  string               `json:"site_code_stock"`
	CreatedAt                      time.Time            `json:"created_at"`
	UpdatedAt                      time.Time            `json:"updated_at"`
	Bundling                       *[]SalesItemBundling `json:"bundling" gorm:"foreignKey:OrderItemID"`
	Options                        *[]SalesItemOption   `json:"options" gorm:"foreignKey:OrderItemID"`
}
