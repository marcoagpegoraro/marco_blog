package initializers

import (
	"time"

	"github.com/patrickmn/go-cache"
)

var Cache *cache.Cache

func ConfigureCache() {
	expirationTime := 5 * time.Minute
	purgeTime := 5 * time.Minute
	Cache = cache.New(expirationTime, purgeTime)
}
