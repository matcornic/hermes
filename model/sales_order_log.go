package model

import "time"

type SalesOrderLog struct {
	ID          int64     `json:"id"`
	OrderID     int64     `json:"order_id"`
	OrderNumber string    `json:"order_number"`
	Status      string    `json:"status"`
	ActionName  string    `json:"action_name"`
	InfoText    string    `json:"info_text"`
	InfoJson    string    `json:"info_json"`
	CreatedAt   time.Time `json:"created_at"`
}
