package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var (
	pomodoroDuration        time.Duration
	breakDuration           time.Duration
	updateInterval          = time.Second
	notificationDuration    string
	warningPomodoroDuration time.Duration
	warningBreakDuration    time.Duration

	// Default durations
	defaultPomodoroDuration        = 1 * time.Minute
	defaultBreakDuration           = 2 * time.Minute
	defaultNotificationDuration    = "1800000" // 30 minutes, time is in milliseconds
	defaultPomodoroWarningDuration = 1 * time.Minute
	defaultBreakWarningDuration    = 1 * time.Minute

	pomodoroWarningSent bool // Flag to track if Pomodoro warning is sent
	breakWarningSent    bool // Flag to track if Break warning is sent
)

func main() {
	var rootCmd = &cobra.Command{Use: "pomo",
		Short: "A Pomodoro timer",
		Long: `A Pomodoro timer helps you manage your work time effectively. It consists of work intervals
(pomodoros) followed by short breaks. The timer will notify you when it's time to take a break
and when the break is over. You can customize the Pomodoro and break durations as well as
the notification duration.

Written by: Neoplatonist`,
	}

	// Define command-line flags for Pomodoro, break, and notification durations
	rootCmd.Flags().DurationVarP(&pomodoroDuration, "pomodoro", "p", defaultPomodoroDuration, "Pomodoro duration")
	rootCmd.Flags().DurationVarP(&breakDuration, "break", "b", defaultBreakDuration, "Break duration")
	rootCmd.Flags().StringVarP(&notificationDuration, "notification", "t", defaultNotificationDuration, "Notification duration (in milliseconds)")
	rootCmd.Flags().DurationVarP(&warningPomodoroDuration, "warning", "m", defaultPomodoroWarningDuration, "Warning time (in minutes)")
	rootCmd.Flags().DurationVarP(&warningBreakDuration, "warning-break", "k", defaultBreakWarningDuration, "Warning time for break (in minutes)")

	var startCmd = &cobra.Command{
		Use:   "start",
		Short: "Start a Pomodoro timer",
		Run:   startPomodoro,
	}

	rootCmd.AddCommand(startCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

func startPomodoro(cmd *cobra.Command, args []string) {
	// Clear the screen
	clearScreen()

	remainingTime := pomodoroDuration

	for remainingTime > 0 {
		progress := float64(pomodoroDuration-remainingTime) / float64(pomodoroDuration)
		loadingBar := createLoadingBar(progress)

		fmt.Println("Pomodoro Timer")
		fmt.Println()
		fmt.Printf("Time remaining: %s\n%s\n", formatDuration(remainingTime), loadingBar)

		// Check if it's time to send the warning notification and the warning hasn't been sent
		if remainingTime <= warningPomodoroDuration && !pomodoroWarningSent {
			sendNotification("Pomodoro Warning", fmt.Sprintf("Pomodoro ends in %s.", formatDuration(remainingTime)))
			pomodoroWarningSent = true
		}

		time.Sleep(updateInterval)
		remainingTime -= updateInterval
		clearScreen()
	}

	// Send a Pomodoro completed notification with adjustable duration and current time
	err := sendNotification("Pomodoro completed! "+time.Now().Format("15:04pm"), "Take a break.")
	if err != nil {
		fmt.Println("Error sending notification:", err)
	}

	// Wait for user input (Enter key) to start the break
	waitForEnter()

	// Clear the screen
	clearScreen()

	remainingTime = breakDuration
	breakWarningSent = false // Reset Break warning flag

	for remainingTime > 0 {
		progress := float64(breakDuration-remainingTime) / float64(breakDuration)
		loadingBar := createLoadingBar(progress)

		fmt.Println("Break Time")
		fmt.Println()
		fmt.Printf("Time remaining: %s\n%s\n", formatDuration(remainingTime), loadingBar)

		// Check if it's time to send the warning notification and the warning hasn't been sent
		if remainingTime <= warningBreakDuration && !breakWarningSent {
			sendNotification("Break Warning", fmt.Sprintf("Break ends in %s.", formatDuration(remainingTime)))
			breakWarningSent = true
		}

		time.Sleep(updateInterval)
		remainingTime -= updateInterval
		clearScreen()
	}

	// Send a break is over notification with adjustable duration and current time
	err = sendNotification("Break is over! "+time.Now().Format("15:04pm"), "Time to start the next Pomodoro.")
	if err != nil {
		fmt.Println("Error sending notification:", err)
	}

	// Wait for user input (Enter key) to start the next Pomodoro
	waitForEnter()

	// Recursively start the next Pomodoro
	startPomodoro(cmd, args)
}

func formatDuration(d time.Duration) string {
	minutes := int(d.Minutes())
	seconds := int(d.Seconds()) % 60
	return fmt.Sprintf("%02d:%02d", minutes, seconds)
}

func createLoadingBar(progress float64) string {
	barWidth := 30
	completed := int(progress * float64(barWidth))
	loadingBar := strings.Repeat("█", completed) + strings.Repeat(" ", barWidth-completed)
	return "[" + loadingBar + "]"
}

func sendNotification(title, message string) error {
	cmd := exec.Command("notify-send", title, message, "-t", notificationDuration)
	return cmd.Run()
}

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func waitForEnter() {
	fmt.Print("Press Enter to continue...")
	fmt.Scanln()
}
