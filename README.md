# gonductor

MTA subway line status command line tool

## Usage

```
gonductor --line=NQR

NQR GOOD SERCIVE
```

## Tmux Status Bar

```
set -g status-right '#(gonductor -t --line=123)'
```
