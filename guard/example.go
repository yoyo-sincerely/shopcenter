package guard

import (
	"github.com/GoAdminGroup/go-admin/context"
)

type ExampleParam struct {
	Param string
}

func (g *Guardian) Example(ctx *context.Context) {

	param := ctx.Query("param")

	ctx.SetUserValue("example", &ExampleParam{
		Param: param,
	})
	ctx.Next()
}

func GetExampleParam(ctx *context.Context) *ExampleParam {
	return ctx.UserValue["example"].(*ExampleParam)
}
