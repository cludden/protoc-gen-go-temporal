package docs

import (
	"bufio"
	"bytes"
	"container/list"
	"fmt"
	"strings"
	"time"

	patch "github.com/alta/protopatch/patch/gopb"
	temporalv1 "github.com/cludden/protoc-gen-go-temporal/gen/temporal/v1"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
)

type (
	Activity struct {
		Descriptor
		Comments               Comments
		Deprecated             bool
		HasNonDefaultOptions   bool
		HeartbeatTimeout       time.Duration
		Input                  string
		Name                   string
		Output                 string
		RetryPolicy            *RetryPolicy
		ScheduleToCloseTimeout time.Duration
		ScheduleToStartTimeout time.Duration
		StartToCloseTimeout    time.Duration
		WaitForCancellation    bool
		TaskQueue              string
	}

	Comments struct {
		Leading       string
		LeadingLines  []string
		Trailing      string
		TrailingLines []string
	}

	Data struct {
		Activities         map[string]Activity
		Enums              map[string]Enum
		Messages           map[string]Message
		ReferencedMessages map[string][]string
		Packages           map[string]Package
		Queries            map[string]Query
		Services           map[string]Service
		Signals            map[string]Signal
		Updates            map[string]Update
		Workflows          map[string]Workflow
	}

	Descriptor struct {
		Name     string
		FullName string
		Package  string
		File     File
	}

	Enum struct {
		Descriptor
		Comments Comments
		Values   []EnumValue
	}

	EnumValue struct {
		Name     string
		Comments Comments
	}

	Field struct {
		Descriptor
		Comments Comments
		GoName   string
		GoTags   string
		JSONName string
		Type     string
	}

	File struct {
		Name string
		Path string
	}

	Message struct {
		Descriptor
		Comments Comments
		Enums    []string
		Fields   map[string]Field
		Messages []string
	}

	Package struct {
		Descriptor
		Enums                []string
		HasTemporalResources bool
		Messages             []string
		ReferencedMessages   []string
		Services             []string
	}

	Query struct {
		Descriptor
		Comments             Comments
		Deprecated           bool
		HasNonDefaultOptions bool
		Input                string
		Name                 string
		Output               string
		XNS                  *XNS
	}

	RetryPolicy struct {
		BackoffCoefficient     float64
		InitialInterval        time.Duration
		MaxAttempts            int
		MaxInterval            time.Duration
		NonRetryableErrorTypes []string
	}

	Service struct {
		Descriptor
		Activities           []string
		Comments             Comments
		Deprecated           bool
		HasTemporalResources bool
		Queries              []string
		Signals              []string
		TaskQueue            string
		Updates              []string
		Workflows            []string
	}

	Signal struct {
		Descriptor
		Comments             Comments
		Deprecated           bool
		HasNonDefaultOptions bool
		Input                string
		Name                 string
		XNS                  *XNS
	}

	Update struct {
		Descriptor
		Comments             Comments
		Deprecated           bool
		HasNonDefaultOptions bool
		Input                string
		Name                 string
		Output               string
		Validate             bool
		WaitPolicy           string
		XNS                  *XNS
	}

	Workflow struct {
		Descriptor
		Comments            Comments
		Deprecated          bool
		ExecutionTimeout    time.Duration
		ID                  string
		IDReusePolicy       string
		Input               string
		Name                string
		Output              string
		ParentClosePolicy   string
		Queries             map[string]WorkflowQuery
		RetryPolicy         *RetryPolicy
		RunTimeout          time.Duration
		SearchAttributes    string
		Signals             map[string]WorkflowSignal
		TaskQueue           string
		TaskTimeout         time.Duration
		Updates             map[string]WorkflowUpdate
		WaitForCancellation bool
		XNS                 *XNS
	}

	WorkflowQuery struct{}

	WorkflowSignal struct {
		Start bool
	}

	WorkflowUpdate struct{}

	XNS struct {
		HeartbeatInterval      time.Duration
		HeartbeatTimeout       time.Duration
		Name                   string
		RetryPolicy            *RetryPolicy
		ScheduleToCloseTimeout time.Duration
		ScheduleToStartTimeout time.Duration
		StartToCloseTimeout    time.Duration
		TaskQueue              string
	}
)

