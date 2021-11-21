package singleton

import "sync"

//Singleton 是单例模式类
type Singleton struct{}

var singleton *Singleton
var once sync.Once

//GetInstance 用于获取单例模式对象
func GetInstance() *Singleton {
	//once.Do 经过我的深度分析，once 内部也是加锁了，而且也进行了两次判断，
	//只不过判断的量done为原子方式的读取，然后新建实例，一样加了锁
	once.Do(func() {
		singleton = &Singleton{}
	})
	return singleton
}
