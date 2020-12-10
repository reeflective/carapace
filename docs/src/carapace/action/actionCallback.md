# ActionCallback
[ActionCallback](https://pkg.go.dev/github.com/rsteube/carapace#ActionCallback) is a special action where the program itself provides the completion dynamically. For this the [hidden subcommand](../hiddenSubcommand.md) is called with an [uid](../uid.md) and the current command line content which then lets cobra parse existing flags and invokes the callback function after that.

```go
carapace.ActionCallback(func(args []string) carapace.Action {
  if conditionCmd.Flag("required").Value.String() == "valid" {
    return carapace.ActionValues("condition fulfilled")
  } else {
    return carapace.ActionMessage("flag --required must be set to valid: " + conditionCmd.Flag("required").Value.String())
  }
})
```

## Testing

Since callbacks are simply invocations of the program they can be tested directly.
```sh
./example _carapace bash '_example__condition#1' example condition --required invalid
#compgen -W $'_\nERR (flag --required must be set to valid: invalid)' -- "${cur//\\ / }" | sed "s!^${curprefix//\\ / }!!"

./example _carapace elvish '_example__condition#1' example condition --required invalid
#edit:complex-candidate ERR &display-suffix=' (flag --required must be set to valid: invalid)'
#edit:complex-candidate _ &display-suffix=' ()'

./example _carapace fish '_example__condition#1' example condition --required invalid
#echo -e ERR\tflag --required must be set to valid: invalid\n_\t\n\n

./example _carapace powershell '_example__condition#1' example condition --required invalid
#[CompletionResult]::new('ERR', 'ERR', [CompletionResultType]::ParameterValue, ' ')
#[CompletionResult]::new('flag --required must be set to valid: invalid', 'flag --required must be set to valid: invalid', [CompletionResultType]::ParameterValue, ' ')

./example _carapace xonsh '_example__condition#1' example condition --required invalid
#{
#  RichCompletion('_', display='_', description='flag --required must be set to valid: invalid', prefix_len=0),
#  RichCompletion('ERR', display='ERR', description='flag --required must be set to valid: invalid', prefix_len=0),
#}

./example _carapace zsh '_example__condition#1' example condition --required invalid
#{local _comp_desc=('_' 'ERR (flag --required must be set to valid: invalid)');compadd -S '' -d _comp_desc '_' 'ERR'}
```