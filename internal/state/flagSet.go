package state

import (
	"io"
	"testing"

	goflag "flag"

	"github.com/rsteube/carapace/internal/assert"
	"github.com/spf13/pflag"
)

type FlagSetState struct {
	*pflag.FlagSet
	dump       string
	flagStates map[*pflag.Flag]*FlagState

	Usage                func()
	SortFlags            bool
	ParseErrorsWhitelist pflag.ParseErrorsWhitelist
	name                 string
	parsed               bool
	actual               map[pflag.NormalizedName]*pflag.Flag
	orderedActual        []*pflag.Flag
	sortedActual         []*pflag.Flag
	formal               map[pflag.NormalizedName]*pflag.Flag
	orderedFormal        []*pflag.Flag
	sortedFormal         []*pflag.Flag
	shorthands           map[byte]*pflag.Flag
	args                 []string
	argsLenAtDash        int
	errorHandling        pflag.ErrorHandling
	output               io.Writer
	interspersed         bool
	normalizeNameFunc    func(f *pflag.FlagSet, name string) pflag.NormalizedName

	addedGoFlagSets []*goflag.FlagSet
}

func NewFlagSetState(fs *pflag.FlagSet) *FlagSetState {
	s := &FlagSetState{
		FlagSet:    fs,
		dump:       dump(fs),
		flagStates: make(map[*pflag.Flag]*FlagState),

		Usage: fs.Usage,

		// TODO rest
	}

	return s
}

func (s *FlagSetState) Restore(t *testing.T) {
	// TODO restore

	if s.flagStates != nil {
		for _, f := range s.flagStates {
			f.Restore(t)
		}
	}

	assert.Equal(t, s.dump, dump(s.FlagSet))
}
