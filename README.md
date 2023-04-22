cobra-example
=============

setup
-----
```
$ go mod init github.com/Amitgb14/cobra-example
$ go get -u github.com/spf13/cobra@latest
$ go install github.com/spf13/cobra-cli@latest
$ ~/go/bin/cobra-cli init . -l none -a "Amit Ghadge"
$ ~/go/bin/cobra-cli add ls -a "Amit Ghadge"
```

Run
---
```
$ go run main.go -h
A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.

Usage:
  cobra-example [command]

Available Commands:
  help        Help about any command
  ls          A brief description of your command
  version     A brief description of your command

Flags:
      --config string   config file (default is $HOME/.cobra-example.yaml)
  -h, --help            help for cobra-example
  -t, --toggle          Help message for toggle

Use "cobra-example [command] --help" for more information about a command.
```
