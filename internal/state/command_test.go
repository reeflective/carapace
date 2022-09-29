package state

import (
	"testing"

	"github.com/spf13/cobra"
)

func TestCommand(t *testing.T) {
	cmd := &cobra.Command{
		Use: "test",
		Run: func(cmd *cobra.Command, args []string) {},
	}

	state := NewCommandState(cmd)

	cmd.Use = "changed"
	cmd.Long = "changed"

	state.Restore(t)
}
