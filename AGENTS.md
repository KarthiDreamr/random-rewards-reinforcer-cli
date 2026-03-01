# AGENTS.md - Development Guidelines

Guidelines for agents working on the Random Rewards Reinforcer CLI codebase.

## Build & Run Commands

### Build the binary
```bash
go build -o random-rewards-reinforcer
```

### Build and name as 'rrr'
```bash
go build -o rrr
```

### Install to system PATH
```bash
go build -o random-rewards-reinforcer
cp random-rewards-reinforcer /home/karthidreamr/.local/bin/
sudo cp random-rewards-reinforcer /usr/local/bin/rrr
```

### Run with default settings (20-35 minutes)
```bash
./random-rewards-reinforcer
```

### Run in test mode (seconds instead of minutes)
```bash
./random-rewards-reinforcer -test-seconds -min 2 -max 5
```

### Run with custom interval
```bash
./random-rewards-reinforcer -min 10 -max 20
```

## Code Style Guidelines

### Imports
- Organize imports in standard Go order: stdlib first, then third-party, separated by blank lines
- Use explicit import blocks for clarity (see main.go lines 3-14)
- No vendoring; use module dependencies via go.mod

### Formatting
- Use `gofmt` for all code (Go's standard formatter)
- Run `gofmt -w main.go` before committing
- Indentation: tabs (Go standard)
- Max line length: ~100 chars (pragmatic, not strict)
- Use trailing commas in multi-line struct/slice literals for clean diffs

### Types & Structs
- Use lowercase `config` struct for internal CLI configuration (line 26)
- Exported types (if any) use PascalCase
- Unexported types use camelCase
- Avoid anonymous structs; define named types in this small codebase

### Naming Conventions
- Function names: descriptive verbs + nouns (e.g., `parseFlags`, `validateConfig`, `printMsg`)
- Variable names: short but clear (e.g., `cfg` for config, `ctx` for context)
- Constants: SCREAMING_SNAKE_CASE for unexported constants (e.g., `colorReset`)
- Single-letter variables OK for loop counters, but prefer readability (e.g., `rangeSize`)

### Error Handling
- Always wrap errors with context using `%w` in fmt.Errorf (line 119: `fmt.Errorf("audio file not found: %w", err)`)
- Check errors immediately after operations; don't ignore them
- Use early returns for validation (lines 36-39 in main())
- Use `defer` for cleanup (lines 42, 64, 142)
- For CLI errors, print user-friendly messages before exiting (line 37)

### Comments
- Add comments before complex logic (lines 46, 57, 75)
- Avoid obvious comments; focus on "why" not "what"
- Comment-to-code ratio should be low (code should be self-documenting)

### Function Organization
- Keep functions small and focused (<50 lines ideal)
- Group related functions together
- Order: main() → utilities → validation → execution helpers
- Current codebase: main (line 34) → print functions → parse → validate → random → audio

### Concurrency
- Use context.Context for cancellation (line 41)
- Leverage signal.NotifyContext for graceful shutdown (line 41)
- Use channels properly: `case <-ctx.Done()` for cleanup (line 70)

### String Formatting
- Use Printf/Sprintf for terminal output with ANSI colors
- Color constants are reusable globals (lines 18-24)
- Use emoji for visual hierarchy in CLI output
- Carriage return `\r` for in-place updates (line 88)

### Testing Notes
Currently no automated tests in the repo. For future test additions:
- Use `*testing.T` convention
- Create `main_test.go` for unit tests
- Run: `go test ./...` (all packages)
- Run single test: `go test -v -run TestFunctionName ./`
- Use table-driven tests for multiple scenarios

## Key Architecture Decisions

### Timer Mystery
- Timer duration is calculated but never shown to user (line 44)
- Only elapsed time is displayed, not remaining time (line 85-93)
- This maintains the "Schrodinger reward" UX philosophy

### Audio Player Selection (lines 145-171)
- Attempts platform-specific players first (darwin → afplay, linux → ffplay/mpg123)
- Falls back to generic ffplay/mpg123 on all platforms
- User can force a specific player with `-player` flag
- Error if no player found

### Graceful Shutdown
- Catches SIGINT (Ctrl+C) and SIGTERM via signal.NotifyContext (line 41)
- Cleans up ticker on exit (defer line 64)
- Timer completes naturally or user interrupts (line 70)

## Git & Commit Conventions

- Commit messages: Present tense, descriptive
- Include reasoning in commit body for non-obvious changes
- Keep binary artifacts out of repo (.gitignore lines 2-3)
- Never force push to main; create feature branches if needed

## Dependencies
- Go 1.25.0+ (from go.mod line 3)
- Only stdlib: context, flag, fmt, math/rand/v2, os, os/exec, os/signal, runtime, syscall, time
- No external dependencies
- Keep it lightweight

## Important Files
- `main.go`: Single-file CLI application (172 lines)
- `README.md`: User-facing documentation with philosophy
- `.gitignore`: Excludes binaries and test artifacts

## Common Tasks

### Add a new flag
1. Add field to `config` struct (line 26)
2. Add flag definition in `parseFlags()` (lines 100-109)
3. Add validation in `validateConfig()` if needed (lines 111-122)
4. Use in main() via `cfg.fieldName`

### Change default audio path
Edit `defaultAudioPath` constant (line 16) and rebuild

### Modify timer display
Edit `printElapsed()` function (lines 85-94) to change format/color

### Support new OS
Add case to choosePlayer() switch (lines 154-162)
