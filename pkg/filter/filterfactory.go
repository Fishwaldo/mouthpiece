package filter

import (
	"context"
	"errors"
	"fmt"
)

type FilterFactoryFN func(context.Context, []Filterconfig) (FilterImplI, error)

type FilterImplDetails struct {
	Factory FilterFactoryFN
}

var (
	FilterImpl map[string]FilterImplDetails
)

func RegisterFilterImpl(name string, factory FilterImplDetails) {
	fmt.Println("RegisterFilterImpl", name)
	if FilterImpl == nil {
		FilterImpl = make(map[string]FilterImplDetails)
	}
	FilterImpl[name] = factory
}

func GetNewFilterImpl(ctx context.Context, name string, config []Filterconfig) (FilterImplI, error) {
	if flt, ok := FilterImpl[name]; ok {
		return flt.Factory(ctx, config)
	}
	return nil, errors.New("FilterImpl Not Found")
}

func GetFilterImpls(ctx context.Context) []string {
	var a []string
	for k := range FilterImpl {
		a = append(a, k)
	}
	return a
}
