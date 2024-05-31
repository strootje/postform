package handlers

import (
	"fmt"

	"git.strooware.nl/postform/modules"
	"github.com/gofiber/fiber/v2"
	lua "github.com/yuin/gopher-lua"
)

func PostFormHandler(c *fiber.Ctx) error {
	state := lua.NewState(lua.Options{
		SkipOpenLibs: false,
	})
	defer state.Close()

	state.PreloadModule("form", modules.NewFormLoader(c))
	state.PreloadModule("headers", modules.NewHeadersLoader(c))
	state.PreloadModule("mailer", modules.NewMailerLoader(c))

	if err := state.DoFile("./scripts/contact.lua"); err != nil {
		return c.SendString(fmt.Sprintf("Lua Error:\n====================\n%s\n", err.Error()))
	}

	return c.RedirectBack("https://strooware.nl")
}
