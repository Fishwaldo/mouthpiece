package filter

import (
	"context"
	"errors"

	"github.com/Fishwaldo/mouthpiece/pkg/interfaces"
)

var (
	FilterImpl map[string]FilterImplFactory
)


type FilterImplFactory interface {
	FilterFactory(context.Context, string) (interfaces.FilterImplI, error)
	DefaultConfig(context.Context) (interfaces.MarshableConfigI)
}

func RegisterFilterImpl(name string, factory FilterImplFactory) {
	if FilterImpl == nil {
		FilterImpl = make(map[string]FilterImplFactory)
	}
	FilterImpl[name] = factory
}

func GetNewFilterImpl(ctx context.Context, name string, config string) (interfaces.FilterImplI, error) {
	if flt, ok := FilterImpl[name]; ok {
		return flt.FilterFactory(ctx, config)
	}
	return nil, errors.New("FilterImpl Not Found")
}

func GetFilterImplDefaultConfig(ctx context.Context, name string) (interfaces.MarshableConfigI, error) {
	if flt, ok := FilterImpl[name]; ok {
		return flt.DefaultConfig(ctx), nil
	}
	return nil, errors.New("FilterImpl Not Found")
}

func GetFilterImpls(ctx context.Context) []string {
	var a []string
	a = make([]string, len(FilterImpl))
	for k := range FilterImpl {
		a = append(a, k)
	}
	return a
}
