package local_cache

import (
	"time"

	"github.com/patrickmn/go-cache"
)

var cacheProxy *cache.Cache

func InitCache() {
	cacheProxy = cache.New(time.Minute*10, time.Minute*5)
}
