# 🎯 Random Rewards Reinforcer CLI

A simple, elegant Go CLI timer that plays a reward song after a **mysterious random interval**. The exact time is hidden from you—you only know the reward will arrive between your set bounds (default 20-35 minutes). Perfect for maintaining focus without clock-watching!

## Philosophy

**Don't know when the reward is coming.** The point is to surprise yourself with an audio alert and enjoy not watching the clock. You'll see how much time has *elapsed* since you started, but you won't know how much time *remains*—that's the mystery, that's the motivation.

## Features

- 🎲 **Mystery Timer**: Set a random interval (default 20-35 minutes) but DON'T KNOW the exact duration
- ⏰ **Elapsed Time Only**: Watch elapsed time update in real-time on the same line (no countdown spoilers!)
- 💡 **Mystery Hint**: At a random point before the reward, get a clue—whether it arrives at an ODD or EVEN minute
- 🎵 **Custom Audio**: Play any MP3 file when the reward arrives
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

### Build & Install (latest from this repo)

Build the binary (any name is fine):

```bash
go build -o random-rewards-reinforcer
```

**User install** (no sudo; ensure `~/.local/bin` is in your PATH):

```bash
cp random-rewards-reinforcer ~/.local/bin/
```

**System-wide install** as `rrr`:

```bash
sudo cp random-rewards-reinforcer /usr/local/bin/rrr
```

**Using an alias:** If you use `rrr` as a shell alias, you don't need to build the binary with the name `rrr`. Build with any name (e.g. `random-rewards-reinforcer`), install it, and point your alias to that path—e.g. `alias rrr='random-rewards-reinforcer'` or `alias rrr='~/.local/bin/random-rewards-reinforcer'`.

### Quick Install

```bash
go install github.com/KarthiDreamr/random-rewards-reinforcer-cli@latest
```

## Usage

### Basic Usage

```bash
rrr
```

Starts a mystery timer with default interval (20-35 minutes). You'll see how much time has elapsed, but not how long until the reward.

**Output:**
```
🎯 Random Rewards Reinforcer started
🎁 Schrodinger reward arrives between 20-35 minutes
🛑 Press Ctrl+C to stop.
⏰    23 minutes
```

### Custom Interval

```bash
rrr -min 10 -max 30
```

Sets a random interval between 10 and 30 minutes (you won't know which).

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

Uses seconds instead of minutes for quick testing (mystery interval 2-5 seconds).

### Stop the Timer

Press `Ctrl+C` to stop the timer at any time.

## Command Line Flags

| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `-min` | int | 20 | Minimum timer interval (minutes or seconds in test mode) |
| `-max` | int | 35 | Maximum timer interval (minutes or seconds in test mode) |
| `-audio` | string | `~/Music/Urban: A Bit Cooler, More Chill - MiniMax.mp3` | Path to audio file to play when reward arrives |
| `-player` | string | (auto-detect) | Force a specific audio player command |
| `-test-seconds` | bool | false | Use seconds instead of minutes for quick testing |

## Output Example

```
🎯 Random Rewards Reinforcer started
🎁 Schrodinger reward arrives between 20-35 minutes
🛑 Press Ctrl+C to stop.
⏰     0 minutes
⏰     1 minutes
⏰     2 minutes
...
⏰     8 minutes
💡 Hint: Reward will arrive at an EVEN minute!
⏰     9 minutes
...
⏰    23 minutes
🎵 Time's up!
```

**Note:** You only see elapsed time, never the remaining time. The mystery is maintained throughout! At a random point before the reward, you'll get a hint about whether the final minute is odd or even.

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

1. On startup, a random **reward time** is calculated between `-min` and `-max` (kept secret!)
2. A second random **hint time** is generated (between 0 and the reward time)
3. Display messages show the bounds but not the exact timer
4. A timer begins, displaying **only elapsed time** that updates every minute (or second in test mode)
5. At the **hint time**, a clue appears: whether the reward arrives at an ODD or EVEN minute
6. Elapsed time is displayed with a carriage return (`\r`), updating in place
7. When the timer completes, the audio file is played
8. The program exits

## The Philosophy Behind It

Schrödinger's reward: Until the audio plays, you exist in a superposition of states—it could arrive now, or in 35 minutes. This uncertainty keeps you engaged without the anxiety of watching a countdown. Focus on your work, not the clock.

The **mystery hint** adds another layer—you'll get one strategic piece of information (odd or even) at an unpredictable time, giving you just enough hope to keep going without spoiling the surprise.

## Contributing

Feel free to fork and submit pull requests for improvements!

## License

MIT
