package video

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

func MetaData(meta *Meta ) error{
	path := meta.FilePath

	cmd := exec.Command("ffprobe", "-v", "error", "-show_entries", "format=duration",
        "-of", "default=noprint_wrappers=1:nokey=1", path)
    output, err := cmd.Output()
    if err != nil {
	fmt.Println(err)
        return  err
    }
    durationStr := strings.TrimSpace(string(output))
	fmt.Printf("Duration in strings %s\n",output)
    meta.Duration,err = strconv.ParseFloat(durationStr, 64)  
	if err !=nil{
	fmt.Println(err)
		return  err
	}
	fmt.Printf("The duration of the video is found out to be %0.2f sec\n",meta.Duration)
	return nil 	
}
