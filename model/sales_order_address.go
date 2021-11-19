package model

import "time"

type SalesOrderAddress struct {
	ID                int64     `json:"id"`
	OrderID           int64     `json:"order_id"`
	AddressID         int64     `json:"address_id"`
	CustomerID        int64     `json:"customer_id"`
	Latitude          float64   `json:"latitude"`
	Longitude         float64   `json:"longitude"`
	StreetMap         string    `json:"street_map"`
	Country           string    `json:"country"`
	CountryCode       string    `json:"country_code"`
	Region            string    `json:"region"`
	RegionCode        string    `json:"region_code"`
	City              string    `json:"city"`
	CityID            int64     `json:"city_id"`
	District          string    `json:"district"`
	DistrictID        int64     `json:"district_id"`
	SubDistrict       string    `json:"sub_district"`
	SubDistrictID     int64     `json:"sub_district_id"`
	Postcode          string    `json:"postcode"`
	Street            string    `json:"street"`
	Telephone         string    `json:"telephone"`
	IsBilling         bool      `json:"is_billing"`
	IsDefaultBilling  bool      `json:"is_default_billing"`
	IsShipping        bool      `json:"is_shipping"`
	IsDefaultShipping bool      `json:"is_default_shipping"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	Email             string    `json:"email"`
}
