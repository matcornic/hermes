package model

import "time"

type SalesItemOption struct {
	ID          int64     `json:"id" gorm:"column:id;primaryKey"`
	Name        string    `json:"name" gorm:"column:name"`
	Value       string    `json:"value" gorm:"column:value"`
	OrderItemID int64     `json:"order_item_id" gorm:"column:order_item_id"`
	OrderID     int64     `json:"order_id" gorm:"column:order_id"`
	CreatedAt   time.Time `json:"created_at" gorm:"<-:create;column:created_at"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"<-:update;column:updated_at"`
}
