package internal

var Tool = new(tool)

type tool struct {}

func (i *tool) GetA1Value(m map[string]string, d string) string {
	if v1, o1 := m["COMMENT"]; o1 {
		return v1
	} else {
		return d
	}
}