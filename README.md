# ruv

A lightweight command-line radio streamer for internet radio stations.

## Features

- Stream internet radio directly from your terminal
- Support for multiple formats (ICY/MP3, HLS/M3U8)
- Simple 4-character station codes
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
ruv nood     # Noods Radio
ruv drmm     # Intergalactic FM Dream Machine
ruv 9128     # 9128.live
ruv alha     # Radio Alhara

# Stop playback
Ctrl+C
```

## Available Stations

| Code | Station |
|------|---------|
| `reso` | Resonance FM |
| `rese` | Resonance Extra |
| `ntso` | NTS Radio 1 |
| `ntst` | NTS Radio 2 |
| `nood` | Noods Radio |
| `drmm` | Intergalactic FM Dream Machine |
| `9128` | 9128.live |
| `alha` | Radio Alhara |
## How It Works

`ruv` uses ffmpeg to handle stream decoding and the [oto](https://github.com/ebitengine/oto) library for cross-platform audio playback. This approach supports a wide variety of streaming formats without complex codec implementation.

## License

See LICENSE file for details.
