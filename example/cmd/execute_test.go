package cmd

import (
	"testing"

	"github.com/rsteube/carapace"
)

func TestExecute(t *testing.T) {
	carapace.Test(t, rootCmd)(func(s carapace.Sandbox) {
		s.Run("execute", "").Expect(carapace.ActionValues("one", "two"))
		s.Run("execute", "o").Expect(carapace.ActionValues("one"))
		s.Run("execute", "one", "").Expect(carapace.ActionValues("three", "four"))
	})
}
