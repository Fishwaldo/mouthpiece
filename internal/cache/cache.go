package cache

import (
	"time"

	"github.com/eko/gocache/v3/cache"
	"github.com/eko/gocache/v3/store"
	"github.com/eko/gocache/v3/metrics"
	gc "github.com/patrickmn/go-cache"
)

var (
	goCachebackend *gc.Cache
	goCacheStore  *store.GoCacheStore
	goCacheMetrics *metrics.Prometheus
)

func Start() {
	goCachebackend = gc.New(5*time.Minute , 10*time.Minute)
	goCacheStore = store.NewGoCache(goCachebackend)
	goCacheMetrics = metrics.NewPrometheus("mouthpiece")
}


func GetNewCache[ T any]() *cache.MetricCache[T] {
	return cache.NewMetric[T](goCacheMetrics,  cache.New[T](goCacheStore))
}