func Parse(p *protogen.Plugin) (*Data, error) {
	data := &Data{
		Activities: map[string]Activity{},
		Enums:      map[string]Enum{},
		Messages:   map[string]Message{},
		Packages:   map[string]Package{},
		Queries:    map[string]Query{},
		Services:   map[string]Service{},
		Signals:    map[string]Signal{},
		Updates:    map[string]Update{},
		Workflows:  map[string]Workflow{},
	}

	refs := newVisitor().walk(p)
	data.ReferencedMessages = refs

	for _, f := range p.Files {
		pkgName := f.Proto.GetPackage()
		file := File{
			Name: string(f.Desc.Name()),
			Path: f.Desc.Path(),
		}

		pkg, ok := data.Packages[pkgName]
		if !ok {
			pkg = Package{
				Descriptor: Descriptor{
					Name: pkgName,
				},
				Enums:              []string{},
				Messages:           []string{},
				ReferencedMessages: refs[pkgName],
				Services:           []string{},
			}
		}

		for _, e := range f.Enums {
			enum := parseEnum(file, pkgName, e)
			pkg.Enums = append(pkg.Enums, enum.FullName)
			data.Enums[string(e.Desc.FullName())] = enum
		}

		seen := map[string]struct{}{}
		messages := list.New()
		for _, msg := range f.Messages {
			messages.PushBack(msg)
			seen[string(msg.Desc.FullName())] = struct{}{}
		}

		for rawmsg := messages.Front(); rawmsg != nil; rawmsg = rawmsg.Next() {
			m := rawmsg.Value.(*protogen.Message)
			msg := Message{
				Comments: parseComments(m.Comments),
				Descriptor: Descriptor{
					Name:     string(m.Desc.Name()),
					FullName: string(m.Desc.FullName()),
					File:     file,
					Package:  pkgName,
				},
				Enums:    []string{},
				Fields:   map[string]Field{},
				Messages: []string{},
			}

			for _, e := range m.Enums {
				enum := parseEnum(file, pkgName, e)
				pkg.Enums = append(pkg.Enums, enum.FullName)
				data.Enums[string(e.Desc.FullName())] = enum
			}

			for _, fl := range m.Fields {
				field := Field{
					Comments: parseComments(fl.Comments),
					Descriptor: Descriptor{
						Name:     string(fl.Desc.Name()),
						FullName: string(fl.Desc.FullName()),
						File:     file,
						Package:  pkgName,
					},
					GoName:   fl.GoName,
					JSONName: fl.Desc.JSONName(),
				}

				if opts, ok := proto.GetExtension(fl.Desc.Options(), patch.E_Field).(*patch.Options); ok {
					if v := opts.GetName(); v != "" {
						field.GoName = v
					}
					if v := opts.GetTags(); v != "" {
						field.GoTags = v
					}
				}

				switch fl.Desc.Kind() {
				case protoreflect.EnumKind:
					field.Type = string(fl.Enum.Desc.FullName())
					msg.Enums = append(msg.Enums, string(fl.Enum.Desc.FullName()))
				case protoreflect.MessageKind:
					field.Type = string(fl.Message.Desc.FullName())
					msg.Messages = append(msg.Messages, string(fl.Message.Desc.FullName()))
					if _, ok := seen[string(fl.Message.Desc.FullName())]; !ok {
						messages.PushBack(fl.Message)
						seen[string(fl.Message.Desc.FullName())] = struct{}{}
					}

				default:
					field.Type = fl.Desc.Kind().String()
				}
				msg.Fields[string(fl.Desc.Name())] = field
			}

			pkg.Messages = append(pkg.Messages, msg.FullName)
			data.Messages[string(m.Desc.FullName())] = msg
		}

		for _, s := range f.Services {
			opts, ok := proto.GetExtension(s.Desc.Options(), temporalv1.E_Service).(*temporalv1.ServiceOptions)
			if !ok {
				continue
			}

			svc := Service{
				Activities: []string{},
				Comments:   parseComments(s.Comments),
				Descriptor: Descriptor{
					Name:     string(s.Desc.Name()),
					FullName: string(s.Desc.FullName()),
					File:     file,
					Package:  pkgName,
				},
				Queries:   []string{},
				Signals:   []string{},
				TaskQueue: opts.GetTaskQueue(),
				Updates:   []string{},
				Workflows: []string{},
			}

			for _, m := range s.Methods {
				if opts, ok := proto.GetExtension(m.Desc.Options(), temporalv1.E_Workflow).(*temporalv1.WorkflowOptions); ok && opts != nil {
					workflow := Workflow{
						Comments:         parseComments(m.Comments),
						Deprecated:       m.Desc.Options().(*descriptorpb.MethodOptions).GetDeprecated(),
						ExecutionTimeout: opts.GetExecutionTimeout().AsDuration(),
						Descriptor: Descriptor{
							Name:     string(m.Desc.Name()),
							FullName: string(m.Desc.FullName()),
							File:     file,
							Package:  pkgName,
						},
						ID:                  opts.GetId(),
						IDReusePolicy:       opts.GetIdReusePolicy().String(),
						Queries:             map[string]WorkflowQuery{},
						RetryPolicy:         parseRetryPolicy(opts.GetRetryPolicy()),
						RunTimeout:          opts.GetRunTimeout().AsDuration(),
						SearchAttributes:    strings.TrimSpace(opts.GetSearchAttributes()),
						Signals:             map[string]WorkflowSignal{},
						TaskQueue:           strings.TrimSpace(opts.GetTaskQueue()),
						TaskTimeout:         opts.GetTaskTimeout().AsDuration(),
						Updates:             map[string]WorkflowUpdate{},
						WaitForCancellation: opts.GetWaitForCancellation(),
						XNS:                 parseXNS(opts.GetXns()),
					}

					if notEmpty(m.Input) {
						workflow.Input = string(m.Input.Desc.FullName())

					}
					if notEmpty(m.Output) {
						workflow.Output = string(m.Output.Desc.FullName())
					}

					if v := opts.GetName(); v != "" {
						workflow.Name = v
					} else {
						workflow.Name = string(m.Desc.FullName())
					}
					if v := opts.GetParentClosePolicy(); v != temporalv1.ParentClosePolicy_PARENT_CLOSE_POLICY_UNSPECIFIED {
						workflow.ParentClosePolicy = v.String()
					}

					svcPrefix := strings.Split(workflow.FullName, ".")
					for _, q := range opts.GetQuery() {
						workflow.Queries[fmt.Sprintf("%s.%s", strings.Join(svcPrefix[:len(svcPrefix)-1], "."), q.GetRef())] = WorkflowQuery{}
					}
					for _, s := range opts.GetSignal() {
						workflow.Signals[fmt.Sprintf("%s.%s", strings.Join(svcPrefix[:len(svcPrefix)-1], "."), s.GetRef())] = WorkflowSignal{
							Start: s.GetStart(),
						}
					}
					for _, u := range opts.GetUpdate() {
						workflow.Updates[fmt.Sprintf("%s.%s", strings.Join(svcPrefix[:len(svcPrefix)-1], "."), u.GetRef())] = WorkflowUpdate{}
					}

					svc.Workflows = append(svc.Workflows, workflow.FullName)
					data.Workflows[workflow.FullName] = workflow
				}

				if opts, ok := proto.GetExtension(m.Desc.Options(), temporalv1.E_Query).(*temporalv1.QueryOptions); ok && opts != nil {
					query := Query{
						Comments:   parseComments(m.Comments),
						Deprecated: m.Desc.Options().(*descriptorpb.MethodOptions).GetDeprecated(),
						Descriptor: Descriptor{
							Name:     string(m.Desc.Name()),
							FullName: string(m.Desc.FullName()),
							File:     file,
							Package:  pkgName,
						},
						XNS: parseXNS(opts.GetXns()),
					}

					if notEmpty(m.Input) {
						query.Input = string(m.Input.Desc.FullName())

					}
					if notEmpty(m.Output) {
						query.Output = string(m.Output.Desc.FullName())
					}

					if v := opts.GetName(); v != "" {
						query.Name = v
					} else {
						query.Name = string(m.Desc.FullName())
					}

					if query.XNS != nil {
						query.HasNonDefaultOptions = true
					}

					svc.Queries = append(svc.Queries, query.FullName)
					data.Queries[query.FullName] = query
				}

				if opts, ok := proto.GetExtension(m.Desc.Options(), temporalv1.E_Signal).(*temporalv1.SignalOptions); ok && opts != nil {
					signal := Signal{
						Comments:   parseComments(m.Comments),
						Deprecated: m.Desc.Options().(*descriptorpb.MethodOptions).GetDeprecated(),
						Descriptor: Descriptor{
							Name:     string(m.Desc.Name()),
							FullName: string(m.Desc.FullName()),
							File:     file,
							Package:  pkgName,
						},
						XNS: parseXNS(opts.GetXns()),
					}

					if notEmpty(m.Input) {
						signal.Input = string(m.Input.Desc.FullName())
					}

					if v := opts.GetName(); v != "" {
						signal.Name = v
					} else {
						signal.Name = string(m.Desc.FullName())
					}

					if signal.XNS != nil {
						signal.HasNonDefaultOptions = true
					}

					svc.Signals = append(svc.Signals, signal.FullName)
					data.Signals[signal.FullName] = signal
				}

				if opts, ok := proto.GetExtension(m.Desc.Options(), temporalv1.E_Update).(*temporalv1.UpdateOptions); ok && opts != nil {
					update := Update{
						Comments:   parseComments(m.Comments),
						Deprecated: m.Desc.Options().(*descriptorpb.MethodOptions).GetDeprecated(),
						Descriptor: Descriptor{
							Name:     string(m.Desc.Name()),
							FullName: string(m.Desc.FullName()),
							File:     file,
							Package:  pkgName,
						},
						Validate: opts.GetValidate(),
						XNS:      parseXNS(opts.GetXns()),
					}

					if notEmpty(m.Input) {
						update.Input = string(m.Input.Desc.FullName())
					}
					if notEmpty(m.Output) {
						update.Output = string(m.Output.Desc.FullName())
					}

					if v := opts.GetName(); v != "" {
						update.Name = v
					} else {
						update.Name = string(m.Desc.FullName())
					}
					if v := opts.GetWaitPolicy(); v != temporalv1.WaitPolicy_WAIT_POLICY_UNSPECIFIED {
						update.WaitPolicy = v.String()
					}

					if update.Validate || update.WaitPolicy != "" || update.XNS != nil {
						update.HasNonDefaultOptions = true
					}

					svc.Updates = append(svc.Updates, update.FullName)
					data.Updates[update.FullName] = update
				}

				if opts, ok := proto.GetExtension(m.Desc.Options(), temporalv1.E_Activity).(*temporalv1.ActivityOptions); ok && opts != nil {
					activity := Activity{
						Comments:   parseComments(m.Comments),
						Deprecated: m.Desc.Options().(*descriptorpb.MethodOptions).GetDeprecated(),
						Descriptor: Descriptor{
							Name:     string(m.Desc.Name()),
							FullName: string(m.Desc.FullName()),
							File:     file,
							Package:  pkgName,
						},
						HeartbeatTimeout:       opts.GetHeartbeatTimeout().AsDuration(),
						RetryPolicy:            parseRetryPolicy(opts.GetRetryPolicy()),
						ScheduleToCloseTimeout: opts.GetHeartbeatTimeout().AsDuration(),
						ScheduleToStartTimeout: opts.GetScheduleToStartTimeout().AsDuration(),
						StartToCloseTimeout:    opts.GetStartToCloseTimeout().AsDuration(),
						WaitForCancellation:    opts.GetWaitForCancellation(),
						TaskQueue:              opts.GetTaskQueue(),
					}

					if notEmpty(m.Input) {
						activity.Input = string(m.Input.Desc.FullName())
					}
					if notEmpty(m.Output) {
						activity.Output = string(m.Output.Desc.FullName())
					}

					if v := opts.GetName(); v != "" {
						activity.Name = v
					} else {
						activity.Name = string(m.Desc.FullName())
					}

					switch {
					case activity.HeartbeatTimeout > 0,
						activity.RetryPolicy != nil,
						activity.ScheduleToCloseTimeout > 0,
						activity.ScheduleToStartTimeout > 0,
						activity.StartToCloseTimeout > 0,
						activity.WaitForCancellation,
						activity.TaskQueue != "":
						activity.HasNonDefaultOptions = true
					}

					svc.Activities = append(svc.Activities, activity.FullName)
					data.Activities[activity.FullName] = activity
				}
			}

			if len(svc.Activities) > 0 || len(svc.Queries) > 0 || len(svc.Signals) > 0 || len(svc.Updates) > 0 || len(svc.Workflows) > 0 {
				pkg.HasTemporalResources, svc.HasTemporalResources = true, true
			}

			pkg.Services = append(pkg.Services, svc.FullName)
			data.Services[string(s.Desc.FullName())] = svc
		}

		data.Packages[pkgName] = pkg
	}

	return data, nil
}

