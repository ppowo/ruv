package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

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

	// Validate station
	if err := station.Validate(); err != nil {
		return fmt.Errorf("invalid station: %w", err)
	}

	fmt.Printf("🎵 Playing %s...\n", station.Name)

	// Create streamer
	streamer := audio.NewStreamer()
	defer streamer.Close()

	// Start streaming using ffmpeg
	err = streamer.Stream(station.URL)
	if err != nil {
		return fmt.Errorf("failed to start stream: %w", err)
	}

	// Handle Ctrl+C gracefully
	handleSignals(streamer)

	return nil
}

// handleSignals waits for Ctrl+C and cleans up
func handleSignals(streamer *audio.Streamer) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("\nPress Ctrl+C to stop")

	<-sigChan

	fmt.Println("\n\n⏹ Stopping stream...")
	streamer.Close()
	fmt.Println("Stream stopped.")
}

// Execute executes the root command
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// Add root command flags here if needed
}

// GetRootCmd returns the root command
func GetRootCmd() *cobra.Command {
	return rootCmd
}
