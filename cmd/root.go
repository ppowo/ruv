package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/eiannone/keyboard"
	"github.com/ppowo/ruv/audio"
	"github.com/ppowo/ruv/stations"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ruv [station]",
	Short: "Stream internet radio stations",
	Long: `Stream internet radio stations using ffmpeg.

Usage:
  ruv          List all available stations
  ruv reso     Play Resonance FM
  ruv rese     Play Resonance Extra
  ruv ntso     Play NTS Radio 1
  ruv ntst     Play NTS Radio 2
  ruv lyll     Play LYL Radio
  ruv cash     Play Cashmere Radio
  ruv lake     Play The Lake Radio
  ruv alha     Play Radio Alhara`,
	Args: cobra.MaximumNArgs(1),
	RunE: runRoot,
}

func runRoot(cmd *cobra.Command, args []string) error {
	// If no arguments, show station list
	if len(args) == 0 {
		return runList()
	}

	// Otherwise, play the station
	stationCode := args[0]

	// Look up the station
	station, err := stations.GetStation(stationCode)
	if err != nil {
		return fmt.Errorf("station not found: %w", err)
	}

	fmt.Printf("🎵 Playing %s...\n", station.Description)

	// Create streamer
	streamer := audio.NewStreamer()
	defer streamer.Close()

	// Start streaming using ffmpeg
	err = streamer.Stream(station.URL)
	if err != nil {
		return fmt.Errorf("failed to start stream: %w", err)
	}

	// Handle both keyboard input and signals
	handleEvents(streamer)

	return nil
}

// handleEvents manages keyboard input and signals
func handleEvents(streamer *audio.Streamer) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Initialize keyboard
	if err := keyboard.Open(); err != nil {
		fmt.Printf("Warning: keyboard input unavailable: %v\n", err)
		fmt.Println("\nPress Ctrl+C to stop")
		<-sigChan
		stopStream(streamer)
		return
	}
	defer keyboard.Close()

	fmt.Println("Controls: [Space/P] Pause/Resume | [Ctrl+C] Exit")

	// Channels for keyboard events and exit signal
	keyChan := make(chan rune, 1)
	exitChan := make(chan bool, 1)

	// Start keyboard listener in separate goroutine
	go func() {
		for {
			char, key, err := keyboard.GetKey()
			if err != nil {
				return
			}

			// Check for Ctrl+C (can be char code 3 or KeyCtrlC depending on terminal)
			if char == 3 || key == keyboard.KeyCtrlC {
				exitChan <- true
				return
			}

			// Send relevant keys to channel
			if char == ' ' || char == 'p' || char == 'P' || key == keyboard.KeySpace {
				keyChan <- char
			}
		}
	}()

	// Event loop - handle signals and keyboard input
	for {
		select {
		case <-sigChan:
			stopStream(streamer)
			return

		case <-exitChan:
			stopStream(streamer)
			return

		case <-keyChan:
			// Handle pause/unpause
			if streamer.IsPaused() {
				if err := streamer.Unpause(); err != nil {
					fmt.Printf("\nError resuming: %v\n", err)
				} else {
					fmt.Print("\r⏯ Resuming stream...                    ")
				}
			} else {
				if err := streamer.Pause(); err != nil {
					fmt.Printf("\nError pausing: %v\n", err)
				} else {
					fmt.Print("\r⏸ Stream paused (press Space/P to resume)")
				}
			}
		}
	}
}

func stopStream(streamer *audio.Streamer) {
	fmt.Println("\n\n⏹ Stopping stream...")
	streamer.Close()
	fmt.Println("Stream stopped.")
}

// Execute executes the root command
func Execute() error {
	return rootCmd.Execute()
}


// GetRootCmd returns the root command
func GetRootCmd() *cobra.Command {
	return rootCmd
}
