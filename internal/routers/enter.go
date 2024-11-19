package routers

import (
	"GO-CRM-API-SHOPDEV/internal/routers/manage"
	"GO-CRM-API-SHOPDEV/internal/routers/user"
)

type RouterGroup struct {
	User   user.UserRouterGroup
	Manage manage.ManageRouteGroup
}

var RouterGroupApp = new(RouterGroup)
