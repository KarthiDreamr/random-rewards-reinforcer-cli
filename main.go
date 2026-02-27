package main

import (
	"context"
	"flag"
	"fmt"
	"math/rand/v2"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

const defaultAudioPath = "/home/karthidreamr/Music/Urban: A Bit Cooler, More Chill - MiniMax.mp3"

const (
	colorReset  = "\033[0m"
	colorCyan   = "\033[36m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorRed    = "\033[31m"
)

type config struct {
	minMinutes int
	maxMinutes int
	audioPath  string
	player     string
	testMode   bool
}

func main() {
	cfg := parseFlags()
	if err := validateConfig(cfg); err != nil {
		printMsg(colorRed, "❌", fmt.Sprintf("Invalid config: %v", err))
		os.Exit(1)
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	unit := "minutes"
	if cfg.testMode {
		unit = "seconds"
	}
	printMsg(colorCyan, "🎯", fmt.Sprintf("Random Rewards Reinforcer started (%d-%d %s)", cfg.minMinutes, cfg.maxMinutes, unit))
	printMsg(colorGreen, "💻", "Focus on your work. Reward alert will surprise you.")
	printMsg(colorYellow, "🛑", "Press Ctrl+C to stop.")

	for cycle := 1; ; cycle++ {
		waitDuration := randomDuration(cfg.minMinutes, cfg.maxMinutes, cfg.testMode)

		select {
		case <-ctx.Done():
			printMsg(colorYellow, "👋", "Stopped. Nice work today.")
			return
		case <-time.After(waitDuration):
			printMsg(colorGreen, "📺", fmt.Sprintf("Cycle %d reward unlocked! Enjoy your episode.", cycle))
			if err := playAudio(ctx, cfg.audioPath, cfg.player); err != nil {
				printMsg(colorRed, "🔇", fmt.Sprintf("Audio failed: %v", err))
			}
			printMsg(colorCyan, "⚙️", "Back to deep work until the next surprise reward.")
		}
	}
}

func printMsg(color, emoji, message string) {
	fmt.Printf("%s%s %s%s\n", color, emoji, message, colorReset)
}

func parseFlags() config {
	cfg := config{}
	flag.IntVar(&cfg.minMinutes, "min", 20, "minimum random interval in minutes")
	flag.IntVar(&cfg.maxMinutes, "max", 35, "maximum random interval in minutes")
	flag.StringVar(&cfg.audioPath, "audio", defaultAudioPath, "audio file path to play when timer finishes")
	flag.StringVar(&cfg.player, "player", "", "force audio player command (optional)")
	flag.BoolVar(&cfg.testMode, "test-seconds", false, "use seconds instead of minutes for quick testing")
	flag.Parse()
	return cfg
}

func validateConfig(cfg config) error {
	if cfg.minMinutes <= 0 || cfg.maxMinutes <= 0 {
		return fmt.Errorf("min and max must be > 0")
	}
	if cfg.minMinutes > cfg.maxMinutes {
		return fmt.Errorf("min cannot be greater than max")
	}
	if _, err := os.Stat(cfg.audioPath); err != nil {
		return fmt.Errorf("audio file not found: %w", err)
	}
	return nil
}

func randomDuration(minUnits, maxUnits int, testMode bool) time.Duration {
	rangeSize := maxUnits - minUnits + 1
	value := minUnits + rand.IntN(rangeSize)
	if testMode {
		return time.Duration(value) * time.Second
	}
	return time.Duration(value) * time.Minute
}

func playAudio(ctx context.Context, audioPath, forcedPlayer string) error {
	playerCmd, args, err := choosePlayer(audioPath, forcedPlayer)
	if err != nil {
		return err
	}

	cmd := exec.CommandContext(ctx, playerCmd, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func choosePlayer(audioPath, forcedPlayer string) (string, []string, error) {
	if forcedPlayer != "" {
		if _, err := exec.LookPath(forcedPlayer); err != nil {
			return "", nil, fmt.Errorf("player %q not found in PATH", forcedPlayer)
		}
		return forcedPlayer, []string{audioPath}, nil
	}

	candidates := [][]string{}
	switch runtime.GOOS {
	case "darwin":
		candidates = append(candidates, []string{"afplay", audioPath})
	case "linux":
		candidates = append(candidates,
			[]string{"ffplay", "-nodisp", "-autoexit", "-loglevel", "error", audioPath},
			[]string{"mpg123", "-q", audioPath},
		)
	}
	candidates = append(candidates, []string{"ffplay", "-nodisp", "-autoexit", "-loglevel", "error", audioPath}, []string{"mpg123", "-q", audioPath})

	for _, entry := range candidates {
		if _, err := exec.LookPath(entry[0]); err == nil {
			return entry[0], entry[1:], nil
		}
	}

	return "", nil, fmt.Errorf("no supported audio player found; install ffplay or mpg123")
}
