# type WaitGroup

A WaitGroup waits for a collection of goroutines to finish.The main goroutine calls Add to set the number of goroutines to wait for. Then each of the goroutines runs and calls Done when finished. At the same time, Wait can be used to block until all goroutines have finished.

# type Mutex

A Mutex is a mutual exclusion lock. 
Lock()
Unlock()

# type RWMutex

A RWMutex is a reader/writer mutual exclusion lock.
The lock can be held by an arbitrary number of readers or a single writer. 
Mainly used when we have multiple read operations and single write operation.
Gives us ability for concurrent read

RLock(), RUnlock -> read operations
Lock(), Unlock -> write operations

RLock is a shared read lock. When a lock is taken with it, other threads* can also take their own lock with RLock. This means multiple threads* can read at the same time. It's semi-exclusive.

If the mutex is read locked, a call to Lock is blocked**. If one or more readers hold a lock, you cannot write.

If the mutex is write locked (with Lock), RLock will block

https://stackoverflow.com/questions/19148809/how-to-use-rwmutex-in-golang


