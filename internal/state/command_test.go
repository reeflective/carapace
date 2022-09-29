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

	state := NewCommandState2(cmd)

	cmd.Use = "changed"

	state.Restore(t)
}
