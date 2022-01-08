package iterator

type Iterator interface {
	First() *Student
	Next() *Student
	IsFinal() bool
}

type Student struct {
	Name string
}

func NewStudent(name string) *Student {
	return &Student{
		Name: name,
	}
}

type StudentIterator struct {
	Index int
	Users []*Student
}

func NewStudentIterator(s []*Student) StudentIterator {
	return StudentIterator {
		Users: s,
	}
}

func (p *StudentIterator) First() *Student {
	if len(p.Users) > 0 {
		return p.Users[0]
	}

	return nil
}

func (p *StudentIterator) Next() *Student {
	if p.IsFinal() {
		nowIndex := p.Index
		p.Index++

		return p.Users[nowIndex]
	}

	return nil
}

func (p *StudentIterator) IsFinal() bool {
	if p.Index < len(p.Users) {
		return true
	}

	return false
}
