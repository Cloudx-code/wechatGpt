package local_cache

import (
	"fmt"
	"testing"
	"time"

	"github.com/patrickmn/go-cache"
)

func TestCache(t *testing.T) {
	var cache1 = cache.New(time.Minute*5, time.Minute*5)
	var cache2 = cache.New(time.Minute*5, time.Minute*5)
	cache1.Set("test", "你好", 0)
	if a, ok := cache1.Get("test"); ok {
		fmt.Println("a")
		fmt.Println(a)
	}
	if a, ok := cache2.Get("test"); ok {
		fmt.Println("b")
		fmt.Println(a)
	}
}
