package convert

import (
	temporalv1 "github.com/cludden/protoc-gen-go-temporal/gen/temporal/v1"
	"go.temporal.io/api/enums/v1"
)

func FromParentClosePolicy(p temporalv1.ParentClosePolicy) enums.ParentClosePolicy {
	switch p {
	case temporalv1.ParentClosePolicy_PARENT_CLOSE_POLICY_ABANDON:
		return enums.PARENT_CLOSE_POLICY_ABANDON
	case temporalv1.ParentClosePolicy_PARENT_CLOSE_POLICY_REQUEST_CANCEL:
		return enums.PARENT_CLOSE_POLICY_REQUEST_CANCEL
	case temporalv1.ParentClosePolicy_PARENT_CLOSE_POLICY_TERMINATE:
		return enums.PARENT_CLOSE_POLICY_TERMINATE
	default:
		return enums.PARENT_CLOSE_POLICY_UNSPECIFIED
	}
}

func ToParentClosePolicy(p enums.ParentClosePolicy) temporalv1.ParentClosePolicy {
	switch p {
	case enums.PARENT_CLOSE_POLICY_ABANDON:
		return temporalv1.ParentClosePolicy_PARENT_CLOSE_POLICY_ABANDON
	case enums.PARENT_CLOSE_POLICY_REQUEST_CANCEL:
		return temporalv1.ParentClosePolicy_PARENT_CLOSE_POLICY_REQUEST_CANCEL
	case enums.PARENT_CLOSE_POLICY_TERMINATE:
		return temporalv1.ParentClosePolicy_PARENT_CLOSE_POLICY_TERMINATE
	default:
		return temporalv1.ParentClosePolicy_PARENT_CLOSE_POLICY_UNSPECIFIED
	}
}
