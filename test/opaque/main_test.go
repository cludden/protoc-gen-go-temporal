package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/cludden/protoc-gen-go-temporal/gen/test/opaque"
	clientmocks "github.com/cludden/protoc-gen-go-temporal/mocks/go.temporal.io/sdk/client"
	"github.com/cludden/protoc-gen-go-temporal/mocks/go.temporal.io/sdk/clientutils"
	"github.com/cludden/protoc-gen-go-temporal/pkg/convert"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/urfave/cli/v2"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/workflow"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/testing/protocmp"
)

func TestUnmarshalHybridExampleCliFlags(t *testing.T) {
	testCliFlags(t, "hybrid", "put-hybrid-example-with-signal-hybrid", opaque.NewHybridCli, opaque.NewHybridCliOptions())
}

func TestUnmarshalOpenExampleCliFlags(t *testing.T) {
	testCliFlags(t, "open", "put-open-example-with-signal-open", opaque.NewOpenCli, opaque.NewOpenCliOptions())
}

func TestUnmarshalOpaqueExampleCliFlags(t *testing.T) {
	testCliFlags(t, "opaque", "put-opaque-example-with-signal-opaque", opaque.NewOpaqueCli, opaque.NewOpaqueCliOptions())
}

func TestUnmarshalOptionalExampleCliFlags(t *testing.T) {
	testCliFlags(t, "optional", "put-optional-example-with-signal-optional", opaque.NewOptionalCli, opaque.NewOptionalCliOptions())
}

// =============================================================================

type testExampleMessage interface {
	GetName() string
	GetScore() float64
	GetScores() []float64
	GetRatio() float32
	GetRatios() []float32
	GetAge() int32
	GetAges() []int32
	GetId() int64
	GetIds() []int64
	GetEmails() []string
	GetExtra() map[string]string
	GetStatus() opaque.Status
	GetStatuses() []opaque.Status
	GetSize() uint32
	GetSizes() []uint32
	GetLength() uint64
	GetLengths() []uint64
	GetConnectionId() int32
	GetConnectionIds() []int32
	GetSessionId() int64
	GetSessionIds() []int64
	GetFixedSize() uint32
	GetFixedSizes() []uint32
	GetFixedLength() uint64
	GetFixedLengths() []uint64
	GetSfixedSize() int32
	GetSfixedSizes() []int32
	GetSfixedLength() int64
	GetSfixedLengths() []int64
	GetIsActive() bool
	GetIsActives() []bool
	GetData() []byte
	GetDatas() [][]byte
	GetAddress() *opaque.Address
	GetPreviousAddresses() []*opaque.Address
	GetOneofName() string
	GetOneofScore() float64
	GetOneofRatio() float32
	GetOneofAge() int32
	GetOneofId() int64
	GetOneofStatus() opaque.Status
	GetOneofSize() uint32
	GetOneofLength() uint64
	GetOneofConnectionId() int32
	GetOneofSessionId() int64
	GetOneofFixedSize() uint32
	GetOneofFixedLength() uint64
	GetOneofSfixedSize() int32
	GetOneofSfixedLength() int64
	GetOneofIsActive() bool
	GetOneofData() []byte
	GetOneofAddress() *opaque.Address
}

