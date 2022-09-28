package state

import (
	"reflect"
	"unsafe"

	"github.com/davecgh/go-spew/spew"
	"github.com/rsteube/carapace/internal/assert"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
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

type flagState struct {
	Changed bool
	Value   string
}

type CommandState struct {
	*cobra.Command
	t      T
	before string
	flags  map[*pflag.Flag]flagState
}

func NewCommandState(t T, cmd *cobra.Command) *CommandState {
	cmd.CompletionOptions = cobra.CompletionOptions{
		DisableDefaultCmd: true,
	}
	cmd.InitDefaultHelpCmd()
	cmd.Flags() // call once
	cfg := spew.NewDefaultConfig()
	cfg.SortKeys = true
	state := &CommandState{
		Command: cmd,
		t:       t,
		//before:  strings.Replace(fmt.Sprintf("%+v", cmd), " ", ", \n", -1),
		before: cfg.Sdump(cmd),
		flags:  make(map[*pflag.Flag]flagState),
	}
	state.store(cmd)
	return state
}

func (c *CommandState) store(cmd *cobra.Command) {
	cmd.LocalFlags().VisitAll(func(f *pflag.Flag) {
		c.flags[f] = flagState{Value: f.Value.String()}
	})

	for _, subCmd := range cmd.Commands() {
		c.store(subCmd)
	}
}

func (c *CommandState) Restore() {
	c.LocalFlags().Args()
	for flag, state := range c.flags {
		flag.Value.Set(state.Value)
		flag.Changed = state.Changed
	}

	if field := reflect.ValueOf(c.Command).Elem().FieldByName("args"); field.IsValid() {
		field = reflect.NewAt(field.Type(), unsafe.Pointer(field.UnsafeAddr())).Elem()
		field.Set(reflect.Zero(field.Type()))
	}

	if field := reflect.ValueOf(c.Command).Elem().FieldByName("sortedFormal"); field.IsValid() {
		field = reflect.NewAt(field.Type(), unsafe.Pointer(field.UnsafeAddr())).Elem()
		field.Set(reflect.Zero(field.Type()))
	}

	if field := reflect.ValueOf(c.Command).Elem().FieldByName("pflags"); field.IsValid() {
		field = reflect.NewAt(field.Type(), unsafe.Pointer(field.UnsafeAddr())).Elem()
		field.Set(reflect.Zero(field.Type()))
	}
	if field := reflect.ValueOf(c.Command).Elem().FieldByName("lflags"); field.IsValid() {
		field = reflect.NewAt(field.Type(), unsafe.Pointer(field.UnsafeAddr())).Elem()
		field.Set(reflect.Zero(field.Type()))
	}
	if field := reflect.ValueOf(c.Command).Elem().FieldByName("iflags"); field.IsValid() {
		field = reflect.NewAt(field.Type(), unsafe.Pointer(field.UnsafeAddr())).Elem()
		field.Set(reflect.Zero(field.Type()))
	}
	if field := reflect.ValueOf(c.Command).Elem().FieldByName("parentsPflags"); field.IsValid() {
		field = reflect.NewAt(field.Type(), unsafe.Pointer(field.UnsafeAddr())).Elem()
		field.Set(reflect.Zero(field.Type()))
	}

	//if field := reflect.ValueOf(c.Command).Elem().FieldByName("commandsAreSorted"); field.IsValid() {
	//	field = reflect.NewAt(field.Type(), unsafe.Pointer(field.UnsafeAddr())).Elem()
	//	field.SetBool(false)
	//}

	if field := reflect.ValueOf(c.Command).Elem().FieldByName("flags"); field.IsValid() {
		if field := field.Elem().FieldByName("args"); field.IsValid() {
			field = reflect.NewAt(field.Type(), unsafe.Pointer(field.UnsafeAddr())).Elem()
			//	field.Set(reflect.Zero(field.Type()))
			field.Set(reflect.MakeSlice(field.Type(), 0, 0))
		}

		if field := field.Elem().FieldByName("sortedFormal"); field.IsValid() {
			field = reflect.NewAt(field.Type(), unsafe.Pointer(field.UnsafeAddr())).Elem()
			field.Set(reflect.Zero(field.Type()))
		}
	}

	//flagSet := reflect.ValueOf(c.Command).Elem().FieldByName("flags")
	//flagSet = reflect.NewAt(flagSet.Type(), unsafe.Pointer(flagSet.UnsafeAddr())).Elem()

	flagSet := c.Command.Flags()
	if field := reflect.ValueOf(flagSet).Elem().FieldByName("parsed"); field.IsValid() {
		field = reflect.NewAt(field.Type(), unsafe.Pointer(field.UnsafeAddr())).Elem()
		field.SetBool(false)
	}
	if field := reflect.ValueOf(flagSet).Elem().FieldByName("args"); field.IsValid() {
		field = reflect.NewAt(field.Type(), unsafe.Pointer(field.UnsafeAddr())).Elem()
		field.Set(reflect.Zero(field.Type()))
	}
	if field := reflect.ValueOf(flagSet).Elem().FieldByName("argsLenAtDash"); field.IsValid() {
		field = reflect.NewAt(field.Type(), unsafe.Pointer(field.UnsafeAddr())).Elem()
		field.SetInt(-1)
	}

	//after := strings.Replace(fmt.Sprintf("%+v", c.Command), " ", ", \n", -1)
	cfg := spew.NewDefaultConfig()
	cfg.SortKeys = true
	after := cfg.Sdump(c.Command)
	assert.Equal(c.t, c.before, after)
}
