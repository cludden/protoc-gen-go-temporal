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
