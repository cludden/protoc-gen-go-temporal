package main

import (
	"fmt"
	"testing"
	"time"

	optionv1 "github.com/cludden/protoc-gen-go-temporal/gen/test/option/v1"
	"github.com/stretchr/testify/suite"
	"go.temporal.io/api/enums/v1"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/testsuite"
	"go.temporal.io/sdk/workflow"
	"google.golang.org/protobuf/proto"
)

type OptionSuite struct {
	suite.Suite
	testsuite.WorkflowTestSuite
	env    *testsuite.TestWorkflowEnvironment
	client optionv1.TestClient
}

func TestOptionSuite(t *testing.T) {
	suite.Run(t, new(OptionSuite))
}

func (s *OptionSuite) SetupTest() {
	s.env = s.NewTestWorkflowEnvironment()
	s.client = optionv1.NewTestTestClient(s.env, &Workflows{}, &Activities{})
}

func (s *OptionSuite) SetupSubTest() {
	s.SetupTest()
}

func (s *OptionSuite) TestActivityOptions() {
	cases := []struct {
		desc     string
		options  []*optionv1.ActivityWithInputActivityOptions
		expected workflow.ActivityOptions
	}{
		{
			desc: "defaults",
			expected: workflow.ActivityOptions{
				HeartbeatTimeout: time.Second * 30,
				RetryPolicy: &temporal.RetryPolicy{
					MaximumInterval: time.Second * 5,
				},
				ScheduleToCloseTimeout: time.Second * 120,
				ScheduleToStartTimeout: time.Second * 10,
				StartToCloseTimeout:    time.Second * 60,
				TaskQueue:              "option-v2",
				WaitForCancellation:    true,
			},
			options: []*optionv1.ActivityWithInputActivityOptions{
				optionv1.NewActivityWithInputActivityOptions(),
				optionv1.NewActivityWithInputActivityOptions().
					WithActivityOptions(workflow.ActivityOptions{}),
			},
		},
		{
			desc: "HeartbeatTimeout",
			expected: workflow.ActivityOptions{
				HeartbeatTimeout: time.Second * 29,
				RetryPolicy: &temporal.RetryPolicy{
					MaximumInterval: time.Second * 5,
				},
				ScheduleToCloseTimeout: time.Second * 120,
				ScheduleToStartTimeout: time.Second * 10,
				StartToCloseTimeout:    time.Second * 60,
				TaskQueue:              "option-v2",
				WaitForCancellation:    true,
			},
			options: []*optionv1.ActivityWithInputActivityOptions{
				// override with activity options
				optionv1.NewActivityWithInputActivityOptions().
					WithActivityOptions(workflow.ActivityOptions{
						HeartbeatTimeout: time.Second * 29,
					}),
				// last activity option override wins
				optionv1.NewActivityWithInputActivityOptions().
					WithActivityOptions(workflow.ActivityOptions{
						HeartbeatTimeout: time.Second * 31,
					}).
					WithActivityOptions(workflow.ActivityOptions{
						HeartbeatTimeout: time.Second * 29,
					}),
				// override with field-specific option
				optionv1.NewActivityWithInputActivityOptions().
					WithHeartbeatTimeout(time.Second * 29),
				// field-specific option takes precedence over activity option
				optionv1.NewActivityWithInputActivityOptions().
					WithHeartbeatTimeout(time.Second * 29).
					WithActivityOptions(workflow.ActivityOptions{
						HeartbeatTimeout: time.Second * 31,
					}),
				// field-specific option takes precedence regardless of calling order
				optionv1.NewActivityWithInputActivityOptions().
					WithActivityOptions(workflow.ActivityOptions{
						HeartbeatTimeout: time.Second * 31,
					}).
					WithHeartbeatTimeout(time.Second * 29),
				// last field-specific option wins
				optionv1.NewActivityWithInputActivityOptions().
					WithHeartbeatTimeout(time.Second * 32).
					WithActivityOptions(workflow.ActivityOptions{
						HeartbeatTimeout: time.Second * 31,
					}).
					WithHeartbeatTimeout(time.Second * 29),
			},
		},
		{
			desc: "ScheduleToCloseTimeout",
			expected: workflow.ActivityOptions{
				HeartbeatTimeout: time.Second * 30,
				RetryPolicy: &temporal.RetryPolicy{
					MaximumInterval: time.Second * 5,
				},
				ScheduleToCloseTimeout: time.Second * 119,
				ScheduleToStartTimeout: time.Second * 10,
				StartToCloseTimeout:    time.Second * 60,
				TaskQueue:              "option-v2",
				WaitForCancellation:    true,
			},
			options: []*optionv1.ActivityWithInputActivityOptions{
				// override with activity options
				optionv1.NewActivityWithInputActivityOptions().
					WithActivityOptions(workflow.ActivityOptions{
						ScheduleToCloseTimeout: time.Second * 119,
					}),
				// last activity option override wins
				optionv1.NewActivityWithInputActivityOptions().
					WithActivityOptions(workflow.ActivityOptions{
						ScheduleToCloseTimeout: time.Second * 121,
					}).
					WithActivityOptions(workflow.ActivityOptions{
						ScheduleToCloseTimeout: time.Second * 119,
					}),
				// override with field-specific option
				optionv1.NewActivityWithInputActivityOptions().
					WithScheduleToCloseTimeout(time.Second * 119),
				// field-specific option takes precedence over activity option
				optionv1.NewActivityWithInputActivityOptions().
					WithScheduleToCloseTimeout(time.Second * 119).
					WithActivityOptions(workflow.ActivityOptions{
						ScheduleToCloseTimeout: time.Second * 121,
					}),
				// field-specific option takes precedence regardless of calling order
				optionv1.NewActivityWithInputActivityOptions().
					WithActivityOptions(workflow.ActivityOptions{
						ScheduleToCloseTimeout: time.Second * 121,
					}).
					WithScheduleToCloseTimeout(time.Second * 119),
				// last field-specific option wins
				optionv1.NewActivityWithInputActivityOptions().
					WithScheduleToCloseTimeout(time.Second * 122).
					WithActivityOptions(workflow.ActivityOptions{
						ScheduleToCloseTimeout: time.Second * 121,
					}).
					WithScheduleToCloseTimeout(time.Second * 119),
			},
		},
		{
			desc: "override with falsey value",
			expected: workflow.ActivityOptions{
				RetryPolicy: &temporal.RetryPolicy{},
				TaskQueue:   "default-test-taskqueue",
			},
			options: []*optionv1.ActivityWithInputActivityOptions{
				// override with activity options
				optionv1.NewActivityWithInputActivityOptions().
					WithHeartbeatTimeout(0).
					WithRetryPolicy(&temporal.RetryPolicy{}).
					WithScheduleToCloseTimeout(0).
					WithScheduleToStartTimeout(0).
					WithStartToCloseTimeout(0).
					WithTaskQueue("").
					WithWaitForCancellation(false),
			},
		},
	}

	for _, c := range cases {
		for i, opts := range c.options {
			s.Run(fmt.Sprintf("%s-%d", c.desc, i), func() {
				var actual workflow.ActivityOptions
				s.env.ExecuteWorkflow(func(ctx workflow.Context) (err error) {
					ctx, err = opts.Build(ctx)
					actual = workflow.GetActivityOptions(ctx)
					return err
				})
				s.Require().True(s.env.IsWorkflowCompleted())
				s.Require().NoError(s.env.GetWorkflowError())
				s.Require().Equal(c.expected, actual)
			})
		}
	}
}

