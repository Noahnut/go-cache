# Go-Cache
## Introduce
Golang Cache in the application layer, use the interface to let's user can simply 
change the cache replacement policy

## Cache Replacement Policy
- [x] Least Frequently Used (LFU)  

- [ ] Tiny Least Frequently Used (TinyLFU) 
 
## How to use
```
package main

import (
	"fmt"
	gocache "go-cache"
)

func main() {

	// new the cache instance and set the policy
	// default is the Least Frequently Used
	// default cache size is 1024
	cache := gocache.NewCache(gocache.LFU_POLICY)

	// set key value
	cache.Set(123, "123")

	// get the value by key, and check exist in the cache or not
	value, exist := cache.Get(123)
	fmt.Println(value, exist)

	// check the key exist in the cache or not
	exist = cache.Contains(123)

	fmt.Println(exist)

	// resize the cache
	cache.Resize(2048)

	// delete the key from the cache
	deleted := cache.Delete(123)

	fmt.Println(deleted)

	// clean all in the cache
	cache.Clean()
}

```