package user

// Gather all routers in this group
type UserRouterGroup struct {
	UserRouter
	ProductRouter
}