func (s *OptionSuite) TestChildWorkflowOptions() {
	cases := []struct {
		desc     string
		expected workflow.ChildWorkflowOptions
		options  []*optionv1.WorkflowWithInputChildOptions
	}{
		{
			desc: "defaults",
			expected: workflow.ChildWorkflowOptions{
				ParentClosePolicy: enums.PARENT_CLOSE_POLICY_REQUEST_CANCEL,
				RetryPolicy: &temporal.RetryPolicy{
					MaximumAttempts: 5,
				},
				SearchAttributes: map[string]any{
					"name": "foo",
				},
				TaskQueue:             "option-v2",
				WaitForCancellation:   true,
				WorkflowID:            "workflow-with-input:foo",
				WorkflowIDReusePolicy: enums.WORKFLOW_ID_REUSE_POLICY_ALLOW_DUPLICATE_FAILED_ONLY,

				WorkflowExecutionTimeout: time.Second * 600,
				WorkflowRunTimeout:       time.Second * 300,
				WorkflowTaskTimeout:      time.Second * 10,
			},
			options: []*optionv1.WorkflowWithInputChildOptions{
				optionv1.NewWorkflowWithInputChildOptions(),
			},
		},
		{
			desc: "overrides",
			expected: workflow.ChildWorkflowOptions{
				ParentClosePolicy: enums.PARENT_CLOSE_POLICY_TERMINATE,
				RetryPolicy: &temporal.RetryPolicy{
					MaximumAttempts: 4,
				},
				SearchAttributes: map[string]any{
					"name": "bar",
				},
				TaskQueue:                "option-v3",
				WaitForCancellation:      true,
				WorkflowID:               "foo",
				WorkflowIDReusePolicy:    enums.WORKFLOW_ID_REUSE_POLICY_REJECT_DUPLICATE,
				WorkflowExecutionTimeout: time.Second * 599,
				WorkflowRunTimeout:       time.Second * 299,
				WorkflowTaskTimeout:      time.Second * 9,
			},
			options: []*optionv1.WorkflowWithInputChildOptions{
				optionv1.NewWorkflowWithInputChildOptions().
					WithChildWorkflowOptions(workflow.ChildWorkflowOptions{
						ParentClosePolicy: enums.PARENT_CLOSE_POLICY_TERMINATE,
						RetryPolicy: &temporal.RetryPolicy{
							MaximumAttempts: 4,
						},
						SearchAttributes: map[string]any{
							"name": "bar",
						},
						TaskQueue:                "option-v3",
						WaitForCancellation:      false,
						WorkflowID:               "foo",
						WorkflowIDReusePolicy:    enums.WORKFLOW_ID_REUSE_POLICY_REJECT_DUPLICATE,
						WorkflowExecutionTimeout: time.Second * 599,
						WorkflowRunTimeout:       time.Second * 299,
						WorkflowTaskTimeout:      time.Second * 9,
					}),
			},
		},
		{
			desc: "field-specific",
			expected: workflow.ChildWorkflowOptions{
				ParentClosePolicy: enums.PARENT_CLOSE_POLICY_TERMINATE,
				RetryPolicy: &temporal.RetryPolicy{
					MaximumAttempts: 4,
				},
				SearchAttributes: map[string]any{
					"name": "bar",
				},
				TaskQueue:                "option-v3",
				WaitForCancellation:      false,
				WorkflowID:               "foo",
				WorkflowIDReusePolicy:    enums.WORKFLOW_ID_REUSE_POLICY_REJECT_DUPLICATE,
				WorkflowExecutionTimeout: time.Second * 599,
				WorkflowRunTimeout:       time.Second * 299,
				WorkflowTaskTimeout:      time.Second * 9,
			},
			options: []*optionv1.WorkflowWithInputChildOptions{
				optionv1.NewWorkflowWithInputChildOptions().
					WithExecutionTimeout(time.Second * 599).
					WithID("foo").
					WithIDReusePolicy(enums.WORKFLOW_ID_REUSE_POLICY_REJECT_DUPLICATE).
					WithParentClosePolicy(enums.PARENT_CLOSE_POLICY_TERMINATE).
					WithRetryPolicy(&temporal.RetryPolicy{
						MaximumAttempts: 4,
					}).
					WithRunTimeout(time.Second * 299).
					WithSearchAttributes(map[string]any{
						"name": "bar",
					}).
					WithTaskQueue("option-v3").
					WithTaskTimeout(time.Second * 9).
					WithWaitForCancellation(false).
					WithChildWorkflowOptions(workflow.ChildWorkflowOptions{
						ParentClosePolicy: enums.PARENT_CLOSE_POLICY_ABANDON,
						RetryPolicy: &temporal.RetryPolicy{
							MaximumAttempts: 6,
						},
						SearchAttributes: map[string]any{
							"name": "baz",
						},
						TaskQueue:                "option-v2",
						WaitForCancellation:      false,
						WorkflowID:               "bar",
						WorkflowIDReusePolicy:    enums.WORKFLOW_ID_REUSE_POLICY_TERMINATE_IF_RUNNING,
						WorkflowExecutionTimeout: time.Second * 601,
						WorkflowRunTimeout:       time.Second * 301,
						WorkflowTaskTimeout:      time.Second * 11,
					}),
			},
		},
	}

	for _, c := range cases {
		for i, opts := range c.options {
			s.Run(fmt.Sprintf("%s-%d", c.desc, i), func() {
				var actual workflow.ChildWorkflowOptions
				s.env.ExecuteWorkflow(func(ctx workflow.Context) (err error) {
					actual, err = opts.Build(ctx, (&optionv1.WorkflowWithInputRequest{Name: "foo"}).ProtoReflect())
					return err
				})
				s.Require().True(s.env.IsWorkflowCompleted())
				s.Require().NoError(s.env.GetWorkflowError())
				s.Require().Equal(c.expected, actual)
			})
		}
	}

}

