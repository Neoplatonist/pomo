# Pomodoro Timer in Go

A simple Pomodoro timer command-line application written in Go.

## Requirements

-   Go 1.16 or higher
-   Linux-based operating system (for notifications using `notify-send`)

## Overview

This Pomodoro timer helps you manage your work time effectively by breaking it into intervals, traditionally 25 minutes in length, known as "Pomodoros," separated by short breaks.

## Features

-   Customizable Pomodoro and break durations.
-   Notifications for Pomodoro completion and break time.
-   Adjustable notification duration.
-   Optional warnings before Pomodoro and break end.
-   Easy-to-use command-line interface.

## Usage

### Installation

You can build and install the Pomodoro timer using the provided Makefile:

```shell
make install
```

This will compile the binary and place it in your `/usr/local/bin` directory for easy access. (It uses `sudo mv`, so you will have to provide your password)

### Starting a Pomodoro

To start a Pomodoro timer with default settings (25 minutes Pomodoro, 5 minutes break. Notifications have a 30 minute duration. You get a 5 minute warning for the Pomodoro session ending, and you get 1 minute warning for a break ending), simply run:

```shell
pomo start
```

You can also customize the timer durations and notification settings using command-line flags:

```shell
pomo start --pomodoro 30m --break 10m --notification 1800000 --warning 5m
```

### Exiting the Timer

The timer will display a notification when it's time for a break or when the Pomodoro is complete. You can press Enter to continue to the next phase of the timer.

## Author

[Neoplatonist](https://github.com/neoplatonist)
