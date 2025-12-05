package audio

import (
	"fmt"
	"os/exec"
	"sync"

	"github.com/ebitengine/oto/v3"
)

// Streamer handles streaming audio from a URL and playing it
type Streamer struct {
	player     *oto.Player
	context    *oto.Context
	cmd        *exec.Cmd
	url        string       // Store URL for restart on unpause
	isPaused   bool         // Track pause state
	pauseMutex sync.RWMutex // Protect concurrent access
}

// NewStreamer creates a new audio streamer
func NewStreamer() *Streamer {
	return &Streamer{}
}

// Stream streams audio from a URL using ffmpeg
func (s *Streamer) Stream(url string) error {
	// Store URL for pause/unpause functionality
	s.url = url

	// Check if ffmpeg is available
	if _, err := exec.LookPath("ffmpeg"); err != nil {
		return fmt.Errorf("ffmpeg is required but not installed. Please install ffmpeg and try again")
	}

	// Create context options
	opts := &oto.NewContextOptions{
		SampleRate:   44100,
		ChannelCount: 2,
		Format:       oto.FormatSignedInt16LE,
	}

	// Initialize the audio context
	context, readyChan, err := oto.NewContext(opts)
	if err != nil {
		return fmt.Errorf("failed to create audio context: %w", err)
	}
	s.context = context

	// Wait for the context to be ready
	<-readyChan

	// Create ffmpeg command
	// -i url: input from URL
	// -f s16le: output as 16-bit signed little-endian PCM
	// -ar 44100: sample rate 44100 Hz
	// -ac 2: stereo (2 channels)
	// -: output to stdout
	cmd := exec.Command("ffmpeg", "-i", url, "-f", "s16le", "-ar", "44100", "-ac", "2", "-")
	s.cmd = cmd

	// Set up stdout for the raw PCM stream
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return fmt.Errorf("failed to get stdout: %w", err)
	}

	// Start the command
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to start ffmpeg: %w", err)
	}

	// Create player directly with the raw PCM stream
	player := context.NewPlayer(stdout)
	s.player = player

	player.Play()

	return nil
}

// Close stops the playback and cleans up resources
func (s *Streamer) Close() error {
	// Pause the player (no need to close in oto v3.4+)
	if s.player != nil {
		s.player.Pause()
		s.player = nil
	}

	// Kill the ffmpeg process and wait for it to finish
	if s.cmd != nil && s.cmd.Process != nil {
		s.cmd.Process.Kill()
		s.cmd.Wait() // Wait for process to finish to avoid zombies
		s.cmd = nil
	}

	// Close the audio context to release audio device resources
	if s.context != nil {
		s.context.Suspend()
		s.context = nil
	}

	return nil
}

// Pause pauses the stream (stops ffmpeg but keeps context alive)
func (s *Streamer) Pause() error {
	s.pauseMutex.Lock()
	defer s.pauseMutex.Unlock()

	if s.isPaused {
		return nil // Already paused
	}

	// Stop ffmpeg but keep context alive for faster resume
	if s.player != nil {
		s.player.Pause()
	}
	if s.cmd != nil && s.cmd.Process != nil {
		s.cmd.Process.Kill()
		s.cmd.Wait() // Avoid zombie processes
		s.cmd = nil
	}

	s.isPaused = true
	return nil
}

// Unpause resumes the stream from live (restarts ffmpeg)
func (s *Streamer) Unpause() error {
	s.pauseMutex.Lock()
	defer s.pauseMutex.Unlock()

	if !s.isPaused {
		return nil // Not paused
	}

	// Restart ffmpeg (reuse existing context)
	cmd := exec.Command("ffmpeg", "-i", s.url, "-f", "s16le", "-ar", "44100", "-ac", "2", "-")
	s.cmd = cmd

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return fmt.Errorf("failed to get stdout: %w", err)
	}

	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to restart ffmpeg: %w", err)
	}

	// Create new player with existing context
	player := s.context.NewPlayer(stdout)
	s.player = player
	player.Play()

	s.isPaused = false
	return nil
}

// IsPaused returns whether the stream is currently paused
func (s *Streamer) IsPaused() bool {
	s.pauseMutex.RLock()
	defer s.pauseMutex.RUnlock()
	return s.isPaused
}
