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
