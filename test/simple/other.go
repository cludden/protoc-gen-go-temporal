package main

import (
	simplepb "github.com/cludden/protoc-gen-go-temporal/gen/test/simple/v1"
)

type (
	OtherWorkflows struct {
		simplepb.OtherWorkflows
	}

	OtherActivities struct {
		simplepb.OtherActivities
	}
)
