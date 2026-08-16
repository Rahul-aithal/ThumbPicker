# ThumbPicker

[![Go](https://img.shields.io/badge/Go-1.21%2B-blue?logo=go&logoColor=white)](https://golang.org/)
[![FFmpeg](https://img.shields.io/badge/FFmpeg-4.0%2B-orange?logo=ffmpeg&logoColor=white)](https://ffmpeg.org/)

A lightweight video thumbnail selection tool built in Go. Extract frames from videos at specific timestamps using FFmpeg, with plans for a smart UI picker and concurrent processing. Perfect for developers building media tools or learning Go hands-on.

This project documents my journey from Go newbie to functional prototype—check out [Part 1 of the build log](https://medium.com/@aithalrahul34/building-thumbpicker-part-1-extracting-video-frames-with-go-96361f8cb145) for the full story, errors, and lessons learned.

## 🚀 Features

### Current (Phase 1: ✅ Complete)

- Extract single or multiple frames from videos at exact timestamps (e.g., `00:01:23`).
- Convert video duration (seconds) to formatted timestamps (`HH:MM:SS`).
- CLI-based execution via Go's `exec` package for seamless FFmpeg integration.
- Handles basic error checking for slices, paths, and command args.

### Planned

- **Phase 2:** Web or desktop UI for browsing/comparing extracted frames (using Fyne or Gin for web).
- **Phase 3:** Smart frame analysis (quality scoring, auto-detection of key moments) + goroutines for concurrent processing.
- **Phase 4:** Export selected thumbnails, API endpoints, and Docker deployment for production.

## 🛠️ Installation

1. **Prerequisites:**
   - Go 1.21+ ([install here](https://go.dev/doc/install)).
   - FFmpeg 4.0+ ([install here](https://ffmpeg.org/download.html))—tested on macOS/Linux; Windows should work with path setup.
   - sqlc ([install here](https://docs.sqlc.dev/en/stable/overview/install.html)).
   - templ ([install here](https://templ.guide/quick-start/installation)).
2. **Clone the Repo:**

   ```bash
   git clone https://github.com/Rahul-aithal/ThumbPicker.git
   cd ThumbPicker
   ```

3. **Build:**
   ```bash
   go mod init github.com/Rahul-aithal/ThumbPicker  # If not already done
   go mod tidy
   go build -o thumbpicker .
   ```

## 📖 Usage

### Basic Frame Extraction

Run the tool to extract a frame at a specific timestamp from a video.

```bash
./thumbpicker -video input.mp4 -timestamp "00:01:23" -output frame.jpg
```

### Example Code Snippet (Core Extraction Logic)

Here's the heart of Phase 1—feel free to tweak in `main.go`:

```go
package main

import (
    "fmt"
    "os"
    "os/exec"
    "strconv"
    "strings"
)

func secondsToTimestamp(seconds float64) string {
    hours := int(seconds / 3600)
    minutes := int((seconds - float64(hours*3600)) / 60)
    secs := seconds - float64(hours*3600) - float64(minutes*60)

    // Format with leading zeros
    hoursStr := fmt.Sprintf("%02d", hours)
    minutesStr := fmt.Sprintf("%02d", minutes)
    secsStr := fmt.Sprintf("%05.2f", secs)[:5]  // Truncate to SS.XX if needed, but blog uses integer secs

    return hoursStr + ":" + minutesStr + ":" + secsStr
}

func extractFrame(inputVideo, timestamp, outputPath string) error {
    cmd := exec.Command(
        "ffmpeg",
        "-i", inputVideo,
        "-ss", timestamp,
        "-frames:v", "1",
        outputPath,
    )
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    return cmd.Run()
}

func main() {
    // Example usage
    timestamps := []string{"00:01:23", "00:02:45"}  // Use append for dynamic slices!
    for _, ts := range timestamps {
        err := extractFrame("input.mp4", ts, fmt.Sprintf("frame_%s.jpg", strings.ReplaceAll(ts, ":", "_")))
        if err != nil {
            fmt.Printf("Error extracting %s: %v\n", ts, err)
        } else {
            fmt.Printf("Extracted frame for %s\n", ts)
        }
    }
}
```

- **Pro Tip:** Always append to slices (`timestamps = append(timestamps, "00:01:23")`) to avoid index errors—slices aren't fixed like arrays!

### Testing

Drop a sample video (`sample.mp4`) in the root, then run:

```bash
go run main.go
```

Outputs: `frame_00_01_23.jpg`, etc.

## 📁 Project Structure

```
ThumbPicker/
├── main.go              # Entry point with extraction logic
├── go.mod               # Module dependencies
├── go.sum               # Dependency checksums
├── README.md            # You're reading it! 😎
├── sample.mp4           # (Optional) Test video
└── outputs/             # Generated frames (gitignored)
```

## 🗺️ Roadmap

- **Part 1:** Frame Extraction (Done—see blog for the gritty details).
- **Part 2:** UI for frame selection (Coming soon—Fyne? Web? Votes in issues?).
- **Part 3:** Concurrency & Optimization.
- **Part 4:** Full service deployment.

Follow the series on [Medium](https://medium.com/@aithalrahul34) .

## 🤝 Contributing

1. Fork the repo.
2. Create a feature branch (`git checkout -b feature/amazing-ui`).
3. Commit changes (`git commit -m "Add UI picker"`).
4. Push & open a PR.

Bug reports or Go/FFmpeg tips? Open an issue—love hearing from fellow builders!

## 📄 License

MIT License—use it, tweak it, build on it. See [LICENSE](LICENSE)

## 🙏 Acknowledgments

- [Go Slices Intro](https://go.dev/blog/slices-intro) for saving my sanity.
- [FFmpeg Docs](https://ffmpeg.org/documentation.html) for endless args.
- That 1-hour YouTube tutorial (shoutout: https://youtu.be/8uiZC0l4Ajw).

Built in public by [Rahul Aithal](https://github.com/Rahul-aithal). Questions? [@rahul_aithal on X](https://x.com/rahul_aithal). Star/fork if it sparks ideas! 🚀
