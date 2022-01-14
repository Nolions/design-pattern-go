package interpret


type Expression interface {
	Interpret(context string) bool
}

type Operate struct {
	//Data string
	Data []string
}

func NewOperate(d []string) *Operate {
	return &Operate{
		Data: d,
	}
}

func (t *Operate) Interpret(s string) bool {
	for _, v:= range t.Data{
		if v == s {
			return true
		}
	}
	//if strings.Contains(s, t.Data) {
	//	return true
	//}
	return false
}
