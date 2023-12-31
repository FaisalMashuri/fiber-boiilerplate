package router

import (
	"github.com/FaisalMashuri/my-wallet/config"
	"github.com/FaisalMashuri/my-wallet/internal/domain/user"
	"github.com/Saucon/errcntrct"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"log"
)

type RouteParams struct {
	UserController user.UserController
}

type router struct {
	RouteParams RouteParams
	Log         logrus.Logger
}

func NewRouter(params RouteParams) router {
	return router{RouteParams: params}
}

func (r *router) SetupRoute(app *fiber.App) {
	if err := errcntrct.InitContract(config.AppConfig.ErrorContract.JSONPathFile); err != nil {
		//logger.Fatal(err, "main : init contract", nil)
		log.Fatal("main : init contract " + err.Error())
	}

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).JSON("HALO")
	})

	// Define routes with auth
	v1 := app.Group("/api/v1")

	v1.Route("/auth", func(router fiber.Router) {
		router.Post("/register", r.RouteParams.UserController.Register)
		router.Post("/login", r.RouteParams.UserController.Login)
	})
}
