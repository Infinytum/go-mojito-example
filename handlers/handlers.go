package handlers

import (
	"github.com/infinytum/go-mojito"
	"github.com/infinytum/go-mojito-example/services"
)

type HomeContext struct {
	Access services.Access `container:"type"`
}

func Home(req *mojito.Request, res *mojito.Response, ctx HomeContext) error {
	res.ViewBag.Set("lastVisited", ctx.Access.LastVisitedAt())
	ctx.Access.RecordVisit()
	return res.View("Home")
}
