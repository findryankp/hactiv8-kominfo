package models

import "time"

type Order struct {
	ID           uint      `json:"id" gorm:"primary_key"`
	CustomerName string    `json:"customer_name"`
	OrderedAt    string    `json:"ordered_at"`
	CreatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type FormOrder struct {
	CustomerName string     `json:"customerName"`
	OrderedAt    string     `json:"orderedAt"`
	Item         []FormItem `json:"items"`
}
