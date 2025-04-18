package main

//goland:noinspection GoNameStartsWithPackageName
const (
	appName   = "say-yes"
	version   = "1.0.0"
	author    = `The Examples Authors`
	desc      = `a good blueprint for you.`
	longDesc  = `a good blueprint for you.`

To get help for say-yes building options, run 
'say-yes --help', or 'say-yes -h'.
`
	examples = `
$ {{.AppName}} gen shell [--bash|--zsh|--auto]
  generate bash/shell completion scripts
$ {{.AppName}} gen man
  generate linux man page 1
$ {{.AppName}} --help
  show help screen.
`
	overview = `` //nolint:varcheck

	zero = 0

	defaultTraceEnabled  = true //nolint:varcheck
	defaultDebugEnabled  = false
	defaultLoggerLevel   = "info"
	defaultLoggerBackend = "logrus"
)
