package modules

import (
	"github.com/gofiber/fiber/v2"
	lua "github.com/yuin/gopher-lua"
)

func NewHeadersLoader(c *fiber.Ctx) lua.LGFunction {
	return func(l *lua.LState) int {
		mod := l.NewUserData()
		mod.Value = c.GetReqHeaders()
		l.SetMetatable(mod, l.NewTable())
		l.SetField(mod.Metatable, "__index", l.NewFunction(headersIndexFunc))
		l.SetField(mod.Metatable, "__newindex", l.NewFunction(headersNewIndexFunc))
		l.Push(mod)
		return 1
	}
}

func headersIndexFunc(l *lua.LState) int {
	if data, ok := l.CheckUserData(1).Value.(map[string][]string); ok {
		if value, ok := data[l.ToString(2)]; ok {
			l.Push(lua.LString(value[0]))
			return 1
		}
	}

	return 0
}

func headersNewIndexFunc(l *lua.LState) int {
	l.Error(lua.LString("Not allowed to set new values"), 0)
	return 0
}
