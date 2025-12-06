package video

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func Combainer(meta *Meta) error {
	if meta == nil {
		panic("No data found from user")
	}

	fmt.Println("Starting to add....")
	fileloc := meta.FilePath
	imageLoc := meta.ThumbLocation[meta.SeletedThumbIndex]
	ext := strings.Split(meta.FilePath, ".")[1]
	outputLoc := strings.Split(meta.FilePath, ".")[0] + "TT." + ext
	// ffmpeg -i input.mp4 -i thumbnail.png -map 1 -map 0 -c copy -disposition:0 attached_pic output.mp4
	cmd := exec.Command("ffmpeg",
		"-i",
		fileloc,
		"-i",
		imageLoc,
		"-map",
		"1",
		"-map",
		"0",
		"-c",
		"copy",
		"-disposition:0",
		"attached_pic",
		outputLoc,
	)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ":" + stderr.String())
		return err
	}
	fmt.Println(out.String())
	return nil
}
