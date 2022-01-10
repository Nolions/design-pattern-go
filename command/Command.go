package command

type Command interface {
	Execute()
}

type OnCommend struct {
	Device Device
}

func (o *OnCommend) Execute()  {
	o.Device.On()
}

type OffCommend struct {
	Device Device
}

func (o *OffCommend) Execute()  {
	o.Device.Off()
}