func (s *OptionSuite) TestLocalActivityOptions() {
	cases := []struct {
		desc     string
		expected workflow.LocalActivityOptions
		options  []*optionv1.ActivityWithInputLocalActivityOptions
	}{
		{
			desc: "defaults",
			expected: workflow.LocalActivityOptions{
				ScheduleToCloseTimeout: time.Second * 120,
				StartToCloseTimeout:    time.Second * 60,
				RetryPolicy: &temporal.RetryPolicy{
					BackoffCoefficient: 2,           // set by sdk.applyRetryPolicyDefaultsForLocalActivity
					InitialInterval:    time.Second, // set by sdk.applyRetryPolicyDefaultsForLocalActivity
					MaximumInterval:    time.Second * 5,
				},
			},
			options: []*optionv1.ActivityWithInputLocalActivityOptions{
				optionv1.NewActivityWithInputLocalActivityOptions(),
				optionv1.NewActivityWithInputLocalActivityOptions().
					WithLocalActivityOptions(workflow.LocalActivityOptions{}),
			},
		},
		{
			desc: "overrides",
			expected: workflow.LocalActivityOptions{
				ScheduleToCloseTimeout: time.Second * 119,
				StartToCloseTimeout:    time.Second * 59,
				RetryPolicy: &temporal.RetryPolicy{
					BackoffCoefficient: 2,           // set by sdk.applyRetryPolicyDefaultsForLocalActivity
					InitialInterval:    time.Second, // set by sdk.applyRetryPolicyDefaultsForLocalActivity
					MaximumInterval:    time.Second * 4,
				},
			},
			options: []*optionv1.ActivityWithInputLocalActivityOptions{
				optionv1.NewActivityWithInputLocalActivityOptions().
					WithLocalActivityOptions(workflow.LocalActivityOptions{
						ScheduleToCloseTimeout: time.Second * 119,
						StartToCloseTimeout:    time.Second * 59,
						RetryPolicy: &temporal.RetryPolicy{
							MaximumInterval: time.Second * 4,
						},
					}),
			},
		},
		{
			desc: "field-specific",
			expected: workflow.LocalActivityOptions{
				ScheduleToCloseTimeout: time.Second * 119,
				StartToCloseTimeout:    time.Second * 59,
				RetryPolicy: &temporal.RetryPolicy{
					BackoffCoefficient: 2,           // set by sdk.applyRetryPolicyDefaultsForLocalActivity
					InitialInterval:    time.Second, // set by sdk.applyRetryPolicyDefaultsForLocalActivity
					MaximumInterval:    time.Second * 4,
				},
			},
			options: []*optionv1.ActivityWithInputLocalActivityOptions{
				optionv1.NewActivityWithInputLocalActivityOptions().
					WithRetryPolicy(&temporal.RetryPolicy{
						MaximumInterval: time.Second * 4,
					}).
					WithScheduleToCloseTimeout(time.Second * 119).
					WithStartToCloseTimeout(time.Second * 59).
					WithLocalActivityOptions(workflow.LocalActivityOptions{
						ScheduleToCloseTimeout: time.Second * 121,
						StartToCloseTimeout:    time.Second * 61,
						RetryPolicy: &temporal.RetryPolicy{
							MaximumInterval: time.Second * 6,
						},
					}),
				optionv1.NewActivityWithInputLocalActivityOptions().
					WithLocalActivityOptions(workflow.LocalActivityOptions{
						ScheduleToCloseTimeout: time.Second * 121,
						StartToCloseTimeout:    time.Second * 61,
						RetryPolicy: &temporal.RetryPolicy{
							MaximumInterval: time.Second * 6,
						},
					}).
					WithRetryPolicy(&temporal.RetryPolicy{
						MaximumInterval: time.Second * 4,
					}).
					WithScheduleToCloseTimeout(time.Second * 119).
					WithStartToCloseTimeout(time.Second * 59),
				optionv1.NewActivityWithInputLocalActivityOptions().
					WithRetryPolicy(&temporal.RetryPolicy{
						MaximumInterval: time.Second * 6,
					}).
					WithScheduleToCloseTimeout(time.Second * 121).
					WithStartToCloseTimeout(time.Second * 61).
					WithRetryPolicy(&temporal.RetryPolicy{
						MaximumInterval: time.Second * 4,
					}).
					WithScheduleToCloseTimeout(time.Second * 119).
					WithStartToCloseTimeout(time.Second * 59),
			},
		},
	}

	for _, c := range cases {
		s.Require().Greater(len(c.options), 0, "invalid test case, at least 1 options case is required")
		for i, opts := range c.options {
			s.Run(fmt.Sprintf("%s-%d", c.desc, i), func() {
				var actual workflow.LocalActivityOptions
				s.env.ExecuteWorkflow(func(ctx workflow.Context) (err error) {
					ctx, err = opts.Build(ctx)
					actual = workflow.GetLocalActivityOptions(ctx)
					return err
				})
				s.Require().True(s.env.IsWorkflowCompleted())
				s.Require().NoError(s.env.GetWorkflowError())
				s.Require().Equal(c.expected, actual)
			})
		}
	}
}

