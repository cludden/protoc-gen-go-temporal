package scheme

import (
	"errors"
	"sync"

	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/dynamicpb"
)

// Option describes a functional configuration option for a Scheme value
type Option func(*Scheme)

// From can initialize a new Scheme value from one or more existing Scheme values
func From(schemes ...*Scheme) Option {
	return func(s *Scheme) {
		s.Merge(schemes...)
	}
}

// Scheme implements a protobuf type registry that can be used to initialize
// dynamic messages by name
type Scheme struct {
	m     sync.Mutex
	types map[string]protoreflect.MessageDescriptor
}

// New initializes a new Scheme value. A given Scheme value can include
// multiple service type registrations.
func New(opts ...Option) *Scheme {
	s := &Scheme{
		types: map[string]protoreflect.MessageDescriptor{},
	}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

// RegisterType registers the specified message with the current Scheme
func (s *Scheme) RegisterType(desc protoreflect.MessageDescriptor) {
	if s == nil {
		s = New()
	}
	s.m.Lock()
	defer s.m.Unlock()
	s.types[string(desc.FullName())] = desc
}

// New initializes a new proto message of the given type
func (s *Scheme) New(t string) (protoreflect.ProtoMessage, error) {
	if t, ok := s.types[t]; ok && t != nil {
		return dynamicpb.NewMessage(t).Interface(), nil
	}
	return nil, errors.New("not found")
}

// Merge can be used to merge multiple schemes into one
func (s *Scheme) Merge(schemes ...*Scheme) {
	if s == nil {
		s = &Scheme{}
	}
	for _, scheme := range schemes {
		scheme.m.Lock()
		defer scheme.m.Unlock()
		for _, t := range scheme.types {
			s.RegisterType(t)
		}
	}
}
