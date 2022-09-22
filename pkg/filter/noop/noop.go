package noopfilter

import (
	"context"
	"encoding/json"

	"github.com/Fishwaldo/mouthpiece/pkg/filter"
	"github.com/Fishwaldo/mouthpiece/pkg/interfaces"
	"github.com/Fishwaldo/mouthpiece/pkg/mperror"
	"github.com/google/uuid"
)

func init() {
	filter.RegisterFilterImpl("NoOpFilter", NoOpFilterFactory{})
	Messages = make(map[uuid.UUID]bool)
}

var Messages map[uuid.UUID]bool


type NoOpFilter struct {
	config      *NoOpFilterConfig
}

type NoOpFilterConfig struct {
}

func (c *NoOpFilterConfig) AsJSON() (string, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *NoOpFilterConfig) FromJSON(data string) error {
	return json.Unmarshal([]byte(data), c)
}

type NoOpFilterFactory struct {
}

func (sff NoOpFilterFactory) FilterFactory(ctx context.Context, config string) (interfaces.FilterImplI, error) {
	var cfg NoOpFilterConfig
	if err := cfg.FromJSON(config); err != nil {
		return nil, mperror.ErrFilterConfigInvalid
	}
	return &NoOpFilter{config: &cfg}, nil
}

func (sff NoOpFilterFactory) DefaultConfig(ctx context.Context) interfaces.MarshableConfigI {
	return &NoOpFilterConfig{
	}
}

func (sf *NoOpFilter) Init(ctx context.Context) error {
	return nil
}

func (sf *NoOpFilter) Process(ctx context.Context, msg interfaces.MessageI) (interfaces.FilterAction, error) {
	Messages[msg.GetID()] = true
	return interfaces.FilterPass, nil
}

func (sf *NoOpFilter) FilterName() string {
	return "NoOpFilter"
}

func (sf *NoOpFilter) SetConfig(ctx context.Context, config interfaces.MarshableConfigI) error {
	var ok bool
	if sf.config, ok = config.(*NoOpFilterConfig); !ok {
		return mperror.ErrFilterConfigInvalid
	}
	return nil
}
func (sf *NoOpFilter) GetConfig(ctx context.Context) (interfaces.MarshableConfigI, error) {
	return sf.config, nil
}

var _ interfaces.FilterImplI = (*NoOpFilter)(nil)