func (s *OptionSuite) TestUpdateOptions() {
	cases := []struct {
		desc     string
		expected client.UpdateWorkflowOptions
		options  []*optionv1.UpdateWithInputOptions
	}{
		{
			desc: "defaults",
			expected: client.UpdateWorkflowOptions{
				Args: []any{
					&optionv1.UpdateWithInputRequest{
						Name: "bar",
					},
				},
				WorkflowID:   "foo",
				RunID:        "",
				UpdateID:     "update-with-input:bar",
				WaitForStage: client.WorkflowUpdateStageCompleted,
			},
			options: []*optionv1.UpdateWithInputOptions{
				optionv1.NewUpdateWithInputOptions(),
			},
		},
		{
			desc: "overrides",
			expected: client.UpdateWorkflowOptions{
				Args: []any{
					&optionv1.UpdateWithInputRequest{
						Name: "bar",
					},
				},
				WorkflowID:   "foo",
				RunID:        "",
				UpdateID:     "bar",
				WaitForStage: client.WorkflowUpdateStageAccepted,
			},
			options: []*optionv1.UpdateWithInputOptions{
				optionv1.NewUpdateWithInputOptions().
					WithUpdateWorkflowOptions(client.UpdateWorkflowOptions{
						WorkflowID: "FOO",
						RunID:      "BAZ",
					}).
					WithUpdateID("bar").
					WithWaitPolicy(client.WorkflowUpdateStageAccepted),
			},
		},
	}

	for _, c := range cases {
		for i, opts := range c.options {
			s.Run(fmt.Sprintf("%s-%d", c.desc, i), func() {
				actual, err := opts.Build("foo", "", &optionv1.UpdateWithInputRequest{
					Name: "bar",
				})
				s.Require().NoError(err)
				s.Require().NotNil(actual)
				s.Require().Equal(c.expected.WorkflowID, actual.WorkflowID)
				s.Require().Equal(c.expected.RunID, actual.RunID)
				s.Require().Equal(c.expected.UpdateID, actual.UpdateID)
				s.Require().Equal(c.expected.WaitForStage, actual.WaitForStage)
				s.Require().Len(actual.Args, len(c.expected.Args))
				for j, arg := range c.expected.Args {
					s.Require().True(proto.Equal(arg.(proto.Message), actual.Args[j].(proto.Message)))
				}
			})
		}
	}
}

