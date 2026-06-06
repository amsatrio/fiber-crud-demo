package auth

import (
	"errors"

	"github.com/amsatrio/fiber-crud-demo/app/dto/response"
	"github.com/amsatrio/fiber-crud-demo/app/util"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

type AuthHandler struct {
	service  AuthService
	validate *validator.Validate
}

func NewAuthHandler(service AuthService, validate *validator.Validate) *AuthHandler {
	return &AuthHandler{
		service:  service,
		validate: validate,
	}
}

func (a *AuthHandler) AuthLogin(c fiber.Ctx) error {

	res := &response.Response{}
	payload := new(AuthLoginRequest)

	if err := c.Bind().Body(payload); err != nil {
		res.ErrMessage(c.Path(), fiber.StatusBadRequest, "parse body error: "+err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(res)
	}

	entity, err := a.service.Login(*payload.Username, *payload.Password)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		res.ErrMessage(c.Path(), fiber.StatusBadRequest, "data not found")
		return c.Status(res.Status).JSON(res)
	}

	if err != nil {
		util.Log("ERROR", "controllers", "AuthIndex", err.Error())
		res.ErrMessage(c.Path(), fiber.StatusBadRequest, "get data error: "+err.Error())
		return c.Status(res.Status).JSON(res)
	}

	res.Ok(c.Path(), entity)
	return c.Status(res.Status).JSON(res)
}

func (a *AuthHandler) AuthRegister(c fiber.Ctx) error {
	res := &response.Response{}

	res.Ok(c.Path(), nil)
	return c.Status(res.Status).JSON(res)
}

func (a *AuthHandler) AuthPublic(c fiber.Ctx) error {
	res := &response.Response{}

	res.Ok(c.Path(), nil)
	return c.Status(res.Status).JSON(res)
}

func (a *AuthHandler) AuthFaskes(c fiber.Ctx) error {
	res := &response.Response{}

	res.Ok(c.Path(), nil)
	return c.Status(res.Status).JSON(res)
}
func (a *AuthHandler) AuthDokter(c fiber.Ctx) error {
	res := &response.Response{}

	res.Ok(c.Path(), nil)
	return c.Status(res.Status).JSON(res)
}
func (a *AuthHandler) AuthPasien(c fiber.Ctx) error {
	res := &response.Response{}

	res.Ok(c.Path(), nil)
	return c.Status(res.Status).JSON(res)
}
func (a *AuthHandler) AuthAdmin(c fiber.Ctx) error {
	res := &response.Response{}

	res.Ok(c.Path(), nil)
	return c.Status(res.Status).JSON(res)
}