var testFlags = map[string]struct {
	args     []string
	expected any
	errors   []string
	get      func(o testExampleMessage) any
}{
	"name": {
		args:     []string{"--name", "foo"},
		expected: "foo",
		get:      func(o testExampleMessage) any { return o.GetName() },
	},
	"score": {
		args:     []string{"--score", "3.14"},
		expected: float64(3.14),
		get:      func(o testExampleMessage) any { return o.GetScore() },
	},
	"scores": {
		args:     []string{"--scores", "3.14", "--scores", "3.15"},
		expected: []float64{3.14, 3.15},
		get:      func(o testExampleMessage) any { return o.GetScores() },
	},
	"ratio": {
		args:     []string{"--ratio", "3.15"},
		expected: float32(3.15),
		get:      func(o testExampleMessage) any { return o.GetRatio() },
	},
	"ratios": {
		args:     []string{"--ratios", "3.15", "--ratios", "3.16"},
		expected: []float32{3.15, 3.16},
		get:      func(o testExampleMessage) any { return o.GetRatios() },
	},
	"age": {
		args:     []string{"--age", "42"},
		expected: int32(42),
		get:      func(o testExampleMessage) any { return o.GetAge() },
	},
	"ages": {
		args:     []string{"--ages", "42", "--ages", "43"},
		expected: []int32{42, 43},
		get:      func(o testExampleMessage) any { return o.GetAges() },
	},
	"id": {
		args:     []string{"--id", "123"},
		expected: int64(123),
		get:      func(o testExampleMessage) any { return o.GetId() },
	},
	"ids": {
		args:     []string{"--ids", "123", "--ids", "124"},
		expected: []int64{123, 124},
		get:      func(o testExampleMessage) any { return o.GetIds() },
	},
	"emails": {
		args:     []string{"--emails", "email1", "--emails", "email2"},
		expected: []string{"email1", "email2"},
		get:      func(o testExampleMessage) any { return o.GetEmails() },
	},
	"extra": {
		args:     []string{"--extra", `{"key":"value"}`},
		expected: map[string]string{"key": "value"},
		get:      func(o testExampleMessage) any { return o.GetExtra() },
	},
	"status": {
		args:     []string{"--status", "STATUS_OK"},
		expected: opaque.Status_STATUS_OK,
		get:      func(o testExampleMessage) any { return o.GetStatus() },
	},
	"statuses": {
		args:     []string{"--statuses", "STATUS_OK", "--statuses", "STATUS_ERROR"},
		expected: []opaque.Status{opaque.Status_STATUS_OK, opaque.Status_STATUS_ERROR},
		get:      func(o testExampleMessage) any { return o.GetStatuses() },
	},
	"size": {
		args:     []string{"--size", "12"},
		expected: uint32(12),
		get:      func(o testExampleMessage) any { return o.GetSize() },
	},
	"sizes": {
		args:     []string{"--sizes", "12", "--sizes", "13"},
		expected: []uint32{12, 13},
		get:      func(o testExampleMessage) any { return o.GetSizes() },
	},
	"length": {
		args:     []string{"--length", "1006"},
		expected: uint64(1006),
		get:      func(o testExampleMessage) any { return o.GetLength() },
	},
	"lengths": {
		args:     []string{"--lengths", "1006", "--lengths", "1007"},
		expected: []uint64{1006, 1007},
		get:      func(o testExampleMessage) any { return o.GetLengths() },
	},
	"connection-id": {
		args:     []string{"--connection-id", "68372"},
		expected: int32(68372),
		get:      func(o testExampleMessage) any { return o.GetConnectionId() },
	},
	"connection-ids": {
		args:     []string{"--connection-ids", "68372", "--connection-ids", "68373"},
		expected: []int32{68372, 68373},
		get:      func(o testExampleMessage) any { return o.GetConnectionIds() },
	},
	"session-id": {
		args:     []string{"--session-id", "9382784"},
		expected: int64(9382784),
		get:      func(o testExampleMessage) any { return o.GetSessionId() },
	},
	"session-ids": {
		args:     []string{"--session-ids", "9382784", "--session-ids", "9382785"},
		expected: []int64{9382784, 9382785},
		get:      func(o testExampleMessage) any { return o.GetSessionIds() },
	},
	"fixed-size": {
		args:     []string{"--fixed-size", "42"},
		expected: uint32(42),
		get:      func(o testExampleMessage) any { return o.GetFixedSize() },
	},
	"fixed-sizes": {
		args:     []string{"--fixed-sizes", "42", "--fixed-sizes", "43"},
		expected: []uint32{42, 43},
		get:      func(o testExampleMessage) any { return o.GetFixedSizes() },
	},
	"fixed-length": {
		args:     []string{"--fixed-length", "1006"},
		expected: uint64(1006),
		get:      func(o testExampleMessage) any { return o.GetFixedLength() },
	},
	"fixed-lengths": {
		args:     []string{"--fixed-lengths", "1006", "--fixed-lengths", "1007"},
		expected: []uint64{1006, 1007},
		get:      func(o testExampleMessage) any { return o.GetFixedLengths() },
	},
	"sfixed-size": {
		args:     []string{"--sfixed-size", "42"},
		expected: int32(42),
		get:      func(o testExampleMessage) any { return o.GetSfixedSize() },
	},
	"sfixed-sizes": {
		args:     []string{"--sfixed-sizes", "42", "--sfixed-sizes", "43"},
		expected: []int32{42, 43},
		get:      func(o testExampleMessage) any { return o.GetSfixedSizes() },
	},
	"sfixed-length": {
		args:     []string{"--sfixed-length", "1006"},
		expected: int64(1006),
		get:      func(o testExampleMessage) any { return o.GetSfixedLength() },
	},
	"sfixed-lengths": {
		args:     []string{"--sfixed-lengths", "1006", "--sfixed-lengths", "1007"},
		expected: []int64{1006, 1007},
		get:      func(o testExampleMessage) any { return o.GetSfixedLengths() },
	},
	"is-active": {
		args:     []string{"--is-active"},
		expected: true,
		get:      func(o testExampleMessage) any { return o.GetIsActive() },
	},
	"is-actives": {
		args:     []string{"--is-actives", "true", "--is-actives", "false"},
		expected: []bool{true, false},
		get:      func(o testExampleMessage) any { return o.GetIsActives() },
	},
	"data": {
		args:     []string{"--data", "Zm9vCg=="},
		expected: "foo\n",
		get:      func(o testExampleMessage) any { return string(o.GetData()) },
	},
	"datas": {
		args:     []string{"--datas", "Zm9vCg==", "--datas", "YmFyCg=="},
		expected: []string{"foo\n", "bar\n"},
		get: func(o testExampleMessage) any {
			return []string{string(o.GetDatas()[0]), string(o.GetDatas()[1])}
		},
	},
	"address": {
		args: []string{"--address", `{"street":"Main St","city":"Springfield","state":"IL","zip":"62701"}`},
		expected: &opaque.Address{
			Street: "Main St",
			City:   "Springfield",
			State:  "IL",
			Zip:    "62701",
		},
		get: func(o testExampleMessage) any { return o.GetAddress() },
	},
	"previous-addresses": {
		args: []string{
			"--previous-addresses", `{"street":"Main St","city":"Springfield","state":"IL","zip":"62701"}`,
			"--previous-addresses", `{"street":"Elm St","city":"Springfield","state":"IL","zip":"62702"}`,
		},
		expected: []*opaque.Address{
			{
				Street: "Main St",
				City:   "Springfield",
				State:  "IL",
				Zip:    "62701",
			},
			{
				Street: "Elm St",
				City:   "Springfield",
				State:  "IL",
				Zip:    "62702",
			},
		},
		get: func(o testExampleMessage) any { return o.GetPreviousAddresses() },
	},
	"oneof-name": {
		args:     []string{"--oneof-name", "foo"},
		expected: "foo",
		get:      func(o testExampleMessage) any { return o.GetOneofName() },
	},
	"oneof-score": {
		args:     []string{"--oneof-score", "3.14"},
		expected: float64(3.14),
		get:      func(o testExampleMessage) any { return o.GetOneofScore() },
	},
	"oneof-ratio": {
		args:     []string{"--oneof-ratio", "3.15"},
		expected: float32(3.15),
		get:      func(o testExampleMessage) any { return o.GetOneofRatio() },
	},
	"oneof-age": {
		args:     []string{"--oneof-age", "42"},
		expected: int32(42),
		get:      func(o testExampleMessage) any { return o.GetOneofAge() },
	},
	"oneof-id": {
		args:     []string{"--oneof-id", "123"},
		expected: int64(123),
		get:      func(o testExampleMessage) any { return o.GetOneofId() },
	},
	"oneof-status": {
		args:     []string{"--oneof-status", "STATUS_OK"},
		expected: opaque.Status_STATUS_OK,
		get:      func(o testExampleMessage) any { return o.GetOneofStatus() },
	},
	"oneof-size": {
		args:     []string{"--oneof-size", "12"},
		expected: uint32(12),
		get:      func(o testExampleMessage) any { return o.GetOneofSize() },
	},
	"oneof-length": {
		args:     []string{"--oneof-length", "1006"},
		expected: uint64(1006),
		get:      func(o testExampleMessage) any { return o.GetOneofLength() },
	},
	"oneof-connection-id": {
		args:     []string{"--oneof-connection-id", "68372"},
		expected: int32(68372),
		get:      func(o testExampleMessage) any { return o.GetOneofConnectionId() },
	},
	"oneof-session-id": {
		args:     []string{"--oneof-session-id", "9382784"},
		expected: int64(9382784),
		get:      func(o testExampleMessage) any { return o.GetOneofSessionId() },
	},
	"oneof-fixed-size": {
		args:     []string{"--oneof-fixed-size", "42"},
		expected: uint32(42),
		get:      func(o testExampleMessage) any { return o.GetOneofFixedSize() },
	},
	"oneof-fixed-length": {
		args:     []string{"--oneof-fixed-length", "1006"},
		expected: uint64(1006),
		get:      func(o testExampleMessage) any { return o.GetOneofFixedLength() },
	},
	"oneof-sfixed-size": {
		args:     []string{"--oneof-sfixed-size", "42"},
		expected: int32(42),
		get:      func(o testExampleMessage) any { return o.GetOneofSfixedSize() },
	},
	"oneof-sfixed-length": {
		args:     []string{"--oneof-sfixed-length", "1006"},
		expected: int64(1006),
		get:      func(o testExampleMessage) any { return o.GetOneofSfixedLength() },
	},
	"oneof-is-active": {
		args:     []string{"--oneof-is-active"},
		expected: true,
		get:      func(o testExampleMessage) any { return o.GetOneofIsActive() },
	},
	"oneof-data": {
		args:     []string{"--oneof-data", "Zm9vCg=="},
		expected: "foo\n",
		get:      func(o testExampleMessage) any { return string(o.GetOneofData()) },
	},
	"oneof-address": {
		args: []string{"--oneof-address", `{"street":"Main St","city":"Springfield","state":"IL","zip":"62701"}`},
		expected: &opaque.Address{
			Street: "Main St",
			City:   "Springfield",
			State:  "IL",
			Zip:    "62701",
		},
		get: func(o testExampleMessage) any { return o.GetOneofAddress() },
	},
	"all": {
		args: []string{
			"--name", "foo",
			"--score", "3.14",
			"--ratio", "3.15",
			"--age", "42",
			"--id", "123",
			"--emails", "email1", "--emails", "email2",
			"--extra", `{"key":"value"}`,
			"--status", "STATUS_OK",
			"--size", "12",
			"--length", "1006",
			"--connection-id", "68372",
			"--session-id", "9382784",
			"--fixed-size", "42",
			"--fixed-length", "1006",
			"--sfixed-size", "42",
			"--sfixed-length", "1006",
			"--is-active",
			"--data", "Zm9vCg==",
			"--address", `{"street":"Main St","city":"Springfield","state":"IL","zip":"62701"}`,
			"--previous-addresses", `{"street":"Main St","city":"Springfield","state":"IL","zip":"62701"}`,
			"--previous-addresses", `{"street":"Elm St","city":"Springfield","state":"IL","zip":"62702"}`,
			"--oneof-address", `{"street":"Main St","city":"Springfield","state":"IL","zip":"62701"}`,
		},
		expected: []string{
			`name:string(foo)`,
			`score:float64(3.14)`,
			`ratio:float32(3.15)`,
			`age:int32(42)`,
			`id:int64(123)`,
			`emails:[]string(["email1","email2"])`,
			`extra:map[string]string({"key":"value"})`,
			`status:opaque.Status(STATUS_OK)`,
			`size:uint32(12)`,
			`length:uint64(1006)`,
			`connection-id:int32(68372)`,
			`session-id:int64(9382784)`,
			`fixed-size:uint32(42)`,
			`fixed-length:uint64(1006)`,
			`sfixed-size:int32(42)`,
			`sfixed-length:int64(1006)`,
			`is-active:bool(true)`,
			"data:[]uint8(foo\n)",
			`address:*opaque.Address({"street":"Main St","city":"Springfield","state":"IL","zip":"62701"})`,
			`previous-addresses:[]*opaque.Address([{"street":"Main St","city":"Springfield","state":"IL","zip":"62701"},{"street":"Elm St","city":"Springfield","state":"IL","zip":"62702"}])`,
			`oneof-address:*opaque.Address({"street":"Main St","city":"Springfield","state":"IL","zip":"62701"})`,
		},
		get: func(o testExampleMessage) any {
			emails, _ := json.Marshal(o.GetEmails())
			extra, _ := json.Marshal(o.GetExtra())
			var adr bytes.Buffer
			json.Compact(&adr, []byte(protojson.Format(o.GetAddress())))
			address := adr.String()
			var oneofAdr bytes.Buffer
			json.Compact(&oneofAdr, []byte(protojson.Format(o.GetOneofAddress())))
			oneofAddress := oneofAdr.String()

			var previousAddresses []string
			for _, a := range o.GetPreviousAddresses() {
				var adr bytes.Buffer
				json.Compact(&adr, []byte(protojson.Format(a)))
				previousAddresses = append(previousAddresses, adr.String())
			}
			if len(previousAddresses) < 2 {
				panic(fmt.Sprintf("expected at least 2 previous addresses, got %d", len(previousAddresses)))
			}

			return []string{
				fmt.Sprintf("name:%T(%s)", o.GetName(), o.GetName()),
				fmt.Sprintf("score:%T(%.02f)", o.GetScore(), o.GetScore()),
				fmt.Sprintf("ratio:%T(%.02f)", o.GetRatio(), o.GetRatio()),
				fmt.Sprintf("age:%T(%d)", o.GetAge(), o.GetAge()),
				fmt.Sprintf("id:%T(%d)", o.GetId(), o.GetId()),
				fmt.Sprintf("emails:%T(%s)", o.GetEmails(), string(emails)),
				fmt.Sprintf("extra:%T(%s)", o.GetExtra(), string(extra)),
				fmt.Sprintf("status:%T(%v)", o.GetStatus(), o.GetStatus()),
				fmt.Sprintf("size:%T(%d)", o.GetSize(), o.GetSize()),
				fmt.Sprintf("length:%T(%d)", o.GetLength(), o.GetLength()),
				fmt.Sprintf("connection-id:%T(%d)", o.GetConnectionId(), o.GetConnectionId()),
				fmt.Sprintf("session-id:%T(%d)", o.GetSessionId(), o.GetSessionId()),
				fmt.Sprintf("fixed-size:%T(%d)", o.GetFixedSize(), o.GetFixedSize()),
				fmt.Sprintf("fixed-length:%T(%d)", o.GetFixedLength(), o.GetFixedLength()),
				fmt.Sprintf("sfixed-size:%T(%d)", o.GetSfixedSize(), o.GetSfixedSize()),
				fmt.Sprintf("sfixed-length:%T(%d)", o.GetSfixedLength(), o.GetSfixedLength()),
				fmt.Sprintf("is-active:%T(%v)", o.GetIsActive(), o.GetIsActive()),
				fmt.Sprintf("data:%T(%s)", o.GetData(), string(o.GetData())),
				fmt.Sprintf("address:%T(%s)", o.GetAddress(), address),
				fmt.Sprintf("previous-addresses:%T([%s,%s])", o.GetPreviousAddresses(), previousAddresses[0], previousAddresses[1]),
				fmt.Sprintf("oneof-address:%T(%s)", o.GetOneofAddress(), oneofAddress),
			}
		},
	},
}

