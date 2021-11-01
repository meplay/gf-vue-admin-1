package interfaces

type Router interface {
	// Public 公有路由组
	Public() Router
	// Private 私有路由组
	Private() Router
	// PublicWithoutRecord 公有路由组 不需要记录日志
	PublicWithoutRecord() Router
	// PrivateWithoutRecord 私有路由组 不需要记录日志
	PrivateWithoutRecord() Router
}

type PublicRouter interface {
	// Public 公有路由组
	Public() PublicRouter
	// PublicWithoutRecord 公有路由组 不需要记录日志
	PublicWithoutRecord() PublicRouter
}

type PrivateRouter interface {
	// Private 私有路由组
	Private() PrivateRouter
	// PrivateWithoutRecord 私有路由组 不需要记录日志
	PrivateWithoutRecord() PrivateRouter
}
