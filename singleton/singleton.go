package singleton

import "sync"

type singleton struct {}

func newSingleton() *singleton {
	return &singleton{}
}

/*************************** 饿汉模式 *****************************/
var (
	instance1 *singleton
)

// 在初始化时就加载
// 优点：1、实例的创建是线程安全的，2、及早暴露问题
// 缺点：不支持延迟加载（懒加载可以避免占用过多内存或影响效率）
func init() {
	instance1 = newSingleton()
}

func GetSingleton1() *singleton {
	return instance1
}

/*************************** 懒汉模式 *****************************/
var (
	instance2 *singleton
	once      sync.Once
	//lock sync.Mutex
)

// 使用时加载
func GetSingleton2() *singleton {
	// 通用写法：双重检测
	//if instance2 == nil {
	//	lock.Lock()
	//	if instance2 == nil {
	//		instance2 = newSingleton()
	//	}
	//	lock.Unlock()
	//}
	//return instance2

	// golang 实现
	if instance2 == nil {
		once.Do(func() {
			instance2 = newSingleton()
		})
	}
	return instance2
}