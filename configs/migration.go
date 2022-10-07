package configs

import (
	"orders/models"
)

func InitMigrate() {
	DB.AutoMigrate(
		&models.Item{},
		&models.Order{},
	)
}
