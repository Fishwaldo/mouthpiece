package filter

import (
	"context"

	"github.com/Fishwaldo/mouthpiece/pkg/interfaces"
	"github.com/Fishwaldo/mouthpiece/pkg/mperror"
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
	return nil, mperror.ErrFilterImplNotFound
}

func GetFilterImplDefaultConfig(ctx context.Context, name string) (interfaces.MarshableConfigI, error) {
	if flt, ok := FilterImpl[name]; ok {
		return flt.DefaultConfig(ctx), nil
	}
	return nil, mperror.ErrFilterImplNotFound
}

func GetFilterImpls(ctx context.Context) []string {
	var a []string
	for k := range FilterImpl {
		a = append(a, k)
	}
	return a
}
