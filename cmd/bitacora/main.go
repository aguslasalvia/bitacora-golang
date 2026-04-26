package main

import (
	"bitacora/internal/config"
	"bitacora/internal/routes"

	_ "bitacora/internal/config"
)

func main() {

	defer config.Database.Close()

	r := routes.SetupRoutes()

	r.Run("127.0.0.1:4040")
}
