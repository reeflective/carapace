package state

import (
	"testing"

	"github.com/spf13/cobra"
)

func TestState(t *testing.T) {
	cmd := &cobra.Command{
		Use: "example",
		Run: func(cmd *cobra.Command, args []string) {},
	}
	cmd.Flags().Bool("first", false, "first flag")
	cmd.Flags().String("second", "", "second flag")

	state := NewCommandState(cmd)
	cmd.SetArgs([]string{"--first", "pos1"})
	state.Restore(t)
}
