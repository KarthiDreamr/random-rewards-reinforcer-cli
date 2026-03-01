# ⏱️ Random Rewards Reinforcer CLI

A simple, elegant Go CLI timer that plays a reward song after a random interval. Perfect for timing focused work, study sessions, or any activity where you want periodic audio alerts.

## Features

- 🎲 **Random Timer**: Set a random interval (default 20-35 minutes) and get surprised by an audio alert
- 📊 **Live Elapsed Time**: Watch elapsed time update in real-time on the same line
- 🎵 **Custom Audio**: Play any MP3 file when timer completes
- ⚡ **Lightweight**: Fast, minimal CLI with no dependencies
- 🧪 **Test Mode**: Use seconds instead of minutes for quick testing
- 🎯 **Cross-platform**: Works on Linux, macOS, and other Unix-like systems

## Installation

### From Source

```bash
git clone https://github.com/KarthiDreamr/random-rewards-reinforcer-cli.git
cd random-rewards-reinforcer-cli
go build -o rrr
sudo mv rrr /usr/local/bin/
```

### Quick Install

```bash
go install github.com/KarthiDreamr/random-rewards-reinforcer-cli@latest
```

## Usage

### Basic Usage

```bash
rrr
```

Starts a timer with default interval (20-35 minutes). The timer will display elapsed time and play an audio alert when complete.

### Custom Interval

```bash
rrr -min 10 -max 30
```

Sets a random interval between 10 and 30 minutes.

### Custom Audio File

```bash
rrr -audio /path/to/your/audio.mp3
```

### Force Specific Audio Player

```bash
rrr -player mpg123
```

### Test Mode (Quick Testing)

```bash
rrr -test-seconds -min 2 -max 5
```

Uses seconds instead of minutes for quick testing (interval 2-5 seconds).

### Stop the Timer

Press `Ctrl+C` to stop the timer at any time.

## Command Line Flags

| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `-min` | int | 20 | Minimum timer interval (minutes or seconds in test mode) |
| `-max` | int | 35 | Maximum timer interval (minutes or seconds in test mode) |
| `-audio` | string | `~/Music/Urban: A Bit Cooler, More Chill - MiniMax.mp3` | Path to audio file to play when timer finishes |
| `-player` | string | (auto-detect) | Force a specific audio player command |
| `-test-seconds` | bool | false | Use seconds instead of minutes for quick testing |

## Output Example

```
⏱️ Timer set for 23 minutes
⏰    12 minutes
```

The elapsed time updates in real-time on the same line without creating new output lines.

## Requirements

- **Go 1.21+** (for building from source)
- **Audio Player**: One of the following installed:
  - `ffplay` (part of FFmpeg)
  - `mpg123`
  - `afplay` (macOS, built-in)

### Install Audio Player (Linux)

```bash
# Ubuntu/Debian
sudo apt-get install ffmpeg

# Or
sudo apt-get install mpg123

# Fedora
sudo dnf install ffmpeg

# Or
sudo dnf install mpg123
```

## Configuration

Default audio file path: `/home/karthidreamr/Music/Urban: A Bit Cooler, More Chill - MiniMax.mp3`

To use a different default audio file, modify the `defaultAudioPath` constant in `main.go` and rebuild.

## How It Works

1. On startup, a random duration is calculated between `-min` and `-max`
2. A timer begins, displaying elapsed time that updates every minute (or second in test mode)
3. Elapsed time is displayed with a carriage return (`\r`), updating in place
4. When the timer completes, the audio file is played
5. The program exits

## Contributing

Feel free to fork and submit pull requests for improvements!

## License

MIT
