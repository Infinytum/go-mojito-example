package handlers

import (
	"github.com/infinytum/go-scalar"
	"github.com/infinytum/go-scalar-example/services"
)

type HomeContext struct {
	Access services.Access `container:"type"`
}

func Home(req *scalar.Request, res *scalar.Response, ctx HomeContext) error {
	res.ViewBag.Set("lastVisited", ctx.Access.LastVisitedAt())
	ctx.Access.RecordVisit()
	return res.View("Home")
}
