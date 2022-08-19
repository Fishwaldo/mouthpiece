package restapi

import (
	"github.com/danielgtaylor/huma"
	"github.com/go-playground/validator/v10"
)

func checkErrors(ctx huma.Context, err error) {
	ve, ok := err.(validator.ValidationErrors)
	if ok {
		for _, e := range ve {
			ctx.AddError(e)
		}
	} else {
		ctx.AddError(err)
	}
}