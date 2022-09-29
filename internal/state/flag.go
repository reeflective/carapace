package state

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/rsteube/carapace/internal/assert"
	"github.com/spf13/pflag"
)

type FlagState struct {
	*pflag.Flag
	dump string

	Name                string
	Shorthand           string
	Usage               string
	Value               string
	DefValue            string
	Changed             bool
	NoOptDefVal         string
	Deprecated          string
	Hidden              bool
	ShorthandDeprecated string
	Annotations         map[string][]string

	// carapace-pflag
	Style int
}

func NewFlagState(f *pflag.Flag) *FlagState {
	s := &FlagState{
		Flag: f,
		dump: dump(f),

		Name:                f.Name,
		Shorthand:           f.Shorthand,
		Usage:               f.Usage,
		Value:               f.Value.String(),
		DefValue:            f.DefValue,
		Changed:             f.Changed,
		NoOptDefVal:         f.NoOptDefVal,
		Deprecated:          f.Deprecated,
		Hidden:              f.Hidden,
		ShorthandDeprecated: f.ShorthandDeprecated,
	}

	if f.Annotations != nil {
		s.Annotations = make(map[string][]string)
		for k, v := range f.Annotations {
			s.Annotations[k] = make([]string, 0, len(v))
			copy(s.Annotations[k], v)
		}
	}

	// TODO set style with reflect
	return s
}

func (s *FlagState) Restore(t T) {
	s.Flag.Name = s.Name
	s.Flag.Shorthand = s.Shorthand
	s.Flag.Usage = s.Usage
	s.Flag.Value.Set(s.Value)
	s.Flag.DefValue = s.DefValue
	s.Flag.Changed = s.Changed
	s.Flag.NoOptDefVal = s.NoOptDefVal
	s.Flag.Deprecated = s.Deprecated
	s.Flag.Hidden = s.Hidden
	s.Flag.ShorthandDeprecated = s.ShorthandDeprecated

	if s.Annotations == nil {
		s.Flag.Annotations = nil
	} else {
		s.Flag.Annotations = make(map[string][]string)
		for k, v := range s.Annotations {
			s.Flag.Annotations[k] = make([]string, 0, len(v))
			copy(s.Flag.Annotations[k], v)
		}
	}

	// TODO set style with reflect

	assert.Equal(t, s.dump, dump(s.Flag))
}

func dump(i interface{}) string {
	cfg := spew.NewDefaultConfig()
	cfg.SortKeys = true
	return cfg.Sdump(i)
}