func notEmpty(msg *protogen.Message) bool {
	return msg.Desc.FullName() != "google.protobuf.Empty"
}

func parseComment(cmt string) (string, []string) {
	var b bytes.Buffer
	var lines []string
	for s := bufio.NewScanner(strings.NewReader(cmt)); s.Scan(); {
		line := strings.TrimPrefix(s.Text(), "// ")
		if line == "//" {
			line = ""
		}
		fmt.Fprintln(&b, line)
		lines = append(lines, line)
	}
	return b.String(), lines
}

func parseComments(c protogen.CommentSet) Comments {
	leading, leadingLines := parseComment(c.Leading.String())
	trailing, trailingLines := parseComment(c.Trailing.String())
	return Comments{
		Leading:       leading,
		LeadingLines:  leadingLines,
		Trailing:      trailing,
		TrailingLines: trailingLines,
	}
}

func parseEnum(file File, pkgName string, e *protogen.Enum) Enum {
	enum := Enum{
		Comments: parseComments(e.Comments),
		Descriptor: Descriptor{
			Name:     string(e.Desc.Name()),
			FullName: string(e.Desc.FullName()),
			File:     file,
			Package:  pkgName,
		},
		Values: []EnumValue{},
	}
	for _, v := range e.Values {
		enum.Values = append(enum.Values, EnumValue{
			Name:     string(v.Desc.Name()),
			Comments: parseComments(v.Comments),
		})
	}
	return enum
}

