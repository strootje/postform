package modules

import (
	"github.com/gofiber/fiber/v2"
	lua "github.com/yuin/gopher-lua"
)

func NewFormLoader(c *fiber.Ctx) lua.LGFunction {
	return func(l *lua.LState) int {
		mod := l.NewUserData()
		mod.Value = c
		l.SetMetatable(mod, l.NewTable())
		l.SetField(mod.Metatable, "__index", l.NewFunction(formIndexFunc))
		l.SetField(mod.Metatable, "__newindex", l.NewFunction(formNewIndexFunc))
		l.Push(mod)
		return 1
	}
}

func formIndexFunc(l *lua.LState) int {
	if data, ok := l.CheckUserData(1).Value.(*fiber.Ctx); ok {
		if value := data.FormValue(l.ToString(2)); value != "" {
			l.Push(lua.LString(value))
			return 1
		}
	}

	return 0
}

func formNewIndexFunc(l *lua.LState) int {
	l.Error(lua.LString("Not allowed to set new values"), 0)
	return 0
}
