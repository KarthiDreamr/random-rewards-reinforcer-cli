# random-rewards-reinforcer-cli

A Go CLI that starts a surprise random work timer and plays an MP3 reward alert between random intervals (default: 20-35 minutes).

## Run

```bash
random-rewards-reinforcer
```

## Local build

```bash
go build -o random-rewards-reinforcer .
```

## Flags

- `-min` minimum interval units (default: 20)
- `-max` maximum interval units (default: 35)
- `-audio` alert audio file path
- `-player` force a specific player command
- `-test-seconds` use seconds instead of minutes for testing
