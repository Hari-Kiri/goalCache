# goalCache
Go language in memory caching library with life time (has expired) support.

## Usage:
- Download this module:
    ```shell
    go get github.com/Hari-Kiri/goalCache
    ```
- Instantiation:
    ```go
    myCache := goalCache.New[string, int]()
    ```
- Set:
    ```go
    myCache.Set("one", 1, 5*time.Second) // Set cache with key name "one", value for integer 1, and expired in five seconds
    ```
- Get:
    ```go
    value, found := myCache.Get("one") // Get cache with key name "one"
	if found {
		fmt.Printf("Value for key 'one': %v\n", value)
	} 
    if !found {
		fmt.Errorf("Key 'one' not found in the cache or has expired")
	}
    ```
- Delete:
    ```go
    myCache.Delete("one") // Delete cache with key name "one"
    ```

### Reference:
- [Quick tip: Implementing an in-memory cache in Go](https://www.alexedwards.net/blog/implementing-an-in-memory-cache-in-go)
