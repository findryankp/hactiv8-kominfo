package models

import "time"

type Item struct {
	ID          uint      `json:"id" gorm:"primary_key"`
	ItemCode    string    `json:"item_code"`
	Description string    `json:"description"`
	Quantity    int       `json:"quantity"`
	OrderId     uint      `json:"order_id"`
	Order       Order     `gorm:"foreignKey:OrderId"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type FormItem struct {
	ItemCode    string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}
