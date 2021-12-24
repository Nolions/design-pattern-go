package factory

type PhoneFactory interface {
	Create() Phone
}

type HTCTaiwanFactory struct {
}

func (h *HTCTaiwanFactory) Create() Phone {
	return &HTC{
		Origin: "Taiwan",
	}
}

type HTCChinaFactory struct {
}

func (h *HTCChinaFactory) Create() Phone {
	return &HTC{
		Origin: "China",
	}
}

type ASUSTaiwanFactory struct {
}

func (a *ASUSTaiwanFactory) Create() Phone {
	return &ASUS{
		Origin: "Taiwan",
	}
}

type ASUSChinaFactory struct {
}

func (h *ASUSChinaFactory) Create() Phone {
	return &ASUS{
		Origin: "China",
	}
}

func CreateFactory(enum PhoneEnum) PhoneFactory {
	switch enum {
	case HTCTaiwanPhoneEnum:
		return &HTCTaiwanFactory{}
	case HTCChinaPhoneEnum:
		return &HTCChinaFactory{}
	case ASUSTaiwanPhoneEnum:
		return &ASUSTaiwanFactory{}
	case ASUSChinaPhoneEnum:
		return &ASUSChinaFactory{}
	default:
		return &HTCTaiwanFactory{}
	}
}

type PhoneEnum int

const (
	HTCTaiwanPhoneEnum PhoneEnum = iota
	HTCChinaPhoneEnum
	ASUSTaiwanPhoneEnum
	ASUSChinaPhoneEnum
)
