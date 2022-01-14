package interpret

type OrOperate struct {
	Expr1 Expression
	Expr2 Expression
}

func NewOrOperate(expr1, expr2 Expression) *OrOperate {
	return &OrOperate{
		Expr1: expr1,
		Expr2: expr2,
	}
}

func (o *OrOperate) Interpret(context string) bool {
	return o.Expr1.Interpret(context) || o.Expr2.Interpret(context)
}
