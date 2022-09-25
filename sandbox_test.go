package carapace

import (
	"testing"

	"github.com/spf13/cobra"
)

func TestSandbox(t *testing.T) {
	cmd := &cobra.Command{
		Use: "sandbox",
		Run: func(cmd *cobra.Command, args []string) {},
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd: true,
		},
	}

	Gen(cmd).PositionalCompletion(
		Batch(
			ActionValues("one", "two"),
			ActionValues("three"),
		).ToA(),
	)

	Test(t, cmd)(func(s Sandbox) {
		s.Run("").Expect(ActionValues("one", "two", "three"))
	})
}
