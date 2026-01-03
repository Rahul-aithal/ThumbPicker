package video

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"

	"github.com/google/uuid"
)

type CombainerStruct struct {
	FilePath      string
	ThumbLocation string
}

type combainerReturn struct {
	FilePath string
	FileName string
}

func Combainer(meta CombainerStruct) (combainerReturn, error) {

	fileloc := meta.FilePath
	imageLoc := meta.ThumbLocation
	fmt.Println("Starting to add....", fileloc, imageLoc)
	ext := strings.Split(meta.FilePath, ".")[1]
	outputLoc := strings.Split(meta.FilePath, ".")[0] + uuid.NewString()[0:8] + "." + ext
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
		return combainerReturn{}, err
	}
	fmt.Println(out.String())
	outSlice := strings.Split(outputLoc, "/")
	return combainerReturn{
		FilePath: outputLoc,
		FileName: outSlice[len(outSlice)-1],
	}, nil
}
