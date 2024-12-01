package goalcache

import (
	"sync"
	"time"
)

// item represents a cache item with a value and an expiration time.
type item[Value any] struct {
	value  Value
	expiry time.Time
}

// isExpired checks if the cache item has expired.
func (Item item[V]) isExpired() bool {
	return time.Now().After(Item.expiry)
}

// LifeTimeCache is a generic cache implementation which has expired.
type LifeTimeCache[Key comparable, Value any] struct {
	items map[Key]item[Value] // The map storing cache items.
	mutex sync.Mutex          // Mutex for controlling concurrent access to the cache.
}

// New creates a new LifeTimeCache instance and run a goroutine to remove
// any item has expired every 1 second with method loop through the cache items and
// delete which has expired.
func New[Key comparable, Value any]() *LifeTimeCache[Key, Value] {
	result := &LifeTimeCache[Key, Value]{
		items: make(map[Key]item[Value]),
	}

	go func() {
		for range time.Tick(1 * time.Second) {
			result.mutex.Lock()

			for key, item := range result.items {
				if !item.isExpired() {
					continue
				}
				delete(result.items, key)
			}

			result.mutex.Unlock()
		}
	}()

	return result
}

// Set will add new item to the cache with key, value, and life time to expired.
func (cache *LifeTimeCache[Key, Value]) Set(key Key, value Value, lifeTime time.Duration) {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()

	cache.items[key] = item[Value]{
		value:  value,
		expiry: time.Now().Add(lifeTime),
	}
}

// Get will retrieve the value associated with the given key from the cache.
// If the key is not found, it will return the zero value for Value and false.
// If the cache found and has expired it will delete the cache and return the
// value for Value and false. Otherwise return the value and true.
func (cache *LifeTimeCache[Key, Value]) Get(key Key) (Value, bool) {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()

	item, found := cache.items[key]
	if !found {
		return item.value, false
	}

	if item.isExpired() {
		delete(cache.items, key)
		return item.value, false
	}

	return item.value, true
}

// Delete the item with the given specific key from the parameter inside cache.
func (cache *LifeTimeCache[Key, Value]) Delete(key Key) {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()

	delete(cache.items, key)
}