func (s *OptionSuite) TestWorkflowOptions() {
	cases := []struct {
		desc     string
		options  []*optionv1.WorkflowWithInputOptions
		expected client.StartWorkflowOptions
	}{
		{
			desc: "defaults",
			expected: client.StartWorkflowOptions{
				EnableEagerStart: true,
				ID:               "workflow-with-input:foo",
				RetryPolicy: &temporal.RetryPolicy{
					MaximumAttempts: 5,
				},
				SearchAttributes: map[string]any{
					"name": "foo",
				},
				TaskQueue:                "option-v2",
				WorkflowIDConflictPolicy: enums.WORKFLOW_ID_CONFLICT_POLICY_USE_EXISTING,
				WorkflowIDReusePolicy:    enums.WORKFLOW_ID_REUSE_POLICY_ALLOW_DUPLICATE_FAILED_ONLY,
				WorkflowExecutionTimeout: time.Second * 600,
				WorkflowRunTimeout:       time.Second * 300,
				WorkflowTaskTimeout:      time.Second * 10,
			},
			options: []*optionv1.WorkflowWithInputOptions{
				optionv1.NewWorkflowWithInputOptions(),
				optionv1.NewWorkflowWithInputOptions().
					WithStartWorkflowOptions(client.StartWorkflowOptions{}),
			},
		},
		{
			desc: "overrides",
			expected: client.StartWorkflowOptions{
				ID: "foo",
				RetryPolicy: &temporal.RetryPolicy{
					MaximumAttempts: 4,
				},
				SearchAttributes: map[string]any{
					"name": "bar",
				},
				TaskQueue:                "option-v2",
				WorkflowIDReusePolicy:    enums.WORKFLOW_ID_REUSE_POLICY_REJECT_DUPLICATE,
				WorkflowIDConflictPolicy: enums.WORKFLOW_ID_CONFLICT_POLICY_FAIL,
				WorkflowExecutionTimeout: time.Second * 599,
				WorkflowRunTimeout:       time.Second * 299,
				WorkflowTaskTimeout:      time.Second * 9,
			},
			options: []*optionv1.WorkflowWithInputOptions{
				optionv1.NewWorkflowWithInputOptions().
					WithStartWorkflowOptions(client.StartWorkflowOptions{
						ID: "foo",
						RetryPolicy: &temporal.RetryPolicy{
							MaximumAttempts: 4,
						},
						SearchAttributes: map[string]any{
							"name": "bar",
						},
						TaskQueue:                "option-v2",
						WorkflowIDReusePolicy:    enums.WORKFLOW_ID_REUSE_POLICY_REJECT_DUPLICATE,
						WorkflowIDConflictPolicy: enums.WORKFLOW_ID_CONFLICT_POLICY_FAIL,
						WorkflowExecutionTimeout: time.Second * 599,
						WorkflowRunTimeout:       time.Second * 299,
						WorkflowTaskTimeout:      time.Second * 9,
					}).
					WithEnableEagerStart(false),
				optionv1.NewWorkflowWithInputOptions().
					WithEnableEagerStart(false).
					WithExecutionTimeout(time.Second * 599).
					WithID("foo").
					WithWorkflowIdConflictPolicy(enums.WORKFLOW_ID_CONFLICT_POLICY_FAIL).
					WithIDReusePolicy(enums.WORKFLOW_ID_REUSE_POLICY_REJECT_DUPLICATE).
					WithRetryPolicy(&temporal.RetryPolicy{
						MaximumAttempts: 4,
					}).
					WithRunTimeout(time.Second * 299).
					WithSearchAttributes(map[string]any{
						"name": "bar",
					}).
					WithTaskQueue("option-v2").
					WithTaskTimeout(time.Second * 9).
					WithStartWorkflowOptions(client.StartWorkflowOptions{
						ID: "bar",
						RetryPolicy: &temporal.RetryPolicy{
							MaximumAttempts: 6,
						},
						SearchAttributes: map[string]any{
							"name": "baz",
						},
						TaskQueue:                "option-v3",
						WorkflowIDReusePolicy:    enums.WORKFLOW_ID_REUSE_POLICY_ALLOW_DUPLICATE,
						WorkflowExecutionTimeout: time.Second * 601,
						WorkflowRunTimeout:       time.Second * 301,
						WorkflowTaskTimeout:      time.Second * 11,
					}),
			},
		},
	}

	for _, c := range cases {
		for i, opts := range c.options {
			s.Run(fmt.Sprintf("%s-%d", c.desc, i), func() {
				actual, err := opts.Build((&optionv1.WorkflowWithInputRequest{Name: "foo"}).ProtoReflect())
				s.Require().NoError(err)
				s.Require().Equal(c.expected, actual)
			})
		}
	}
}
