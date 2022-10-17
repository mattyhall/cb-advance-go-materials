1. When do you need to use a mutex?

2. Research what contention is. How does it relate to locking?

3. What is a data race?

4. Implement a map that is protected by a lock using the starting code below. Fill it with lots of values and then create several goroutines that randomly get/set values. Run using `go run -race main.go` - if you haven't protected the map correctly it will output a warning about a data race

```go
type ThreadSafeMap struct {
	mu sync.Mutex
	m map[string]int
}

// Get returns a pointer to the integer and whether the key was in the map
func (m *ThreadSafeMap) Get(key string) (*int, bool) {}

// Put inserts the key into the map with the value "value". It should overwrite any value that is already there for that key
func (m *ThreadSafeMap) Put(key string, value int) {}
```
