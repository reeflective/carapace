package cmd

import (
	"testing"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/style"
)

func TestMultiparts(t *testing.T) {
	carapace.Test(t, rootCmd)(func(s carapace.Sandbox) {
		s.Files(
			"dirA/file1.txt", "",
			"dirA/file2.png", "",
			"dirB/dirC/file3.go", "",
			"dirB/file4.md", "",
			"file5.go", "",
		)
		s.Run("multiparts", "").Expect(carapace.ActionValues("DIRECTORY", "FILE", "VALUE").Invoke(carapace.Context{}).Suffix("=").ToA())
		s.Run("multiparts", "D").Expect(carapace.ActionValues("DIRECTORY").Invoke(carapace.Context{}).Suffix("=").ToA())
		s.Run("multiparts", "DIRECTORY").Expect(carapace.ActionValues("DIRECTORY").Invoke(carapace.Context{}).Suffix("=").ToA())
		s.Run("multiparts", "DIRECTORY=").Expect(carapace.ActionValues("dirA/", "dirB/").StyleF(style.ForPathExt).Invoke(carapace.Context{}).Prefix("DIRECTORY=").ToA())
		s.Run("multiparts", "VALUE=").Expect(carapace.ActionValues("one", "two", "three").Invoke(carapace.Context{}).Prefix("VALUE=").ToA())
		s.Run("multiparts", "VALUE=o").Expect(carapace.ActionValues("one").Invoke(carapace.Context{}).Prefix("VALUE=").ToA())
		s.Run("multiparts", "VALUE=one,").Expect(carapace.ActionValues("DIRECTORY", "FILE").Invoke(carapace.Context{}).Prefix("VALUE=one,").Suffix("=").ToA())
		s.Run("multiparts", "VALUE=one,F").Expect(carapace.ActionValues("FILE").Invoke(carapace.Context{}).Prefix("VALUE=one,").Suffix("=").ToA())
		s.Run("multiparts", "VALUE=one,FILE=").Expect(carapace.ActionValues("dirA/", "dirB/", "file5.go").StyleF(style.ForPathExt).Invoke(carapace.Context{}).Prefix("VALUE=one,FILE=").ToA())
		s.Run("multiparts", "VALUE=one,FILE=dirB/").Expect(carapace.ActionValues("dirC/", "file4.md").StyleF(style.ForPathExt).Invoke(carapace.Context{}).Prefix("VALUE=one,FILE=dirB/").ToA())
	})
}
