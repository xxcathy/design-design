package adapter

type ITarget interface {
	f1()
	f2()
	f3()
}

type Adaptee struct{}

func (adaptee Adaptee) fa() {}
func (adaptee Adaptee) fb() {}
func (adaptee Adaptee) fc() {}

type Adaptor struct {
	adaptee *Adaptee
}

func NewAdaptor(adaptee *Adaptee) *Adaptor {
	return &Adaptor{
		adaptee: adaptee,
	}
}

func (adaptor *Adaptor) f1() {
	adaptor.adaptee.fa()
}

func (adaptor *Adaptor) f2() {
	adaptor.adaptee.fa()
	adaptor.adaptee.fb()
}

func (adaptor *Adaptor) f3() {
	adaptor.adaptee.fa()
	adaptor.adaptee.fb()
	adaptor.adaptee.fc()
}
