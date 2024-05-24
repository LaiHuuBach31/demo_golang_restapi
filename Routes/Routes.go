package Routes

import (
	Controller "demo_api/Controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/cashier/:cashierId/login", Controller.Login)
	app.Post("/cashier/:cashierId/logout", Controller.Logout)
	app.Post("/cashier/:cashierId/passcode", Controller.Passcode)

	// cashier routes
	app.Get("/cashiers", Controller.GetCashierList)
	app.Get("/cashiers/:cashierId", Controller.GetCashierDetail)
	app.Post("/cashiers", Controller.CreateCashier)
	app.Put("/cashiers/:cashierId", Controller.UpdateCashier)
	app.Delete("/cashiers/:cashierId", Controller.DeleteCashier)
}
