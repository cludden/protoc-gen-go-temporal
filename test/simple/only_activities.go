package main

import (
	"context"

	simplepb "github.com/cludden/protoc-gen-go-temporal/gen/simple"
)

type OnlyActivites struct{}

func (a *OnlyActivites) LonelyActivity1(ctx context.Context, req *simplepb.LonelyActivity1Request) (*simplepb.LonelyActivity1Response, error) {
	return &simplepb.LonelyActivity1Response{}, nil
}
