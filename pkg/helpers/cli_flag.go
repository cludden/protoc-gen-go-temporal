package helpers

// UnmarshalCliFlagsOptions configures the behavior of cli flag unmarshalling
type UnmarshalCliFlagsOptions struct {
	FromFile    string
	Prefix      string              // flag prefix
	PrefixFlags map[string]struct{} // flags that should be prefixed
}

func (o *UnmarshalCliFlagsOptions) FlagName(short string) string {
	if o.PrefixFlags == nil {
		return short
	} else if _, ok := o.PrefixFlags[short]; ok {
		return o.Prefix + "-" + short
	}
	return short
}

func FlattenUnmarshalCliFlagsOptions(opts ...UnmarshalCliFlagsOptions) UnmarshalCliFlagsOptions {
	if len(opts) == 0 {
		return UnmarshalCliFlagsOptions{}
	}
	result := opts[0]
	for _, opt := range opts[1:] {
		if opt.FromFile != "" {
			result.FromFile = opt.FromFile
		}
		if opt.Prefix != "" {
			result.Prefix = opt.Prefix
		}
		if opt.PrefixFlags != nil {
			if result.PrefixFlags == nil {
				result.PrefixFlags = make(map[string]struct{})
			}
			for k := range opt.PrefixFlags {
				result.PrefixFlags[k] = struct{}{}
			}
		}
	}
	return result
}
