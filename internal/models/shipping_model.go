package models

type Shipping struct {
	BaseModel
	OrderID      uint    `json:"order_id" gorm:"column:order_id;not null"`
	ShippingCost float64 `json:"shipping_cost" gorm:"column:shipping_cost;not null"`
	ShippingType string  `json:"shipping_type" gorm:"column:shipping_type;not null"`
}
