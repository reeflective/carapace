package cmd

import (
	"testing"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/style"
)

func TestFiles(t *testing.T) {
	carapace.Test(t, rootCmd)(func(s carapace.Sandbox) {
		s.Files("example.txt", "")
		s.Run("action", "--files", "e").Expect(carapace.ActionValues("example.txt").StyleF(style.ForPathExt))
	})
}

func TestShells(t *testing.T) {
	carapace.Test(t, rootCmd)(func(s carapace.Sandbox) {
		s.Reply("chsh", "--list-shells").With("/bin/bash\n/bin/zsh")
		s.Run("action", "--shells", "").Expect(carapace.ActionValues("/bin/bash", "/bin/zsh"))
	})
}
