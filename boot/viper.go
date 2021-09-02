package boot

var Viper = new(_viper)

type _viper struct {
	err  error
	path string
}

func (v *_viper) Initialize() {

}
