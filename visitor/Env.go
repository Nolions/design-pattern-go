package visitor

type Visitor interface {
	Visit() string
}

type IElement interface {
	accept(visitor Visitor) string
}

type Element struct {
}

func (e Element) accept(visitor Visitor) string {
	return visitor.Visit()
}

type Env struct {
	Element
}

func (e Env) Print(visitor Visitor) string {
	return e.Element.accept(visitor)
}
