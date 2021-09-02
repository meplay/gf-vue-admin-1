package boot

var Gorm = new(_gorm)

type _gorm struct{}

func (g *_gorm) Initialize() {}
