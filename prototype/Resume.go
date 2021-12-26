package prototype

type Resume interface {
	clone() Resume
	copy() Resume
}

type Basic struct {
	name       string
	age        int
	resumeType string
}

func (b *Basic) clone() *Basic {
	c := *b
	return &c
}

func (b *Basic) copy() Basic {
	c := *b
	return c
}
