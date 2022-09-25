package cmd

import (
	"testing"

	"github.com/rsteube/carapace"
)

func TestCondition(t *testing.T) {
	carapace.Test(t, rootCmd)(func(s carapace.Sandbox) {
		s.Run("condition", "").Expect(carapace.ActionMessage("flag --required must be set to valid: "))
		s.Run("condition", "").Expect(carapace.ActionMessage("flag --required must be set to valid: "))
		s.Run("condition", "--required", "invalid", "").Expect(carapace.ActionMessage("flag --required must be set to valid: invalid"))
		s.Run("condition", "--required", "valid", "").Expect(carapace.ActionValues("condition fulfilled"))
	})
}
