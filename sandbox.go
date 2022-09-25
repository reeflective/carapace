package carapace

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"

	"github.com/rsteube/carapace/internal/assert"
	"github.com/rsteube/carapace/internal/common"
	"github.com/rsteube/carapace/internal/state"
	"github.com/spf13/cobra"
)

type T interface {
	Cleanup(func())
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Fail()
	FailNow()
	Failed() bool
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
	Helper()
	Log(args ...interface{})
	Logf(format string, args ...interface{})
	Name() string
	Setenv(key, value string)
	Skip(args ...interface{})
	SkipNow()
	Skipf(format string, args ...interface{})
	Skipped() bool
	TempDir() string
}

type Sandbox struct {
	t             T
	cmd           *state.CommandState
	dir           string
	mockedReplies map[string]string
	Keep          bool
}

func newSandbox(t T, cmd *cobra.Command) Sandbox {
	os.Mkdir(fmt.Sprintf("%v/carapace-test/", os.TempDir()), 0755)
	tempDir, err := os.MkdirTemp(fmt.Sprintf("%v/carapace-test/", os.TempDir()), "sandbox_"+t.Name()+"_")
	if err != nil {
		t.Fatal("failed to create sandbox dir: " + err.Error())
	}
	return Sandbox{
		t:             t,
		cmd:           state.NewCommandState(t, cmd),
		mockedReplies: make(map[string]string),
		dir:           tempDir,
	}
}

func (s *Sandbox) mockCommand(name string, arg ...string) *exec.Cmd {
	m, _ := json.Marshal(append([]string{name}, arg...))
	if reply, ok := s.mockedReplies[string(m)]; ok {
		return exec.Command("echo", "-n", reply)
	} else {
		return exec.Command("sh", "-c", fmt.Sprintf("echo 'unexpected call:' >/dev/stderr ; false"))
	}
}

func (s *Sandbox) remove() {
	if !s.Keep {
		// TODO cleanup folder
	}
}

func (s *Sandbox) Files(args ...string) {
	if len(args)%2 != 0 {
		s.t.Errorf("invalid amount of arguments: %v", len(args))
	}

	if !strings.HasPrefix(s.dir, os.TempDir()) {
		s.t.Errorf("sandbox dir not in os.TempDir: ", s.dir)
	}

	for i := 0; i < len(args); i += 2 {
		file := args[i]
		content := args[i+1]

		if strings.HasPrefix(file, "../") {
			s.t.Fatalf("invalid filename: %v", file)
		}

		path := fmt.Sprintf("%v/%v", s.dir, file)

		if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil && !os.IsExist(err) {
			s.t.Fatal(err.Error())
		}

		if err := os.WriteFile(path, []byte(content), 0644); err != nil {
			s.t.Fatal(err.Error())
		}
	}

} // TODO
func (s *Sandbox) Reply(args ...string) reply {
	m, _ := json.Marshal(args)
	return reply{s, string(m)}
}

type reply struct {
	*Sandbox
	call string
}

func (r reply) With(s string) {
	r.mockedReplies[r.call] = s
}

func (s *Sandbox) Run(args ...string) run {
	c := Context{
		Args:          []string{},
		CallbackValue: "",
		Dir:           s.dir,
		sandbox:       s,
	}
	if len(args) > 0 {
		c.Args = args[:len(args)-1]
		c.CallbackValue = args[len(args)-1]
	}

	return run{
		c,
        ActionCallback(func(c Context) Action {
	        s.cmd.Restore() // TODO restore right before ActionExecute
		    return ActionExecute(s.cmd.Command)
        }),
	}
}

type run struct {
	context Context
	actual  Action
}

func (r run) Expect(e Action) {
	ar := common.RawValues(r.actual.Invoke(r.context).rawValues).FilterPrefix(r.context.CallbackValue)
	sort.Sort(common.ByValue(ar))
	actual, err := json.MarshalIndent(ar, "", "  ")
	if err != nil {
		r.context.sandbox.t.Fatal(err.Error())
	}

	er := common.RawValues(e.Invoke(r.context).rawValues).FilterPrefix(r.context.CallbackValue)
	sort.Sort(common.ByValue(er))
	expected, err := json.MarshalIndent(er, "", "  ")
	if err != nil {
		r.context.sandbox.t.Fatal(err.Error())
	}

	assert.Equal(r.context.sandbox.t, string(expected), string(actual))
}
