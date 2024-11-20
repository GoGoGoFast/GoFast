package sync

import "sync"

// Counter is a thread-safe counter.
type Counter struct {
	mu    sync.Mutex
	value int
}

// NewCounter creates a new Counter.
func NewCounter() *Counter {
	return &Counter{}
}

// Increment increases the counter by 1.
func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

// Decrement decreases the counter by 1.
func (c *Counter) Decrement() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value--
}

// Value returns the current value of the counter.
func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

// WorkerPool is a simple worker pool implementation.
type WorkerPool struct {
	tasks chan func()
	wg    sync.WaitGroup
}

// NewWorkerPool creates a new WorkerPool.
func NewWorkerPool(workerCount int) *WorkerPool {
	pool := &WorkerPool{
		tasks: make(chan func()),
	}
	pool.wg.Add(workerCount)
	for i := 0; i < workerCount; i++ {
		go pool.worker()
	}
	return pool
}

// worker is the worker goroutine in the worker pool.
func (p *WorkerPool) worker() {
	defer p.wg.Done()
	for task := range p.tasks {
		task()
	}
}

// Submit submits a task to the worker pool.
func (p *WorkerPool) Submit(task func()) {
	p.tasks <- task
}

// ConcurrentQueue is a simple thread-safe queue.
type ConcurrentQueue struct {
	items []interface{}
	mu    sync.Mutex
}

// NewConcurrentQueue creates a new ConcurrentQueue.
func NewConcurrentQueue() *ConcurrentQueue {
	return &ConcurrentQueue{
		items: make([]interface{}, 0),
	}
}

// Enqueue adds an element to the end of the queue.
func (q *ConcurrentQueue) Enqueue(item interface{}) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.items = append(q.items, item)
}

// Dequeue removes an element from the beginning of the queue and returns it.
func (q *ConcurrentQueue) Dequeue() (interface{}, bool) {
	q.mu.Lock()
	defer q.mu.Unlock()
	if len(q.items) == 0 {
		return nil, false
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item, true
}

// Size returns the number of elements in the queue.
func (q *ConcurrentQueue) Size() int {
	q.mu.Lock()
	defer q.mu.Unlock()
	return len(q.items)
}

// Semaphore is a simple semaphore implementation.
type Semaphore struct {
	ch chan struct{}
}

// NewSemaphore creates a new Semaphore.
func NewSemaphore(max int) *Semaphore {
	return &Semaphore{
		ch: make(chan struct{}, max),
	}
}

// Acquire acquires a semaphore.
func (s *Semaphore) Acquire() {
	s.ch <- struct{}{}
}

// Release releases a semaphore.
func (s *Semaphore) Release() {
	<-s.ch
}

// ConcurrentMap is a thread-safe Map implementation.
type ConcurrentMap struct {
	mu    sync.RWMutex
	store map[interface{}]interface{}
}

// NewConcurrentMap creates a new ConcurrentMap.
func NewConcurrentMap() *ConcurrentMap {
	return &ConcurrentMap{
		store: make(map[interface{}]interface{}),
	}
}

// Put puts a key-value pair into the Map.
func (m *ConcurrentMap) Put(key, value interface{}) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.store[key] = value
}

// Get gets a value from the Map by key.
func (m *ConcurrentMap) Get(key interface{}) (interface{}, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	value, ok := m.store[key]
	return value, ok
}

// Delete deletes a key-value pair from the Map.
func (m *ConcurrentMap) Delete(key interface{}) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.store, key)
}

// Size returns the number of key-value pairs in the Map.
func (m *ConcurrentMap) Size() int {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return len(m.store)
}

// ConcurrentSet is a thread-safe Set implementation.
type ConcurrentSet struct {
	mu    sync.RWMutex
	store map[interface{}]struct{}
}

// NewConcurrentSet creates a new ConcurrentSet.
func NewConcurrentSet() *ConcurrentSet {
	return &ConcurrentSet{
		store: make(map[interface{}]struct{}),
	}
}

// Add adds an element to the Set.
func (s *ConcurrentSet) Add(item interface{}) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.store[item] = struct{}{}
}

// Remove removes an element from the Set.
func (s *ConcurrentSet) Remove(item interface{}) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.store, item)
}

// Contains checks if an element is in the Set.
func (s *ConcurrentSet) Contains(item interface{}) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	_, ok := s.store[item]
	return ok
}

// Size returns the number of elements in the Set.
func (s *ConcurrentSet) Size() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.store)
}

// Shutdown shuts down the worker pool and waits for all tasks to complete.
func (p *WorkerPool) Shutdown() {
	close(p.tasks)
	p.wg.Wait()
}
