package main

import (
	"orders/configs"
	"orders/routes"
)

func main() {
	configs.InitEnv()
	configs.InitDB()
	configs.InitMigrate()

	r := routes.InitRoutes()
	r.Run()
}
