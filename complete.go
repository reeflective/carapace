package carapace

import (
	"os"

	"github.com/carapace-sh/carapace/internal/config"
	"github.com/carapace-sh/carapace/internal/shell/bash"
	"github.com/carapace-sh/carapace/internal/shell/cmd_clink"
	"github.com/carapace-sh/carapace/internal/shell/nushell"
	"github.com/carapace-sh/carapace/pkg/ps"
	"github.com/spf13/cobra"
)

// Complete can be used by Go programs wishing to produce completions for
// themselves, without passing through shell snippets/output or export formats.
//
// The `onFinalize` function parameter, if non nil, will be called after having
// generated the completions from the given command/tree. This function is generally
// used to reset the command tree, which is needed when the Go program is a shell itself.
// Also, and before calling `onFinalize` if not nil, the completion storage is cleared.
func Complete(cmd *cobra.Command, args []string, onFinalize func()) (common.RawValues, common.Meta) {
	// Generate the completion as normally done for an external system shell
	action, current := generate(cmd, args)

	// And adapt/fetch the results from invoked action
	return internalValues(action, current, onFinalize)
}

func complete(cmd *cobra.Command, args []string) (string, error) {
	switch len(args) {
	case 0:
		return Gen(cmd).Snippet(ps.DetermineShell())
	case 1:
		return Gen(cmd).Snippet(args[0])
	default:
		initHelpCompletion(cmd)

		switch ps.DetermineShell() {
		case "nushell":
			args = nushell.Patch(args) // handle open quotes
			LOG.Printf("patching args to %#v", args)
		case "bash": // TODO what about oil and such?
			LOG.Printf("COMP_LINE is %#v", os.Getenv("COMP_LINE"))
			LOG.Printf("COMP_POINT is %#v", os.Getenv("COMP_POINT"))
			LOG.Printf("COMP_WORDBREAKS is %#v", os.Getenv("COMP_WORDBREAKS"))
			var err error
			args, err = bash.Patch(args) // handle redirects
			LOG.Printf("patching args to %#v", args)
			if err != nil {
				context := NewContext(args...)
				if _, ok := err.(bash.RedirectError); ok {
					LOG.Printf("completing redirect target for %#v", args)
					return ActionFiles().Invoke(context).value(args[0], args[len(args)-1]), nil
				}
				return ActionMessage(err.Error()).Invoke(context).value(args[0], args[len(args)-1]), nil
			}
		case "cmd-clink":
			var err error
			args, err = cmd_clink.Patch(args)
			LOG.Printf("patching args to %#v", args)
			if err != nil {
				context := NewContext(args...)
				return ActionMessage(err.Error()).Invoke(context).value(args[0], args[len(args)-1]), nil
			}
		}

		action, context := traverse(cmd, args[2:])
		if err := config.Load(); err != nil {
			action = ActionMessage("failed to load config: " + err.Error())
		}
		return action.Invoke(context).value(args[0], args[len(args)-1]), nil
	}
}

func internalValues(a InvokedAction, current string, onFinalize func()) (common.RawValues, common.Meta) {
	unsorted := common.RawValues(a.rawValues)
	sorted := make(common.RawValues, 0)

	// Ensure values are sorted.
	unsorted.EachTag(func(tag string, values common.RawValues) {
		vals := make(common.RawValues, len(values))
		for index, val := range values {
			if !a.meta.Nospace.Matches(val.Value) {
				val.Value += " "
			}
			if val.Style != "" {
				val.Style = style.SGR(val.Style)
			}

			vals[index] = val
		}
		sorted = append(sorted, vals...)
	})

	// Merge/filter completions and meta stuff.
	filtered := sorted.FilterPrefix(current)
	filtered = a.meta.Messages.Integrate(filtered, current)

	// Reset the storage (empty all commands) and run the finalize function, which is
	// generally in charge of binding new command instances, with blank flags.
	if onFinalize != nil {
		storage = make(_storage)
		onFinalize()
	}

	return filtered, a.meta
}
