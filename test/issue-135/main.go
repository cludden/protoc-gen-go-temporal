package main

import (
	"context"

	issue_135v1 "github.com/cludden/protoc-gen-go-temporal/gen/test/issue-135/v1"
)

type Activities struct{}

func (a *Activities) Do(
	ctx context.Context,
	req *issue_135v1.DoRequest,
) (*issue_135v1.DoResponse, error) {
	return &issue_135v1.DoResponse{}, nil
}
