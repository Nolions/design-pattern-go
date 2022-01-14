package interpret

type AndOperate struct {
	Expr1 Expression
	Expr2 Expression
}

func NewAndOperate(expr1, expr2 Expression) *AndOperate {
	return &AndOperate{
		Expr1: expr1,
		Expr2: expr2,
	}
}

func (a *AndOperate) Interpret(context string) bool {
	return a.Expr1.Interpret(context) && a.Expr2.Interpret(context)
}
