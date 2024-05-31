package modules

import (
	"net/smtp"

	"git.strooware.nl/postform/common"
	"github.com/gofiber/fiber/v2"
	lua "github.com/yuin/gopher-lua"
)

func NewMailerLoader(c *fiber.Ctx) lua.LGFunction {
	return func(l *lua.LState) int {
		mod := l.NewUserData()

		indexTable := l.NewTable()
		l.SetFuncs(indexTable, map[string]lua.LGFunction{
			"send": mailerSend,
		})

		l.SetMetatable(mod, l.NewTable())
		l.SetField(mod.Metatable, "__index", indexTable)

		l.Push(mod)
		return 1
	}
}

func mailerSend(l *lua.LState) int {
	body := l.ToString(1)

	addr := common.Config.Smtp.Server
	auth := smtp.CRAMMD5Auth(common.Config.Smtp.Username, common.Config.Smtp.Password)
	from := common.Config.Smtp.FromAddress
	to := "bas@strootje.com"

	fullBody := "From: " + from + "\r\n" +
		"To: " + to + "\r\n" +
		"Subject: Testing this email..\r\n\r\n" +
		body + "\r\n"

	if err := smtp.SendMail(addr, auth, from, []string{to}, []byte(fullBody)); err != nil {
		panic(err)
	}

	return 0
}
