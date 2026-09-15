package xns

import (
	"errors"
	"math"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	commonpb "go.temporal.io/api/common/v1"
	deploymentpb "go.temporal.io/api/deployment/v1"
	"go.temporal.io/api/enums/v1"
	workflowpb "go.temporal.io/api/workflow/v1"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/converter"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/worker"
)

func TestErrorToApplicationError(t *testing.T) {
	require := require.New(t)

	err := ErrorToApplicationError(nil)
	require.Nil(err)

	err = errors.New("uh oh")
	require.ErrorIs(ErrorToApplicationError(err), err)

	err = temporal.NewNonRetryableApplicationError("uh oh", "Foo", nil)
	require.ErrorIs(ErrorToApplicationError(err), err)

	err = &temporal.WorkflowExecutionError{}
	require.NotNil(Unwrap(ErrorToApplicationError(err)))
	err = ErrorToApplicationError(err)
	require.Equal("WorkflowExecutionError", Code(err))
	require.True(IsNonRetryable(err))

	err = &temporal.CanceledError{}
	require.NotNil(Unwrap(ErrorToApplicationError(err)))
	err = ErrorToApplicationError(err)
	require.Equal("CanceledError", Code(err))
	require.True(IsNonRetryable(err))

	err = &temporal.TerminatedError{}
	require.NotNil(Unwrap(ErrorToApplicationError(err)))
	err = ErrorToApplicationError(err)
	require.Equal("TerminatedError", Code(err))
	require.True(IsNonRetryable(err))

	err = &temporal.ChildWorkflowExecutionError{}
	require.NotNil(Unwrap(ErrorToApplicationError(err)))
	err = ErrorToApplicationError(err)
	require.Equal("ChildWorkflowExecutionError", Code(err))
	require.True(IsNonRetryable(err))
}

func TestStartWorkflowOptionsRoundTrip(t *testing.T) {
	cases := map[string]client.StartWorkflowOptions{
		"empty": {},
		"pinned versioning override": {
			VersioningOverride: &client.PinnedVersioningOverride{
				Version: worker.WorkerDeploymentVersion{
					DeploymentName: "checkout",
					BuildID:        "abc123",
				},
			},
		},
		"auto upgrade versioning override": {
			VersioningOverride: &client.AutoUpgradeVersioningOverride{},
		},
		"all supported fields": {
			ID:                                       "workflow-id",
			TaskQueue:                                "task-queue",
			WorkflowExecutionTimeout:                 time.Hour,
			WorkflowRunTimeout:                       time.Minute * 30,
			WorkflowTaskTimeout:                      time.Second * 10,
			WorkflowIDConflictPolicy:                 enums.WORKFLOW_ID_CONFLICT_POLICY_USE_EXISTING,
			WorkflowIDReusePolicy:                    enums.WORKFLOW_ID_REUSE_POLICY_REJECT_DUPLICATE,
			WorkflowExecutionErrorWhenAlreadyStarted: true,
			RetryPolicy: &temporal.RetryPolicy{
				BackoffCoefficient:     2,
				InitialInterval:        time.Second,
				MaximumAttempts:        5,
				MaximumInterval:        time.Minute,
				NonRetryableErrorTypes: []string{"Fatal"},
			},
			CronSchedule:     "0 * * * *",
			Memo:             map[string]any{"foo": "bar"},
			EnableEagerStart: true,
			StartDelay:       time.Second * 5,
			StaticSummary:    "a summary",
			StaticDetails:    "some details",
			Priority: temporal.Priority{
				PriorityKey:    3,
				FairnessKey:    "tenant-a",
				FairnessWeight: 1.5,
			},
			TypedSearchAttributes: temporal.NewSearchAttributes(
				temporal.NewSearchAttributeKeyString("text").ValueSet("hello"),
				temporal.NewSearchAttributeKeyKeyword("keyword").ValueSet("kw"),
				temporal.NewSearchAttributeKeyBool("bool").ValueSet(true),
				temporal.NewSearchAttributeKeyInt64("int").ValueSet(1<<62),
				temporal.NewSearchAttributeKeyFloat64("float").ValueSet(1.25),
				temporal.NewSearchAttributeKeyTime("time").ValueSet(time.Date(2026, 9, 15, 12, 0, 0, 0, time.UTC)),
				temporal.NewSearchAttributeKeyKeywordList("list").ValueSet([]string{"a", "b"}),
			),
			VersioningOverride: &client.PinnedVersioningOverride{
				Version: worker.WorkerDeploymentVersion{
					DeploymentName: "checkout",
					BuildID:        "abc123",
				},
			},
		},
		"untyped search attributes": {
			SearchAttributes: map[string]any{"foo": "bar"},
		},
	}

	for name, expected := range cases {
		t.Run(name, func(t *testing.T) {
			pb, err := MarshalStartWorkflowOptions(expected)
			require.NoError(t, err)
			require.Equal(t, expected, UnmarshalStartWorkflowOptions(pb))
		})
	}
}

