package decorator

type IFunc interface {
	f()
}

type FuncA struct {
	otherFunc IFunc
}

func NewFuncA(b IFunc) *FuncA {
	return &FuncA{
		otherFunc: b,
	}
}

func (a FuncA) f() {
	// do something
	a.otherFunc.f()
	// do something
}