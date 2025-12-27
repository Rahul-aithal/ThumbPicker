package video

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

func Extractor(meta *Meta) error {
	directory := "./pub/" + strings.Split(filepath.Base(meta.FilePath), ".")[0]

	err := os.MkdirAll(directory, 0o777)
	if err != nil {
		return err
	}

	fileLocs := make([]string, 0, meta.NumberOfFrames)

	for i, val := range meta.TimeStamps {
		fmt.Printf("Proceesing on index %d value %0.2f\n", i, val)

		timestamp := getTimeStampFromSeconds(val)
		fileLoc := directory + "/frame" + strconv.Itoa(i) + ".png"
		cmd := exec.Command("ffmpeg",
			"-ss", timestamp,
			"-i", meta.FilePath,
			// "-vf", "fps=1",
			"-frames:v", "1",
			fileLoc)

		var out bytes.Buffer
		var stderr bytes.Buffer
		cmd.Stdout = &out
		cmd.Stderr = &stderr
		err := cmd.Run()
		if err != nil {
			fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
			return err
		}

		fileLocs = append(fileLocs, fileLoc)

	}
	meta.ThumbLocation = fileLocs
	return nil
}

func getTimeStampFromSeconds(val float64) string {
	min := int(val) / 60
	hour := min / 60
	second := int(val) % 60
	fmt.Println("Time: ", hour, min, second)

	hr := strconv.Itoa(hour)
	mi := strconv.Itoa(min)
	se := strconv.Itoa(second)
	if min < 9 {
		mi = "0" + strconv.Itoa(min)
	}
	if hour < 9 {
		hr = "0" + strconv.Itoa(hour)
	}
	if second < 9 {
		se = "0" + strconv.Itoa(second)
	}

	timestamp := hr + ":" + mi + ":" + se
	return timestamp
}
