# gonductor

[![Build Status](https://travis-ci.org/itsmeduncan/gonductor.svg?branch=master)](https://travis-ci.org/itsmeduncan/gonductor) [![Code Climate](https://codeclimate.com/github/itsmeduncan/gonductor/badges/gpa.svg)](https://codeclimate.com/github/itsmeduncan/gonductor)

MTA subway line status command line tool

## Usage

### Command line interface
```
$ gonductor --help
NAME:
gonductor - Simple tool for MTA subway status

USAGE:
gonductor [global options] command [command options]

VERSION:
1.0.0

AUTHOR:
Author - <unknown@email>

COMMANDS:
help, h      Shows a list of commands or help for on

GLOBAL OPTIONS:
--line, -l           subway line to check the status
--tmux, -t           turn tmux colorization on
--help, -h           show help
--version, -v        print the version
```

Getting the status of a specifc line
```
$ gonductor -l NQR
#=> GOOD SERVICE
```

Tmux formatting the status of a specific line
```
gonductor -t --line=NQR
#=> #[fg=colour3]NQR
```
