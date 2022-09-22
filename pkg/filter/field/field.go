package field

import (
	"context"
	"encoding/json"
	"regexp"
	"strings"

	"github.com/Fishwaldo/mouthpiece/pkg/filter"
	"github.com/Fishwaldo/mouthpiece/pkg/interfaces"
	"github.com/Fishwaldo/mouthpiece/pkg/mperror"
)

func init() {
	filter.RegisterFilterImpl("FieldFilter", FieldFilterFactory{})
}

type FieldFilterOp int

const (
	FieldFilterOpEQ FieldFilterOp = iota
	FieldFilterOpContains
	FieldFilterOpPresent
	FieldFilterOpMissing
	FieldFilterOpRegex
	FieldFilterOpNoOp
)

type FieldFilter struct {
	config      *FieldFilterConfig
	regexpmatch *regexp.Regexp
}

type FieldFilterConfig struct {
	Op    FieldFilterOp
	Field string
	Value string
}

func (c *FieldFilterConfig) AsJSON() (string, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *FieldFilterConfig) FromJSON(data string) error {
	return json.Unmarshal([]byte(data), c)
}

type FieldFilterFactory struct {
}

func (sff FieldFilterFactory) FilterFactory(ctx context.Context, config string) (interfaces.FilterImplI, error) {
	var cfg FieldFilterConfig
	if err := cfg.FromJSON(config); err != nil {
		return nil, mperror.ErrFilterConfigInvalid
	}
	return &FieldFilter{config: &cfg}, nil
}

func (sff FieldFilterFactory) DefaultConfig(ctx context.Context) interfaces.MarshableConfigI {
	return &FieldFilterConfig{
		Op: FieldFilterOpNoOp,
	}
}

func (sf *FieldFilter) Init(ctx context.Context) error {
	return nil
}

func (sf *FieldFilter) Process(ctx context.Context, msg interfaces.MessageI) (interfaces.FilterAction, error) {
	switch sf.config.Op {
	case FieldFilterOpEQ:
		{
			if val, err := msg.GetField(ctx, sf.config.Field); err != nil {
				return interfaces.FilterNoMatch, err
			} else {
				if val == sf.config.Value {
					return interfaces.FilterMatch, nil
				}
				return interfaces.FilterNoMatch, nil
			}
		}
	case FieldFilterOpContains:
		{
			if val, err := msg.GetField(ctx, sf.config.Field); err != nil {
				return interfaces.FilterNoMatch, err
			} else {
				if strings.Contains(val, sf.config.Value) {
					return interfaces.FilterMatch, nil
				} else {
					return interfaces.FilterNoMatch, nil
				}
			}
		}
	case FieldFilterOpPresent:
		{
			if _, err := msg.GetField(ctx, sf.config.Field); err != nil {
				return interfaces.FilterNoMatch, nil
			} else {
				return interfaces.FilterMatch, nil
			}
		}
	case FieldFilterOpMissing:
		{
			if _, err := msg.GetField(ctx, sf.config.Field); err != nil {
				return interfaces.FilterMatch, nil
			} else {
				return interfaces.FilterNoMatch, nil
			}
		}
	case FieldFilterOpRegex:
		{
			if val, err := msg.GetField(ctx, sf.config.Field); err != nil {
				return interfaces.FilterNoMatch, err
			} else {
				ok, err := regexp.MatchString(sf.config.Value, val)
				if err != nil {
					return interfaces.FilterPass, err
				} else {
					if ok {
						return interfaces.FilterMatch, nil
					} else {
						return interfaces.FilterNoMatch, nil
					}
				}
			}
		}
	case FieldFilterOpNoOp:
		{
			return interfaces.FilterPass, nil
		}
	default:
		{
			return interfaces.FilterPass, nil
		}
	}
}

func (sf *FieldFilter) FilterName() string {
	return "FieldFilter"
}

func (sf *FieldFilter) SetConfig(ctx context.Context, config interfaces.MarshableConfigI) error {
	var ok bool
	if sf.config, ok = config.(*FieldFilterConfig); !ok {
		return mperror.ErrFilterConfigInvalid
	}
	if sf.config.Op > FieldFilterOpNoOp {
		return mperror.ErrFilterConfigInvalid
	}
	return nil
}
func (sf *FieldFilter) GetConfig(ctx context.Context) (interfaces.MarshableConfigI, error) {
	return sf.config, nil
}

var _ interfaces.FilterImplI = (*FieldFilter)(nil)
