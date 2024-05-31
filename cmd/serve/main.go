package main

import (
	"git.strooware.nl/postform/common"
	"git.strooware.nl/postform/handlers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	if err := common.InitConfig(); err != nil {
		panic(err)
	}

	app := fiber.New(fiber.Config{
		DisableStartupMessage: common.Config.Debug,
		EnablePrintRoutes:     common.Config.Debug,
		PassLocalsToViews:     true,
		ServerHeader:          "",
	})

	app.Route("/f/:formId", func(r fiber.Router) {
		r.Use(handlers.FindFormMiddleware)
		r.Use(handlers.ValidateReferrerMiddleware)

		r.Post("/", handlers.PostFormHandler)
	})

	if err := app.Listen(":8080"); err != nil {
		panic(err)
	}
}