func parseRetryPolicy(rp *temporalv1.RetryPolicy) *RetryPolicy {
	if rp == nil {
		return nil
	}
	result := &RetryPolicy{
		BackoffCoefficient:     rp.GetBackoffCoefficient(),
		InitialInterval:        rp.GetInitialInterval().AsDuration(),
		MaxAttempts:            int(rp.GetMaxAttempts()),
		MaxInterval:            rp.GetMaxInterval().AsDuration(),
		NonRetryableErrorTypes: rp.GetNonRetryableErrorTypes(),
	}
	if result.BackoffCoefficient != 0 || result.InitialInterval > 0 || result.MaxAttempts != 0 || result.MaxInterval > 0 || len(result.NonRetryableErrorTypes) > 0 {
		return result
	}
	return nil
}

func parseXNS(opts *temporalv1.XNSActivityOptions) *XNS {
	if opts == nil {
		return nil
	}
	result := &XNS{
		HeartbeatInterval:      opts.GetHeartbeatInterval().AsDuration(),
		HeartbeatTimeout:       opts.GetHeartbeatTimeout().AsDuration(),
		Name:                   opts.GetName(),
		RetryPolicy:            parseRetryPolicy(opts.GetRetryPolicy()),
		ScheduleToCloseTimeout: opts.GetScheduleToCloseTimeout().AsDuration(),
		ScheduleToStartTimeout: opts.GetScheduleToStartTimeout().AsDuration(),
		StartToCloseTimeout:    opts.GetStartToCloseTimeout().AsDuration(),
		TaskQueue:              opts.GetTaskQueue(),
	}
	if result.HeartbeatInterval != 0 || result.HeartbeatTimeout != 0 || result.Name != "" || result.RetryPolicy != nil || result.ScheduleToCloseTimeout > 0 || result.ScheduleToStartTimeout > 0 || result.StartToCloseTimeout > 0 || result.TaskQueue != "" {
		return result
	}
	return nil
}
