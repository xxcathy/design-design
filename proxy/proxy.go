package proxy

import "time"

type IUser interface {
	Register(telephone, password string)
}

type User struct{}

func (user *User) Register(telephone, password string) {}

type UserProxy struct {
	user   *User
	metric *Metric
}

func NewUserProxy() *UserProxy {
	return &UserProxy{
		user:   &User{},
		metric: &Metric{},
	}
}
func (userProxy *UserProxy) Register(telephone, password string) {
	startTime := time.Now().Unix()
	userProxy.user.Register(telephone, password)
	endTime := time.Now().Unix()
	userProxy.metric.Report(startTime, endTime)
}

type Metric struct{}

func (m *Metric) Report(startTime, endTime int64) {}

//func main() {
//	user := NewUserProxy()
//	user.Register("xxxx", "xxxxxx")
//}