func TestMarshalVersioningOverride(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		vo, err := marshalVersioningOverride(nil)
		require.NoError(t, err)
		require.Nil(t, vo)
	})

	t.Run("typed nil", func(t *testing.T) {
		// a non-nil interface holding a nil pointer must not panic
		var pinned *client.PinnedVersioningOverride
		vo, err := marshalVersioningOverride(pinned)
		require.NoError(t, err)
		require.Nil(t, vo)
	})

	t.Run("auto upgrade sets only the override oneof", func(t *testing.T) {
		vo, err := marshalVersioningOverride(&client.AutoUpgradeVersioningOverride{})
		require.NoError(t, err)
		require.True(t, vo.GetAutoUpgrade())
		// the deprecated fields stay empty: nothing on this wire is read by a server
		require.Equal(t, enums.VERSIONING_BEHAVIOR_UNSPECIFIED, vo.GetBehavior()) //nolint:staticcheck
		require.Empty(t, vo.GetPinnedVersion())                                   //nolint:staticcheck
		require.Nil(t, vo.GetDeployment())                                        //nolint:staticcheck
	})

	t.Run("pinned sets the override oneof", func(t *testing.T) {
		vo, err := marshalVersioningOverride(&client.PinnedVersioningOverride{
			Version: worker.WorkerDeploymentVersion{DeploymentName: "checkout", BuildID: "abc123"},
		})
		require.NoError(t, err)
		require.Equal(t, workflowpb.VersioningOverride_PINNED_OVERRIDE_BEHAVIOR_PINNED, vo.GetPinned().GetBehavior())
		require.Equal(t, "checkout", vo.GetPinned().GetVersion().GetDeploymentName())
		require.Equal(t, "abc123", vo.GetPinned().GetVersion().GetBuildId())
	})
}

func TestUnmarshalVersioningOverride(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		require.Nil(t, unmarshalVersioningOverride(nil))
	})

	t.Run("auto upgrade behavior", func(t *testing.T) {
		require.Equal(t, &client.AutoUpgradeVersioningOverride{}, unmarshalVersioningOverride(&workflowpb.VersioningOverride{
			Behavior: enums.VERSIONING_BEHAVIOR_AUTO_UPGRADE,
		}))
	})

	t.Run("pinned deployment", func(t *testing.T) {
		require.Equal(t, &client.PinnedVersioningOverride{
			Version: worker.WorkerDeploymentVersion{DeploymentName: "checkout", BuildID: "abc123"},
		}, unmarshalVersioningOverride(&workflowpb.VersioningOverride{
			Behavior:   enums.VERSIONING_BEHAVIOR_PINNED,
			Deployment: &deploymentpb.Deployment{SeriesName: "checkout", BuildId: "abc123"},
		}))
	})

	t.Run("pinned version string", func(t *testing.T) {
		require.Equal(t, &client.PinnedVersioningOverride{
			Version: worker.WorkerDeploymentVersion{DeploymentName: "checkout", BuildID: "abc123"},
		}, unmarshalVersioningOverride(&workflowpb.VersioningOverride{
			Behavior:      enums.VERSIONING_BEHAVIOR_PINNED,
			PinnedVersion: "checkout.abc123",
		}))
	})

	t.Run("empty override", func(t *testing.T) {
		require.Nil(t, unmarshalVersioningOverride(&workflowpb.VersioningOverride{}))
	})

	t.Run("auto upgrade set to false", func(t *testing.T) {
		require.Nil(t, unmarshalVersioningOverride(&workflowpb.VersioningOverride{
			Override: &workflowpb.VersioningOverride_AutoUpgrade{AutoUpgrade: false},
		}))
	})

	t.Run("pinned version string without a separator", func(t *testing.T) {
		require.Nil(t, unmarshalVersioningOverride(&workflowpb.VersioningOverride{
			Behavior:      enums.VERSIONING_BEHAVIOR_PINNED,
			PinnedVersion: "checkout",
		}))
	})

	t.Run("override oneof wins over the deprecated fields", func(t *testing.T) {
		require.Equal(t, &client.PinnedVersioningOverride{
			Version: worker.WorkerDeploymentVersion{DeploymentName: "checkout", BuildID: "abc123"},
		}, unmarshalVersioningOverride(&workflowpb.VersioningOverride{
			Behavior:      enums.VERSIONING_BEHAVIOR_AUTO_UPGRADE,
			PinnedVersion: "other.def456",
			Override: &workflowpb.VersioningOverride_Pinned{
				Pinned: &workflowpb.VersioningOverride_PinnedOverride{
					Behavior: workflowpb.VersioningOverride_PINNED_OVERRIDE_BEHAVIOR_PINNED,
					Version: &deploymentpb.WorkerDeploymentVersion{
						DeploymentName: "checkout",
						BuildId:        "abc123",
					},
				},
			},
		}))
	})
}

