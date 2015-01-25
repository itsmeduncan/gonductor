# gonductor

[![Build Status](https://travis-ci.org/itsmeduncan/gonductor.svg?branch=master)](https://travis-ci.org/itsmeduncan/gonductor)

MTA subway line status command line tool

## Usage

### Command line interface
```
gonductor --line=NQR
#=> GOOD SERVICE
```

### In a tmux status bar
```
gonductor -t --line=NQR
#=> #[fg=colour3]NQR
```
