package main

import (
	"fmt"
	"github.com/akyoto/cache"
	"time"
)

func main() {
	// New cache
	c := cache.New(5 * time.Minute)

	// Put something into the cache
	c.Set("a", "b", 1*time.Minute)

	// Read from the cache
	obj, _ := c.Get("a")

	// Convert the type
	fmt.Println(obj.(string))
}