func TestMarshalStartWorkflowOptionsPriorityOverflow(t *testing.T) {
	// PriorityKey is an int in the SDK but an int32 on the wire
	_, err := MarshalStartWorkflowOptions(client.StartWorkflowOptions{
		Priority: temporal.Priority{PriorityKey: math.MaxInt32 + 1},
	})
	require.ErrorContains(t, err, "priority key")
}

func TestTypedSearchAttributesRoundTrip(t *testing.T) {
	ts := time.Date(2026, 9, 15, 12, 0, 0, 0, time.UTC)
	cases := map[string]struct {
		attr      temporal.SearchAttributeUpdate
		valueType enums.IndexedValueType
		value     any
	}{
		"bool":         {temporal.NewSearchAttributeKeyBool("b").ValueSet(true), enums.INDEXED_VALUE_TYPE_BOOL, true},
		"float64":      {temporal.NewSearchAttributeKeyFloat64("f").ValueSet(1.25), enums.INDEXED_VALUE_TYPE_DOUBLE, 1.25},
		"int64":        {temporal.NewSearchAttributeKeyInt64("i").ValueSet(1 << 62), enums.INDEXED_VALUE_TYPE_INT, int64(1 << 62)},
		"keyword":      {temporal.NewSearchAttributeKeyKeyword("k").ValueSet("kw"), enums.INDEXED_VALUE_TYPE_KEYWORD, "kw"},
		"keyword list": {temporal.NewSearchAttributeKeyKeywordList("kl").ValueSet([]string{"a", "b"}), enums.INDEXED_VALUE_TYPE_KEYWORD_LIST, []string{"a", "b"}},
		"text":         {temporal.NewSearchAttributeKeyString("s").ValueSet("hello"), enums.INDEXED_VALUE_TYPE_TEXT, "hello"},
		"time":         {temporal.NewSearchAttributeKeyTime("t").ValueSet(ts), enums.INDEXED_VALUE_TYPE_DATETIME, ts},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			pb, err := marshalTypedSearchAttributes(temporal.NewSearchAttributes(c.attr))
			require.NoError(t, err)
			require.Len(t, pb.GetIndexedFields(), 1)

			// the indexed value type travels with the payload, not the key name
			for _, payload := range pb.GetIndexedFields() {
				require.Equal(t, c.valueType.String(), string(payload.GetMetadata()["type"]))
			}

			got := unmarshalTypedSearchAttributes(pb)
			require.Equal(t, 1, got.Size())
			for key, value := range got.GetUntypedValues() {
				require.Equal(t, c.valueType, key.GetValueType())
				require.Equal(t, c.value, value)
			}
		})
	}
}

func TestUnmarshalTypedSearchAttributesSkipsUnusable(t *testing.T) {
	dc := converter.GetDefaultDataConverter()
	payload := func(value any, valueType string) *commonpb.Payload {
		p, err := dc.ToPayload(value)
		require.NoError(t, err)
		if valueType != "" {
			p.Metadata["type"] = []byte(valueType)
		}
		return p
	}

	got := unmarshalTypedSearchAttributes(&commonpb.SearchAttributes{
		IndexedFields: map[string]*commonpb.Payload{
			"good":     payload("hello", enums.INDEXED_VALUE_TYPE_TEXT.String()),
			"no type":  payload("hello", ""),
			"unknown":  payload("hello", "NotAnIndexedValueType"),
			"mistyped": payload("hello", enums.INDEXED_VALUE_TYPE_INT.String()),
			"no data": {
				Metadata: map[string][]byte{"type": []byte(enums.INDEXED_VALUE_TYPE_TEXT.String())},
			},
		},
	})

	require.Equal(t, 1, got.Size())
	value, ok := got.GetString(temporal.NewSearchAttributeKeyString("good"))
	require.True(t, ok)
	require.Equal(t, "hello", value)
}
