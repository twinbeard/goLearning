package routers

import (
	"github.com/twinbeard/goLearning/internal/routers/manager"
	"github.com/twinbeard/goLearning/internal/routers/user"
)

type RouterGroup struct {
	User    user.UserRouterGroup
	Manager manager.ManagerRouterGrop
}

var RouterGroupApp = new(RouterGroup)

// new is a built-in function that allocates memory, not initialize.
