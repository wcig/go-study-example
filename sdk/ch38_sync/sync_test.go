package ch38_sync

// sync: 提供基本同步原语

// 类型
// 1.Cond: 表示一条件变量
// type Cond struct {
//    // L is held while observing or changing the condition
//    L Locker
//    // contains filtered or unexported fields
// }
// func NewCond(l Locker) *Cond // 基于Locker l创建一Cond
// func (c *Cond) Broadcast()   // 唤醒所有等待c的goroutines (允许但不需要调用者在调用时保持c.L)
// func (c *Cond) Signal()      // 如果有的话,唤醒等待c的一goroutine (允许但不需要调用者在调用时保持c.L)
// func (c *Cond) Wait()        // 原子的解锁c.L并暂停调用goroutine的执行,后续恢复执行后,Wait在返回之前锁定c.L

// 2.Locker: 表示一可以锁定和解锁的对象
// type Locker interface {
//    Lock()
//    Unlock()
// }

// 3.Map: 跟Go的map[interface{}]interface{}比较类似,但是在多个goroutines并发使用是安全的,而且无序额外的锁或条件.加载,存储和删除开销在固定时间.
// type Map struct{
//    // contains filtered or unexported fields
// }
// func (m *Map) Delete(key interface{})                                               // 从m中删除key和其对应val
// func (m *Map) Load(key interface{}) (value interface{}, ok bool)                    // 从map中查找key对应val,如果没有则返回nil
// func (m *Map) LoadAndDelete(key interface{}) (value interface{}, loaded bool)       // 返回先前key的val,同时删除key,如果先前key存在则loaded返回true
// func (m *Map) LoadOrStore(key, value interface{}) (actual interface{}, loaded bool) // 返回先前存在的key对应val,不存在则存在并返回给定的key,val值,loaded表示key先前是否存在
// func (m *Map) Range(f func(key, value interface{}) bool)                            // 遍历m,如果f返回false则遍历停止
// func (m *Map) Store(key, value interface{})                                         // 存储key,val

// 4.Mutex: 互斥锁,零值是一个未锁定的互斥锁 (第一次使用后不能复制)
// type Mutex struct{
//    // contains filtered or unexported fields
// }
// func (m *Mutex) Lock()   // 锁定,如果锁已经在使用,则调用的goroutine会阻塞直到互斥锁可用
// func (m *Mutex) Unlock() // 解锁,改调用之前如果m是未锁定的则会导致运行时错误 (互斥锁与特定goroutine无关,允许一个goroutine锁定互斥锁另一个goroutine解锁)

// 5.type Once: 用于只执行一次操作的对象 (第一次使用后不能复制)
// type Once struct {
//    // contains filtered or unexported fields
// }
// func (o *Once) Do(f func()) // 只调用一次的函数 (解释Do被调用多次,也只有第一次调用才会执行函数f)

// 6.Pool: 一组可以单独保存和检索的临时对象. (存在在Pool的任何项目可能会在没有被通知情况下被自动删除,一个Pool可以在安全的被多个goroutines使用) (第一次使用后不能复制)
// type Pool struct {
//    New func() interface{}
//    // contains filtered or unexported fields
// }
// func (p *Pool) Get() interface{}  // 从Pool中选择任意项,将其从Pool中移除并返回给调用者
// func (p *Pool) Put(x interface{}) // 添加x到Pool中

// 7.RWMutex: 读写互斥锁,该锁可以被任意数量的readers和单个writer持有,零值是一未锁定的互斥锁 (它不限制并发的读，但是读写、写写无法同时进行) (第一次使用后不能复制)
// type RWMutex struct {
//    // contains filtered or unexported fields
// }
// func (rw *RWMutex) Lock()           // 锁定rw来写,读和写都会被锁定直到锁释放
// func (rw *RWMutex) RLock()          // 锁定rw来读,此时写会被锁定
// func (rw *RWMutex) RLocker() Locker // 返回Locker接口,该接口通过调用rw.RLock和rw.RUnlock实现Lock和Unlock方法
// func (rw *RWMutex) RUnlock()        // 取消rw的读锁,在调用RUnlock之前没有锁定将导致运行时错误
// func (rw *RWMutex) Unlock()         // 取消rw的写锁,在调用Unlock之前没有锁定将导致运行时错误

// 8.WaitGroup: 用于等待一组goroutines完成
// type WaitGroup struct {
//    // contains filtered or unexported fields
// }
// func (wg *WaitGroup) Add(delta int) // 增加待处理任务数
// func (wg *WaitGroup) Done()         // 表示一个任务已完成
// func (wg *WaitGroup) Wait()         // 等待所有任务完成
