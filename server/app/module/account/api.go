package account

import (
	"fcompressor/app"
	"fcompressor/common/response"

	"github.com/gofiber/fiber/v2"
)

type accountService struct {
	app   *fiber.App
	store Store
}

func newService(app *fiber.App) app.Service {
	return &accountService{
		app:   app,
		store: NewStore(),
	}
}

func (s *accountService) initSession(ctx *fiber.Ctx) error {
	user := ctx.Locals("user_id").(string)

	res, err := s.store.Session(user)
	if err != nil {
		return response.InternalServerError(ctx, err)
	}

	return response.Ok(ctx, res)
}

func (s *accountService) CreateRoutes() {
	r := s.app.Group("/api/v1/account")

	r.Get("/session", s.initSession)
}

func init() {
	app.RegisterService(newService)
}
