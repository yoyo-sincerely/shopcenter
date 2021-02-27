package controller

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/types"
)

type Handler struct {
	HTML     func(ctx *context.Context, panel types.Panel, ops ...template.ExecuteOptions)
	HTMLMenu func(ctx *context.Context, panel types.Panel, ops ...template.ExecuteOptions)
}

func NewHandler( /*...*/ ) *Handler {
	return &Handler{
		// ...
	}
}

func (h *Handler) Update( /*...*/ ) {
	// ...
}
