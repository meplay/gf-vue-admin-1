package interfaces

type Router interface {
	Init()
}

func RouterInit(arg ...Router) {
	for i := 0; i < len(arg); i++ {
		arg[i].Init()
	}
}
