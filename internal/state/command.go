package state

import (
	"bytes"
	"context"
	"io"
	"testing"

	"github.com/rsteube/carapace/internal/assert"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type CommandState2 struct {
	*cobra.Command
	dump          string
	flagSetStates map[*pflag.FlagSet]*FlagSetState

	Use                    string
	Aliases                []string
	SuggestFor             []string
	Short                  string
	Long                   string
	Example                string
	ValidArgs              []string
	ValidArgsFunction      func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective)
	Args                   cobra.PositionalArgs
	ArgAliases             []string
	BashCompletionFunction string
	Deprecated             string
	Annotations            map[string]string
	Version                string
	PersistentPreRun       func(cmd *cobra.Command, args []string)
	PersistentPreRunE      func(cmd *cobra.Command, args []string) error
	PreRun                 func(cmd *cobra.Command, args []string)
	PreRunE                func(cmd *cobra.Command, args []string) error
	Run                    func(cmd *cobra.Command, args []string)
	RunE                   func(cmd *cobra.Command, args []string) error
	PostRun                func(cmd *cobra.Command, args []string)
	PostRunE               func(cmd *cobra.Command, args []string) error
	PersistentPostRun      func(cmd *cobra.Command, args []string)
	PersistentPostRunE     func(cmd *cobra.Command, args []string) error
	args                   []string
	flagErrorBuf           *bytes.Buffer
	flags                  *pflag.FlagSet
	pflags                 *pflag.FlagSet
	lflags                 *pflag.FlagSet
	iflags                 *pflag.FlagSet
	parentsPflags          *pflag.FlagSet
	globNormFunc           func(f *pflag.FlagSet, name string) pflag.NormalizedName
	usageFunc              func(*cobra.Command) error
	usageTemplate          string
	flagErrorFunc          func(*cobra.Command, error) error
	helpTemplate           string
	helpFunc               func(*cobra.Command, []string)
	helpCommand            *cobra.Command
	versionTemplate        string
	inReader               io.Reader
	outWriter              io.Writer
	errWriter              io.Writer
	FParseErrWhitelist     cobra.FParseErrWhitelist
	CompletionOptions      cobra.CompletionOptions
	commandsAreSorted      bool
	commandCalledAs        struct {
		name   string
		called bool
	}
	ctx                        context.Context
	commands                   []*cobra.Command
	parent                     *cobra.Command
	commandsMaxUseLen          int
	commandsMaxCommandPathLen  int
	commandsMaxNameLen         int
	TraverseChildren           bool
	Hidden                     bool
	SilenceErrors              bool
	SilenceUsage               bool
	DisableFlagParsing         bool
	DisableAutoGenTag          bool
	DisableFlagsInUseLine      bool
	DisableSuggestions         bool
	SuggestionsMinimumDistance int
}

func NewCommandState2(cmd *cobra.Command) *CommandState2 {
	s := &CommandState2{
		Command: cmd,
		dump:    dump(cmd),
		// flagSetStates map[*pflag.FlagSet]*FlagSetState

		Use: cmd.Use,
        Aliases: copyStringSlice(cmd.Aliases),
        SuggestFor: copyStringSlice(cmd.SuggestFor),
        Short: cmd.Short,
		//Long string
		//Example string
		//ValidArgs []string
		//ValidArgsFunction func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective)
		//Args cobra.PositionalArgs
		//ArgAliases []string
		//BashCompletionFunction string
		//Deprecated string
		//Annotations map[string]string
		//Version string
		//PersistentPreRun func(cmd *cobra.Command, args []string)
		//PersistentPreRunE func(cmd *cobra.Command, args []string) error
		//PreRun func(cmd *cobra.Command, args []string)
		//PreRunE func(cmd *cobra.Command, args []string) error
		//Run func(cmd *cobra.Command, args []string)
		//RunE func(cmd *cobra.Command, args []string) error
		//PostRun func(cmd *cobra.Command, args []string)
		//PostRunE func(cmd *cobra.Command, args []string) error
		//PersistentPostRun func(cmd *cobra.Command, args []string)
		//PersistentPostRunE func(cmd *cobra.Command, args []string) error
		//args []string
		//flagErrorBuf *bytes.Buffer
		//flags *pflag.FlagSet
		//pflags *pflag.FlagSet
		//lflags *pflag.FlagSet
		//iflags *pflag.FlagSet
		//parentsPflags *pflag.FlagSet
		//globNormFunc func(f *pflag.FlagSet, name string) pflag.NormalizedName
		//usageFunc func(*cobra.Command) error
		//usageTemplate string
		//flagErrorFunc func(*cobra.Command, error) error
		//helpTemplate string
		//helpFunc func(*cobra.Command, []string)
		//helpCommand *cobra.Command
		//versionTemplate string
		//inReader io.Reader
		//outWriter io.Writer
		//errWriter io.Writer
		//FParseErrWhitelist cobra.FParseErrWhitelist
		//CompletionOptions cobra.CompletionOptions
		//commandsAreSorted bool
		//commandCalledAs struct {
		//	name   string
		//	called bool
		//}
		//ctx context.Context
		//commands []*cobra.Command
		//parent *cobra.Command
		//commandsMaxUseLen         int
		//commandsMaxCommandPathLen int
		//commandsMaxNameLen        int
		//TraverseChildren bool
		//Hidden bool
		//SilenceErrors bool
		//SilenceUsage bool
		//DisableFlagParsing bool
		//DisableAutoGenTag bool
		//DisableFlagsInUseLine bool
		//DisableSuggestions bool
		//SuggestionsMinimumDistance int

	}

	return s
}

func (s *CommandState2) Restore(t *testing.T) {
	// TODO restore

	if s.flagSetStates != nil {
		for _, f := range s.flagSetStates {
			f.Restore(t)
		}
	}

	assert.Equal(t, s.dump, dump(s.Command))
}