func testCliFlags[
	CliOptions interface {
		WithClient(func(*cli.Context) (client.Client, error)) CliOptions
	},
](t *testing.T, prefix, cmd string, appCtor func(...CliOptions) (*cli.App, error), options CliOptions) {
	for _, desc := range workflow.DeterministicKeys(testFlags) {
		tc := testFlags[desc]
		for _, useSignal := range []bool{false, true} {
			suffix := " workflow-input"
			if useSignal {
				suffix = " signal-input"
			}
			t.Run(prefix+" "+desc+suffix, func(t *testing.T) {
				c := clientmocks.NewMockClient(t)
				c.EXPECT().Close().Return()

				var signalInput, workflowInput testExampleMessage
				c.EXPECT().
					SignalWithStartWorkflow(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					RunAndReturn(func(ctx context.Context, workflowID, signalName string, s any, swo client.StartWorkflowOptions, workflow any, args ...any) (client.WorkflowRun, error) {
						var ok bool
						signalInput, ok = s.(testExampleMessage)
						require.True(t, ok)
						require.NotNil(t, signalInput)
						workflowInput, ok = args[0].(testExampleMessage)
						require.True(t, ok)
						require.NotNil(t, workflowInput)
						run := clientutils.NewMockWorkflowRun(t)
						run.EXPECT().Get(mock.Anything, mock.Anything).RunAndReturn(func(ctx context.Context, i interface{}) error {
							return nil
						})
						return run, nil
					})

				app, err := appCtor(options.WithClient(func(*cli.Context) (client.Client, error) { return c, nil }))
				require.NoError(t, err)

				args := tc.args
				if useSignal {
					args, err = convert.MapSliceFunc(args, func(s string) (string, error) {
						if !strings.HasPrefix(s, "--") {
							return s, nil
						}
						return "--signal-" + prefix + "-" + s[2:], nil
					})
					require.NoError(t, err)
				}
				require.NoError(t, app.Run(append([]string{"opaque", cmd}, args...)))

				if len(tc.errors) > 0 {
					require.Error(t, err)
					for _, msg := range tc.errors {
						require.ErrorContains(t, err, msg)
					}
				} else {
					input := workflowInput
					if useSignal {
						input = signalInput
					}
					require.NoError(t, err)
					require.NotNil(t, input)
					out := tc.get(input)

					tt := reflect.TypeOf(out)
					if tt.Kind() == reflect.Slice {
						tt = tt.Elem()
					}

					if tt.Kind() == reflect.Ptr && tt.Implements(reflect.TypeOf((*proto.Message)(nil)).Elem()) {
						require.Empty(t, cmp.Diff(out, tc.expected, protocmp.Transform()))
					} else {
						require.Equal(t, tc.expected, out)
					}
				}
			})
		}
	}
}
