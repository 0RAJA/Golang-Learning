package msgpool

import "sync"

/*
保证一个类仅有一个实例，并提供一个访问它的全局访问点。
有一些对象通常我们只需要一个共享的实例，比如线程池、全局缓存、对象池等，这种场景下就适合使用单例模式
注意:
1. 限制调用者直接实例化该对象
2. 为该对象的单例提供一个全局唯一的访问方法
在实际开发中，我们经常会遇到需要频繁创建和销毁的对象。
频繁的创建和销毁一则消耗CPU，二则内存的利用率也不高，通常我们都会使用对象池技术来进行优化。
考虑我们需要实现一个消息对象池，因为是全局的中心点，管理所有的Message实例，所以将其实现成单例
*/

type Message struct {
	cnt int
}

type messagePool struct {
	pool *sync.Pool
}

var msgPool = &messagePool{pool: &sync.Pool{New: func() interface{} { return new(Message) }}}

// Instance 访问消息池单例的唯一方法
func Instance() *messagePool {
	return msgPool
}

// AddMsg 添加消息
func (m *messagePool) AddMsg(msg *Message) {
	m.pool.Put(msg)
}

// GetMsg 获取消息
func (m *messagePool) GetMsg() *Message {
	return m.pool.Get().(*Message)
}
