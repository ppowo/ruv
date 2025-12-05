# ruv

A lightweight command-line radio streamer for internet radio stations.

## Features

- Stream internet radio directly from your terminal
- Support for multiple formats (ICY/MP3, HLS/M3U8)
- Simple 4-letter station codes
- Clean resource management (no memory leaks)

## Requirements

- Go 1.25.4+
- [ffmpeg](https://ffmpeg.org/) (must be installed and in PATH)

## Installation

```bash
# Clone the repository
git clone https://github.com/ppowo/ruv.git
cd ruv

# Build
go build -o bin/ruv .

# Or use Mage
mage build

# Install to ~/.bio/bin or ~/.local/bin
mage install
```

## Usage

```bash
# List all available stations
ruv

# Play a station
ruv reso     # Resonance FM
ruv rese     # Resonance Extra
ruv ntso     # NTS Radio 1
ruv ntst     # NTS Radio 2
ruv lyll     # LYL Radio
ruv cash     # Cashmere Radio
ruv lake     # The Lake Radio
ruv alha     # Radio Alhara

# Stop playback
Ctrl+C
```

## Available Stations

| Code | Station | Location |
|------|---------|----------|
| `reso` | Resonance FM | London |
| `rese` | Resonance Extra | London |
| `ntso` | NTS Radio 1 | Global |
| `ntst` | NTS Radio 2 | Global |
| `lyll` | LYL Radio | London |
| `cash` | Cashmere Radio | Berlin |
| `lake` | The Lake Radio | Online |
| `alha` | Radio Alhara | Palestine |

## How It Works

`ruv` uses ffmpeg to handle stream decoding and the [oto](https://github.com/ebitengine/oto) library for cross-platform audio playback. This approach supports a wide variety of streaming formats without complex codec implementation.

## License

See LICENSE file for details.
