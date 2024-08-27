package ctxpropagation

import (
	"context"
	"fmt"
	"reflect"

	"github.com/cludden/protoc-gen-go-temporal/pkg/ctxkey"
	"go.temporal.io/sdk/converter"
	"go.temporal.io/sdk/workflow"
)

type (
	Key interface {
		Value() reflect.Value
		Serializable() bool
		String() string
	}

	propagator struct {
		dc   converter.DataConverter
		keys []Key
	}

	Option func(*propagator)
)

func New(opts ...Option) workflow.ContextPropagator {
	p := &propagator{}
	for _, opt := range opts {
		opt(p)
	}
	if p.dc == nil {
		p.dc = converter.GetDefaultDataConverter()
	}
	return p
}

func (p *propagator) Extract(ctx context.Context, r workflow.HeaderReader) (context.Context, error) {
	for _, k := range p.keys {
		if k.Serializable() {
			payload, ok := r.Get(k.String())
			if !ok {
				continue
			}

			val := k.Value()
			var canAddr bool
			if val.CanAddr() {
				val, canAddr = val.Addr(), true
			}
			if err := p.dc.FromPayload(payload, val.Interface()); err != nil {
				return ctx, fmt.Errorf("error decoding payload to value of type %s", val.Type().String())
			}

			v := val.Interface()
			if canAddr {
				v = val.Elem().Interface()
			}
			ctx = context.WithValue(ctx, k, v)
		}
	}
	return ctx, nil
}

func (p *propagator) ExtractToWorkflow(ctx workflow.Context, r workflow.HeaderReader) (workflow.Context, error) {
	for _, k := range p.keys {
		if k.Serializable() {
			payload, ok := r.Get(k.String())
			if !ok {
				continue
			}

			val := k.Value()
			var canAddr bool
			if val.CanAddr() {
				val, canAddr = val.Addr(), true
			}

			if err := p.dc.FromPayload(payload, val.Interface()); err != nil {
				return ctx, fmt.Errorf("error decoding payload to value of type %s", val.Type().String())
			}

			v := val.Interface()
			if canAddr {
				v = val.Elem().Interface()
			}
			ctx = workflow.WithValue(ctx, k, v)
		}
	}
	return ctx, nil
}

func (p *propagator) Inject(ctx context.Context, w workflow.HeaderWriter) error {
	return p.inject(ctx, w)
}

func (p *propagator) InjectFromWorkflow(ctx workflow.Context, w workflow.HeaderWriter) error {
	return p.inject(ctx, w)
}

func WithDataConverter(dc converter.DataConverter) Option {
	return func(p *propagator) {
		p.dc = dc
	}
}

func WithKeys(keys ...Key) Option {
	return func(p *propagator) {
		p.keys = append(p.keys, keys...)
	}
}

func (p *propagator) inject(ctx ctxkey.Context, w workflow.HeaderWriter) error {
	for _, k := range p.keys {
		if k.Serializable() {
			payload, err := p.dc.ToPayload(ctx.Value(k))
			if err != nil {
				return err
			}
			w.Set(k.String(), payload)
		}
	}
	return nil
}
