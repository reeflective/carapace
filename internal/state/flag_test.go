package state

import (
	"testing"

	"github.com/spf13/pflag"
)

func TestFlag(t *testing.T) {
	fs := pflag.NewFlagSet("", pflag.ContinueOnError)
	fs.String("test", "", "usage")

	f := fs.Lookup("test")

	state := NewFlagState(f)
	f.Deprecated = "Deprecated"
	f.DefValue = "arr"
	f.Changed = true
	f.Value.Set("arr")
	state.Restore(t)
}
