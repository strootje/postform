package handlers

import (
	"fmt"
	"net/url"
	"slices"

	"github.com/gofiber/fiber/v2"
)

type (
	CustomError struct {
		Message string
	}
	FormSettings map[string]FormOpts

	FormOpts struct {
		AllowedHosts []string
	}
)

var forms = FormSettings{
	"811bb6c8": {
		AllowedHosts: []string{
			"magic-nonsense.com",
			"localhost:1313",
		},
	},
}

func FindFormMiddleware(c *fiber.Ctx) error {
	formId := c.Params("formId")

	if formId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(CustomError{
			Message: "formId is not provided",
		})
	}

	formSettings, found := forms[formId]

	if !found {
		return c.Status(fiber.StatusNotFound).JSON(CustomError{
			Message: fmt.Sprintf("form(%s) not found", formId),
		})
	}

	c.Locals("formSettings", formSettings)
	return c.Next()
}

func ValidateReferrerMiddleware(c *fiber.Ctx) error {
	ref, err := url.Parse(string(c.Request().Header.Referer()))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(CustomError{
			Message: err.Error(),
		})
	}

	formOpts, ok := c.Locals("formSettings").(FormOpts)

	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(CustomError{
			Message: "Unable to cast form..",
		})
	}

	if !slices.Contains(formOpts.AllowedHosts, ref.Host) {
		return c.Status(fiber.StatusForbidden).JSON(CustomError{
			Message: fmt.Sprintf("Ref(%s) is not allowed", ref.Host),
		})
	}

	return c.Next()
}
