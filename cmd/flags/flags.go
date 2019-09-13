package flags

import (
	flag "github.com/spf13/pflag"
)

// GlobalFlags is the flags that contains the global flags
type GlobalFlags struct {
	Silent bool
	NoWarn bool
	Debug  bool

	Namespace   string
	KubeContext string
	Profile     string
	Vars        []string
}

// SetGlobalFlags applies the global flags
func SetGlobalFlags(flags *flag.FlagSet) *GlobalFlags {
	globalFlags := &GlobalFlags{
		Vars: []string{},
	}

	flags.BoolVar(&globalFlags.NoWarn, "no-warn", false, "If true does not show any warning when deploying into a different namespace or kube-context than before")
	flags.BoolVar(&globalFlags.Debug, "debug", false, "Prints the stack trace if an error occurs")
	flags.BoolVar(&globalFlags.Silent, "silent", false, "Run in silent mode and prevents any devspace log output except panics & fatals")

	flags.StringVarP(&globalFlags.Profile, "profile", "p", "", "The devspace profile to use (if there is any)")
	flags.StringVarP(&globalFlags.Namespace, "namespace", "n", "", "The kubernetes namespace to use")
	flags.StringVar(&globalFlags.KubeContext, "kube-context", "", "The kubernetes context to use")
	flags.StringSliceVar(&globalFlags.Vars, "var", []string{}, "Variables to override during execution (e.g. --var=MYVAR=MYVALUE)")

	return globalFlags
}
